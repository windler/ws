package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/windler/ws/app/commands"
)

func TestDefaultApp(t *testing.T) {
	app := CreateNewApp("myVersion", nil)

	assert.Equal(t, "myVersion", app.app.Version)
	assert.Equal(t, "ws", app.app.Name)
	assert.Equal(t, "Dev Workspace Swiss Knife.", app.app.Description)
	assert.Equal(t, "workspace hero", app.app.Usage)
	assert.Equal(t, "Nico Windler", app.app.Author)
	assert.Equal(t, "2017", app.app.Copyright)
	assert.True(t, app.app.EnableBashCompletion)
}

func TestAddCommands(t *testing.T) {
	app := CreateNewApp("myVersion", nil)
	assert.True(t, len(app.app.Commands) == 0)

	app.AddCommand(BaseCommand{
		Command: "testcommand",
	}, nil)

	assert.True(t, len(app.app.Commands) == 1)
	assert.Equal(t, "testcommand", app.app.Commands[0].Name)
}

type BaseCommand struct {
	Description string
	Aliases     []string
	Command     string
	Action      func(c commands.WSCommandContext)
	Subcommands []commands.WSCommand
	Flags       []commands.WSCommandFlag
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

func (b BaseCommand) GetAction() func(c commands.WSCommandContext) {
	return b.Action
}

func (b BaseCommand) GetSubcommands() []commands.WSCommand {
	return b.Subcommands
}

func (b BaseCommand) GetStringFlags() []commands.WSCommandFlag {
	return b.Flags
}
