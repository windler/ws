package commands

import (
	"bytes"
	"errors"
	"html/template"
	"sort"
	"strings"

	"github.com/fatih/color"

	"github.com/urfave/cli"
	"github.com/windler/ws/app/commands/internal/commandCommons"
	"github.com/windler/ws/app/config"
)

//ListWsFactory creates commands to list workspace information
type ListWsFactory struct {
	InfoRetriever commandCommons.WsInfoRetriever
	UserInterface UI
}

type tableData [][]string

//CreateCommand creates a ListWsCommand
func (factory *ListWsFactory) CreateCommand() BaseCommand {

	return BaseCommand{
		Command:     CmdListWs,
		Description: "List all workspaces with fancy information.",
		Aliases:     []string{},
		Action: func(c *cli.Context) error {
			return factory.listWsExecAll(c)
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "table",
				Usage: "formats the table using the `template`",
			},
		},
		Subcommands: []BaseCommand{},
	}
}

func (factory *ListWsFactory) UI() UI {
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
	conf := config.Repository()

	wsDir := conf.WsDir

	if wsDir == "" {
		factory.UI().PrintHeader("Panic!")
		factory.UI().PrintString(" >> No workspaces defined to scan <<", color.FgRed)
		RecommendFromError(CmdSetup, factory.UI())

		return nil
	}

	dirs := []string{}
	if onlyCurrent {
		dirs = append(dirs, commandCommons.GetCurrentWorkspace(wsDir))
	} else {
		dirs = commandCommons.GetWsDirs(wsDir)
	}

	dataChannel := factory.channelFileInfos(dirs)
	fanOutChannels := []<-chan []string{}

	tableFormat := getTableFormat(c)

	if conf.ParallelProcessing == 0 {
		return errors.New("ParallelProcessing has to be > 0")
	}

	for i := 0; i < conf.ParallelProcessing; i++ {
		fanOutChannels = append(fanOutChannels, factory.collectWsData(dataChannel, onlyCurrent, tableFormat))
	}

	rows := tableData{}
	fanInChannel := factory.fanIn(fanOutChannels)
	for i := 0; i < len(dirs); i++ {
		rows = append(rows, <-fanInChannel)
	}

	if len(rows) > 0 {
		sort.Sort(rows)

		funcMap := commandCommons.GetHeaderFunctionMap()

		buf := new(bytes.Buffer)
		t := template.Must(template.New("header").Funcs(funcMap).Parse(tableFormat))
		t.Execute(buf, "")

		factory.UI().PrintTable(strings.Split(buf.String(), "|"), rows)
	} else {
		factory.printError(onlyCurrent)
	}

	return nil
}

func getTableFormat(c *cli.Context) string {
	conf := config.Repository()

	tableFormat := "{{wsRoot .}}|{{gitStatus .}}|{{gitBranch .}}"
	if c.String("table") != "" {
		tableFormat = c.String("table")
	} else if conf.TableFormat != "" {
		tableFormat = conf.TableFormat
	}

	return tableFormat
}

func (factory *ListWsFactory) printError(onlyCurrent bool) {
	if onlyCurrent {
		factory.UI().PrintString("Current directory is not within a workspace.", color.FgYellow)
	} else {
		factory.UI().PrintString("No workspaces found!", color.FgRed)
		RecommendFromError(CmdSetup, factory.UI())
	}
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

func (factory *ListWsFactory) collectWsData(in <-chan string, onlyCurrent bool, pattern string) <-chan []string {
	out := make(chan []string)
	go func() {
		funcMap := commandCommons.GetRowsFunctionMap(factory.InfoRetriever, !onlyCurrent)

		for dir := range in {
			buf := new(bytes.Buffer)
			t := template.Must(template.New("table").Funcs(funcMap).Parse(pattern))
			t.Execute(buf, dir)

			out <- strings.Split(buf.String(), "|")
		}
		close(out)
	}()
	return out
}

type tableTemplateData struct {
	Dir string
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
