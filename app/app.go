package app

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/urfave/cli"
	"github.com/windler/projhero/app/commands"
	"github.com/windler/projhero/config"
)

//ProjHeroApp is the main cli app
type ProjHeroApp struct {
	app *cli.App
}

//CreateNewApp create a new app with the given version
func CreateNewApp(version string) *ProjHeroApp {
	a := &ProjHeroApp{
		app: cli.NewApp(),
	}

	a.configureApp(version)

	return a
}

func (app ProjHeroApp) configureApp(version string) {
	cliApp := app.app

	cliApp.Name = "projhero"
	cliApp.Description = "Dev Workspace Swiss Knife"
	cliApp.Author = "Nico Windler"
	cliApp.Copyright = "2017"
	cliApp.Email = "nico.windler@gmail.com"
	cliApp.Version = version

	cliApp.EnableBashCompletion = true

	cliApp.Action = func(c *cli.Context) {
		fmt.Println("Welcome star")
	}

	usr, err := user.Current()

	if err != nil {
		log.Fatal("can not obtain user ", err)
	}

	config.Prepare(usr.HomeDir + "/.projherocfg")
}

//AddCommand adds a new cli command
func (app *ProjHeroApp) AddCommand(factory commands.BaseCommandFactory) {
	command := commands.CreateCliCommand(factory)

	cliApp := app.app
	cliApp.Commands = append(cliApp.Commands, *command)
}

//Start launches the cli app
func (app *ProjHeroApp) Start() {
	app.app.Run(os.Args)
}
