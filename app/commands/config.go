package commands

type Config interface {
	GetWsDir() string
	GetParallelProcessing() int
	GetCustomCommands() []CustomCommand
	GetTableFormat() string
}

type CustomCommand interface {
	GetName() string
	GetDescription() string
	GetCmd() string
	GetArgs() []string
}
