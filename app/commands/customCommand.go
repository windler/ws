package commands

import (
	"fmt"
)

//SetupAppFactory creates commands to list workspace information
type CustomCommandFactory struct {
	UserInterface UI
	Cmd           CustomCommand
	WSRetriever   WorkspaceRetriever
	Executor      CustomCommandExecutor
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

	if (*c).GetArgs()[0] != "" {
		ws := factory.WSRetriever.GetWorkspaceByPattern((*c).GetConfig().GetWsDir(), (*c).GetArgs()[0])
		factory.Executor.Exec(&factory.Cmd, ws, c)
	} else {
		factory.Executor.ExecInCurrentWs(&factory.Cmd, c)
	}
}
