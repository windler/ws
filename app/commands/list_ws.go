package commands

import (
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/windler/workspacehero/app/common"

	"github.com/urfave/cli"
	"github.com/windler/workspacehero/app/git"
	"github.com/windler/workspacehero/config"
)

//ListWsFactory creates commands to list workspace information
type ListWsFactory struct{}

type tableData [][]string

//ensure interface
var _ BaseCommandFactory = &ListWsFactory{}

//CreateCommand creates a ListWsCommand
func (factory *ListWsFactory) CreateCommand() BaseCommand {

	return BaseCommand{
		Command:     CmdListWs,
		Description: "List all workspaces with fancy information.",
		Aliases:     []string{},
		Action:      listWsExecAll,
		Subcommands: []BaseCommand{},
	}
}

func (c tableData) Len() int           { return len(c) }
func (c tableData) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c tableData) Less(i, j int) bool { return strings.Compare(c[i][0], c[j][0]) == -1 }

//ListWsExecCurrent prints infos about the current ws-directory
func ListWsExecCurrent(c *cli.Context) error {
	return listWsExec(c, true)
}

func listWsExecAll(c *cli.Context) error {
	return listWsExec(c, false)
}

func listWsExec(c *cli.Context, onlyCurrent bool) error {
	conf := config.Repository(c)

	wsDir := conf.WsDir

	if wsDir == "" {
		common.PrintHeader("Panic!")
		color.Red(" >> No workspaces defined to scan <<")
		common.RecommendFromError(CmdSetup)
		return nil
	}

	fileInfo, err := ioutil.ReadDir(wsDir)
	if err != nil {
		return err
	}
	dirs := getDirs(fileInfo, wsDir, onlyCurrent)

	rows := tableData{}

	dataChannel := channelFileInfos(dirs)

	fanOutChannels := []<-chan []string{}

	for i := 0; i < conf.ParallelProcessing; i++ {
		fanOutChannels = append(fanOutChannels, collectWsData(dataChannel, onlyCurrent))
	}

	fanInChannel := fanIn(fanOutChannels)
	for i := 0; i < len(dirs); i++ {
		rows = append(rows, <-fanInChannel)
	}

	if len(rows) > 0 {
		sort.Sort(rows)
		printTable([]string{"dir", "git status", "branch"}, rows)
	} else {
		color.Red("No workspaces found!")
		common.RecommendFromError(CmdSetup)
	}

	return nil
}

func getDirs(fileInfos []os.FileInfo, root string, onlyCurrent bool) []string {
	result := []string{}

	wd, _ := os.Getwd()
	for _, dir := range fileInfos {
		fullDir := (root + dir.Name())
		if !onlyCurrent || strings.HasPrefix(wd, fullDir) {
			result = append(result, fullDir)
		}
	}

	return result
}

func channelFileInfos(dirs []string) <-chan string {
	out := make(chan string)
	go func() {
		for _, dir := range dirs {
			out <- dir
		}
		close(out)
	}()
	return out
}

func collectWsData(in <-chan string, onlyCurrent bool) <-chan []string {
	out := make(chan []string)
	go func() {
		wd, _ := os.Getwd()

		for dir := range in {
			g := git.For(dir)
			status := g.Status()
			branch := g.CurrentBranch()

			if !onlyCurrent && strings.HasPrefix(wd, dir) {
				dir = dir + " <--"
			}

			out <- []string{dir, status.Status, branch}
		}
		close(out)
	}()
	return out
}

func fanIn(input []<-chan []string) <-chan []string {
	c := make(chan []string)
	for _, ch := range input {
		go func(channel <-chan []string) {
			for msg := range channel {
				c <- msg
			}
		}(ch)
	}

	return c
}
