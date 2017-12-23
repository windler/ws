package main

import (
	asd "github.com/windler/workspacehero/app"
	"github.com/windler/workspacehero/app/commands"
)

func main() {
	app := asd.CreateNewApp("0.0.1")

	app.AddCommand(&commands.ListWsFactory{})
	app.AddCommand(&commands.SetupAppFactory{})

	app.Start()
}
