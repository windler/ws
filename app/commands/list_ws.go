package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/urfave/cli"
	"github.com/windler/projhero/app/git"
	"github.com/windler/projhero/config"
)

//ListWsFactory creates commands to list workspace information
type ListWsFactory struct{}

//ensure interface
var _ BaseCommandFactory = &ListWsFactory{}

//CreateCommand creates a ListWsCommand
func (factory *ListWsFactory) CreateCommand() BaseCommand {
	return BaseCommand{
		Command:     "workspace",
		Description: "List all workspaces with fancy information",
		Aliases:     []string{"ws"},
		Action:      new(action),
	}
}

type action struct{}

func (a action) Exec(c *cli.Context) error {
	conf := config.Repository()

	fmt.Println(conf.Get(config.WsDir))
	fileInfo, err := ioutil.ReadDir(conf.Get(config.WsDir))
	if err != nil {
		return err
	}

	rows := [][]string{}

	dataChannel := channelFileInfos(fileInfo, conf.Get(config.WsDir))

	fanOutChannels := []<-chan []string{}

	parallel, _ := strconv.Atoi(conf.Get(config.ParallelProcessing))
	fmt.Println(string(parallel))
	for i := 0; i < parallel; i++ {
		fanOutChannels = append(fanOutChannels, collectWsData(dataChannel))
	}

	fanInChannel := fanIn(fanOutChannels)
	for i := 0; i < len(fileInfo); i++ {
		rows = append(rows, <-fanInChannel)
	}

	printTable([]string{"dir", "git status", "branch"}, rows)

	return nil
}

func channelFileInfos(dirs []os.FileInfo, root string) <-chan string {
	out := make(chan string)
	go func() {
		for _, dir := range dirs {
			out <- (root + "/" + dir.Name())
		}
		close(out)
	}()
	return out
}

func collectWsData(in <-chan string) <-chan []string {
	out := make(chan []string)
	go func() {
		for dir := range in {
			g := git.For(dir)
			status := g.Status()
			branch := g.CurrentBranch()

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
