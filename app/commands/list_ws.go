package commands

type ListWsFactory struct{}

//ensure interface
var _ BaseCommandFactory = &ListWsFactory{}

func (factory *ListWsFactory) CreateCommand() BaseCommand {
	return BaseCommand{
		Name: "List Workspaces",
	}
}
