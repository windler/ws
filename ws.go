package main

import (
	"github.com/urfave/cli"
	ws "github.com/windler/ws/app"
	"github.com/windler/ws/app/commands"
	"github.com/windler/ws/app/config"
	"github.com/windler/ws/app/git"
	"github.com/windler/ws/app/ui"
)

func main() {
	app := ws.CreateNewApp("1.0.0")
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

	for _, cmd := range config.Repository().CustomCommands {
		app.AddCommand(&commands.CustomCommandFactory{
			UserInterface: ui,
			Cmd:           cmd,
		})
	}

	app.Start()
}
