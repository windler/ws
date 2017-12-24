package commands

import (
	"errors"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/windler/workspacehero/app/commands/contracts"

	"github.com/fatih/color"

	"github.com/urfave/cli"
	"github.com/windler/workspacehero/config"
)

//ListWsFactory creates commands to list workspace information
type ListWsFactory struct {
	InfoRetriever contracts.WsInfoRetriever
	UserInterface contracts.UI
}

type tableData [][]string

//ensure interface
var (
	_ BaseCommandFactory = &ListWsFactory{}
)

//CreateCommand creates a ListWsCommand
func (factory *ListWsFactory) CreateCommand() BaseCommand {

	return BaseCommand{
		Command:     CmdListWs,
		Description: "List all workspaces with fancy information.",
		Aliases:     []string{},
		Action: func(c *cli.Context) error {
			return factory.listWsExecAll(c)
		},
		Subcommands: []BaseCommand{},
	}
}

func (factory *ListWsFactory) UI() contracts.UI {
	return factory.UserInterface
}

func (c tableData) Len() int           { return len(c) }
func (c tableData) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c tableData) Less(i, j int) bool { return strings.Compare(c[i][0], c[j][0]) == -1 }

//ListWsExecCurrent prints infos about the current ws-directory
func (factory *ListWsFactory) ListWsExecCurrent(c *cli.Context) error {
	return factory.listWsExec(c, true)
}

func (factory *ListWsFactory) listWsExecAll(c *cli.Context) error {
	return factory.listWsExec(c, false)
}

func (factory *ListWsFactory) listWsExec(c *cli.Context, onlyCurrent bool) error {
	conf := config.Repository(c)

	wsDir := conf.WsDir

	if wsDir == "" {
		factory.UI().PrintHeader("Panic!")
		factory.UI().PrintString(" >> No workspaces defined to scan <<", color.FgRed)
		RecommendFromError(CmdSetup, factory.UI())

		return nil
	}

	fileInfo, err := ioutil.ReadDir(wsDir)
	if err != nil {
		return err
	}
	dirs := factory.getDirs(fileInfo, wsDir, onlyCurrent)

	rows := tableData{}

	dataChannel := factory.channelFileInfos(dirs)

	fanOutChannels := []<-chan []string{}

	if conf.ParallelProcessing == 0 {
		return errors.New("ParallelProcessing has to be > 0")
	}

	for i := 0; i < conf.ParallelProcessing; i++ {
		fanOutChannels = append(fanOutChannels, factory.collectWsData(dataChannel, onlyCurrent))
	}

	fanInChannel := factory.fanIn(fanOutChannels)
	for i := 0; i < len(dirs); i++ {
		rows = append(rows, <-fanInChannel)
	}

	if len(rows) > 0 {
		sort.Sort(rows)
		factory.UI().PrintTable([]string{"dir", "git status", "branch"}, rows)
	} else {
		factory.UI().PrintString("No workspaces found!", color.FgRed)
		RecommendFromError(CmdSetup, factory.UI())
	}

	return nil
}

func (factory *ListWsFactory) getDirs(fileInfos []os.FileInfo, root string, onlyCurrent bool) []string {
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

func (factory *ListWsFactory) channelFileInfos(dirs []string) <-chan string {
	out := make(chan string, len(dirs))
	go func() {
		for _, dir := range dirs {
			out <- dir
		}
		close(out)
	}()
	return out
}

func (factory *ListWsFactory) collectWsData(in <-chan string, onlyCurrent bool) <-chan []string {
	out := make(chan []string)
	go func() {
		wd, _ := os.Getwd()

		for dir := range in {
			status := factory.InfoRetriever.Status(dir)
			branch := factory.InfoRetriever.CurrentBranch(dir)

			if !onlyCurrent && strings.HasPrefix(wd, dir) {
				dir = dir + " <--"
			}

			out <- []string{dir, status, branch}
		}
		close(out)
	}()
	return out
}

func (factory *ListWsFactory) fanIn(input []<-chan []string) <-chan []string {
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
