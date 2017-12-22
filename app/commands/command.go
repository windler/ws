package commands

import (
	"github.com/urfave/cli"
)

type BaseCommand struct {
	Name string
}

func CreateCliCommand(factory BaseCommandFactory) *cli.Command {
	bc := factory.CreateCommand()

	return &cli.Command{
		Name: bc.Name,
	}
}

type BaseCommandFactory interface {
	CreateCommand() BaseCommand
}
