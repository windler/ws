package app

import (
	"os"

	"github.com/urfave/cli"
	"github.com/windler/asd/app/commands"
	"github.com/windler/asd/config"
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

	cliApp.Name = "asd"
	cliApp.Description = "Dev Workspace Swiss Knife."
	cliApp.Usage = "workspace hero"
	cliApp.Author = "Nico Windler"
	cliApp.Copyright = "2017"
	cliApp.Version = version

	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  config.ConfigFlag + ", c",
			Usage: "Load configuration from `FILE`",
		},
	}

	cliApp.EnableBashCompletion = true
}

func (app *ProjHeroApp) SetAction(fn func(c *cli.Context) error) {
	app.app.Action = fn
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
