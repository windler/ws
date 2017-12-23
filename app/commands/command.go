package commands

import (
	"github.com/urfave/cli"
)

const (
	//CmdListWs is the ls command
	CmdListWs string = "ls"
	//CmdSetup is the setup command
	CmdSetup string = "setup"
)

//BaseCommand represents wraps the cli commands
type BaseCommand struct {
	Description string
	Aliases     []string
	Command     string
	Action      func(c *cli.Context) error
	Subcommands []BaseCommand
}

//CreateCliCommand creates the command for the cli-app
func CreateCliCommand(factory BaseCommandFactory) *cli.Command {
	bc := factory.CreateCommand()

	cmd := &cli.Command{
		Name:        bc.Command,
		Usage:       bc.Description,
		Aliases:     bc.Aliases,
		ArgsUsage:   "[command]",
		Subcommands: createSubCommands(bc.Subcommands),
	}

	if bc.Action != nil {
		cmd.Action = bc.Action
	}

	return cmd
}

func createSubCommands(cmds []BaseCommand) []cli.Command {
	subCommands := []cli.Command{}
	for _, bc := range cmds {

		command := cli.Command{
			Name:        bc.Command,
			Usage:       bc.Description,
			Aliases:     bc.Aliases,
			ArgsUsage:   "[command]",
			Subcommands: createSubCommands(bc.Subcommands),
		}

		if bc.Action != nil {
			command.Action = bc.Action
		}

		subCommands = append(subCommands, command)
	}
	return subCommands
}

//BaseCommandFactory creates commands
type BaseCommandFactory interface {
	CreateCommand() BaseCommand
}

//CommandAction represents the action executed when command is chosen
type CommandAction interface {
	Exec(c *cli.Context) error
}
