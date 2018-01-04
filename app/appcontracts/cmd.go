package appcontracts

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
	GetFirstArg() string
	GetConfig() Config
}
