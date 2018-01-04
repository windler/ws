package main

import (
	ws "github.com/windler/ws/app"
	"github.com/windler/ws/app/commands"
	"github.com/windler/ws/app/config"
	"github.com/windler/ws/app/git"
	"github.com/windler/ws/app/ui"
)

func main() {
	yamlRepo := config.CreateYamlRepository()
	app := ws.CreateNewApp("1.1.0", yamlRepo)
	ui := ui.ConsoleUI{}

	listWsFactory := &commands.ListWsFactory{
		InfoRetriever: git.New(),
		UserInterface: ui,
	}

	app.AddCommand(listWsFactory.CreateCommand(), yamlRepo)

	for _, cmd := range yamlRepo.GetCustomCommands() {
		ccFactory := &commands.CustomCommandFactory{
			UserInterface: ui,
			Cmd:           cmd,
		}
		app.AddCommand(ccFactory.CreateCommand(), yamlRepo)
	}

	app.Start()
}
