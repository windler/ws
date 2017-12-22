package main

import (
	projhero "github.com/windler/projhero/app"
	"github.com/windler/projhero/app/commands"
)

func main() {
	app := projhero.CreateNewApp("1.0.0")
	app.AddCommand(&commands.ListWsFactory{})
	app.Start()
}
