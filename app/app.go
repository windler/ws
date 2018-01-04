package app

import (
	"os"

	"github.com/urfave/cli"
	"github.com/windler/ws/app/appcontracts"
)

//ProjHeroApp is the main cli app
type ProjHeroApp struct {
	app *cli.App
}

//CreateNewApp create a new app with the given version
func CreateNewApp(version string, cfg appcontracts.Config) *ProjHeroApp {
	a := &ProjHeroApp{
		app: cli.NewApp(),
	}

	a.configureApp(version, cfg)

	return a
}

func (app ProjHeroApp) configureApp(version string, cfg appcontracts.Config) {
	cliApp := app.app

	cliApp.Name = "ws"
	cliApp.Description = "Dev Workspace Swiss Knife."
	cliApp.Usage = "workspace hero"
	cliApp.Author = "Nico Windler"
	cliApp.Copyright = "2017"
	cliApp.Version = version

	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	cliApp.EnableBashCompletion = true
}

func (app *ProjHeroApp) SetAction(fn func(c *cli.Context) error) {
	app.app.Action = fn
}

//AddCommand adds a new cli command
func (app *ProjHeroApp) AddCommand(cmd appcontracts.WSCommand, cfg appcontracts.Config) {
	command := CreateCliCommand(cmd, cfg)

	cliApp := app.app
	cliApp.Commands = append(cliApp.Commands, *command)
}

//Start launches the cli app
func (app *ProjHeroApp) Start() {
	app.app.Run(os.Args)
}
