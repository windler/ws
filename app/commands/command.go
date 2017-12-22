package commands

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

//BaseCommand represents wraps the cli commands
type BaseCommand struct {
	Description string
	Aliases     []string
	Command     string
	Action      CommandAction
}

//CreateCliCommand creates the command for the cli-app
func CreateCliCommand(factory BaseCommandFactory) *cli.Command {
	bc := factory.CreateCommand()

	return &cli.Command{
		Name:      bc.Command,
		Usage:     bc.Description,
		Aliases:   bc.Aliases,
		ArgsUsage: "[command]",
		Action:    bc.Action.Exec,
	}
}

//BaseCommandFactory creates commands
type BaseCommandFactory interface {
	CreateCommand() BaseCommand
}

//CommandAction represents the action executed when command is chosen
type CommandAction interface {
	Exec(c *cli.Context) error
}

func printTable(header []string, rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorder(false)
	table.AppendBulk(rows)
	table.Render()
}
