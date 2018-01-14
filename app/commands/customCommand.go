package commands

import (
	"fmt"
)

//SetupAppFactory creates commands to list workspace information
type CustomCommandFactory struct {
	UserInterface UI
	Cmd           CustomCommand
}

//CreateCommand creates a ListWsCommand
func (factory *CustomCommandFactory) CreateCommand() BaseCommand {
	return BaseCommand{
		Command:     factory.Cmd.GetName(),
		Description: factory.Cmd.GetDescription(),
		Action: func(c WSCommandContext) {
			factory.action(&c)
		},
	}
}

func (factory *CustomCommandFactory) UI() UI {
	return factory.UserInterface
}

func (factory *CustomCommandFactory) action(c *WSCommandContext) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Custom command is invalid. Check config.", r)
		}
	}()

	var output string
	if (*c).GetFirstArg() != "" {
		ws := GetWorkspaceByPattern((*c).GetConfig().GetWsDir(), (*c).GetFirstArg())
		output = ExecCustomCommand(&factory.Cmd, ws, c)
	} else {
		output = ExecCustomCommandInCurrentWs(&factory.Cmd, c)
	}

	factory.UI().PrintString(output)
}
