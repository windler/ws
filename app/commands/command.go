package commands

type WSCommand interface {
	GetDescription() string
	GetAliases() []string
	GetCommand() string
	GetAction() func(c WSCommandContext)
	GetSubcommands() []WSCommand
	GetStringFlags() []WSCommandFlag
}

type WSCommandFlag interface {
	GetName() string
	GetUsage() string
	GetType() string
}

type WSCommandContext interface {
	GetStringFlag(flag string) string
	GetBoolFlag(flag string) bool
	GetIntFlag(flag string) int
	GetArgs() []string
	GetConfig() Config
}

type BaseCommand struct {
	Description string
	Aliases     []string
	Command     string
	Action      func(c WSCommandContext)
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

func (b BaseCommand) GetAction() func(c WSCommandContext) {
	return b.Action
}

func (b BaseCommand) GetSubcommands() []WSCommand {
	res := []WSCommand{}
	for _, sc := range b.Subcommands {
		res = append(res, sc)
	}
	return res
}

func (b BaseCommand) GetStringFlags() []WSCommandFlag {
	res := []WSCommandFlag{}
	for _, f := range b.Flags {
		res = append(res, f)
	}
	return res
}
