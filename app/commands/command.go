package commands

import (
	"github.com/windler/ws/app/appcontracts"
)

//BaseCommand represents wraps the cli commands
type BaseCommand struct {
	Description string
	Aliases     []string
	Command     string
	Action      func(c appcontracts.WSCommandContext)
	Subcommands []BaseCommand
	Flags       []StringFlag
}

func (b BaseCommand) GetDescription() string {
	return b.Description
}

func (b BaseCommand) GetAliases() []string {
	return b.Aliases
}

func (b BaseCommand) GetCommand() string {
	return b.Command
}

func (b BaseCommand) GetAction() func(c appcontracts.WSCommandContext) {
	return b.Action
}

func (b BaseCommand) GetSubcommands() []appcontracts.WSCommand {
	res := []appcontracts.WSCommand{}
	for _, sc := range b.Subcommands {
		res = append(res, sc)
	}
	return res
}

func (b BaseCommand) GetStringFlags() []appcontracts.WSCommandFlag {
	res := []appcontracts.WSCommandFlag{}
	for _, f := range b.Flags {
		res = append(res, f)
	}
	return res
}
