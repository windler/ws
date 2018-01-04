package commands

import (
	"fmt"

	"github.com/windler/ws/app/appcontracts"
	"github.com/windler/ws/app/commands/internal/commandCommons"
)

//SetupAppFactory creates commands to list workspace information
type CustomCommandFactory struct {
	UserInterface UI
	Cmd           appcontracts.CustomCommand
}

//CreateCommand creates a ListWsCommand
func (factory *CustomCommandFactory) CreateCommand() BaseCommand {
	return BaseCommand{
		Command:     factory.Cmd.Name,
		Description: factory.Cmd.Description,
		Action: func(c appcontracts.WSCommandContext) {
			factory.action(&c)
		},
	}
}

func (factory *CustomCommandFactory) UI() UI {
	return factory.UserInterface
}

func (factory *CustomCommandFactory) action(c *appcontracts.WSCommandContext) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Custom command is invalid. Check config.", r)
		}
	}()

	factory.UI().PrintString(factory.Cmd.Name+":", "green")

	var output string
	if (*c).GetFirstArg() != "" {
		ws := commandCommons.GetWorkspaceByPattern((*c).GetConfig().GetWsDir(), (*c).GetFirstArg())
		output = commandCommons.ExecCustomCommand(&factory.Cmd, ws, c)
	} else {
		output = commandCommons.ExecCustomCommandInCurrentWs(&factory.Cmd, c)
	}

	factory.UI().PrintString(output)
}
