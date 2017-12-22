package app

import (
	"testing"

	"github.com/windler/projhero/app/commands"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDefaultApp(t *testing.T) {
	app := CreateNewApp("myVersion")

	assert.Equal(t, "myVersion", app.app.Version)
	assert.Equal(t, "projhero", app.app.Name)
	assert.Equal(t, "Dev Workspace Swiss Knife", app.app.Description)
	assert.Equal(t, "Dev Workspace Swiss Knife", app.app.Usage)
	assert.Equal(t, "Nico Windler", app.app.Author)
	assert.Equal(t, "2017", app.app.Copyright)
	assert.Equal(t, "nico.windler@gmail.com", app.app.Email)
	assert.True(t, app.app.EnableBashCompletion)
}

type BaseCommandFactoryMock struct {
	mock.Mock
}

func (m *BaseCommandFactoryMock) CreateCommand() commands.BaseCommand {
	args := m.Called()

	return args.Get(0).(commands.BaseCommand)
}

func TestAddCommands(t *testing.T) {
	app := CreateNewApp("myVersion")
	assert.True(t, len(app.app.Commands) == 0)

	factory := new(BaseCommandFactoryMock)
	factory.On("CreateCommand").Return(commands.BaseCommand{
		Command: "testcommand",
	})

	app.AddCommand(factory)

	assert.True(t, len(app.app.Commands) == 1)
	assert.Equal(t, "testcommand", app.app.Commands[0].Name)
}
