package main

import (
	asd "github.com/windler/workspacehero/app"
	"github.com/windler/workspacehero/app/commands"
	"github.com/windler/workspacehero/app/ui"
)

func main() {
	ui.SetUI(ui.ConsoleUI{})

	app := asd.CreateNewApp("0.0.1")

	app.AddCommand(&commands.ListWsFactory{})
	app.AddCommand(&commands.SetupAppFactory{})

	app.Start()
}
