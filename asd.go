package main

import (
	"github.com/urfave/cli"
	asd "github.com/windler/workspacehero/app"
	"github.com/windler/workspacehero/app/commands"
	"github.com/windler/workspacehero/app/git"
	"github.com/windler/workspacehero/app/ui"
)

func main() {
	app := asd.CreateNewApp("0.0.1")
	ui := ui.ConsoleUI{}

	listWsFactory := &commands.ListWsFactory{
		InfoRetriever: git.New(),
		UserInterface: ui,
	}

	app.SetAction(func(c *cli.Context) error {
		return listWsFactory.ListWsExecCurrent(c)
	})

	app.AddCommand(listWsFactory)
	app.AddCommand(&commands.SetupAppFactory{
		UserInterface: ui,
	})

	app.Start()
}
