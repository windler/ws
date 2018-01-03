package commands

import (
	"github.com/urfave/cli"
	"github.com/windler/ws/app/config"
)

//SetupAppFactory creates commands to list workspace information
type CustomCommandFactory struct {
	UserInterface UI
	cmd           config.CustomCommand
}

//ensure interface
var _ BaseCommandFactory = &CustomCommandFactory{}

//CreateCommand creates a ListWsCommand
func (factory *CustomCommandFactory) CreateCommand() BaseCommand {
	return BaseCommand{
		Command:     factory.cmd.Name,
		Description: "Configure ws to unleash the beauty. Alternatively, you can edit your personal config file.",
		Action: func(c *cli.Context) error {
			return factory.action(c)
		},
	}
}

func (factory *CustomCommandFactory) UI() UI {
	return factory.UserInterface
}

func (factory *CustomCommandFactory) action(c *cli.Context) error {
	/*	factory.UI().PrintHeader(factory.cmd.Command)
		data, err := exec.Command(factory.cmd.Command, factory.cmd.Args...).Output()
		if err != nil {
			panic(err)
		}

		factory.UI().PrintString(string(data))
	*/
	return nil
}
