package appcontracts

type Config interface {
	GetWsDir() string
	GetParallelProcessing() int
	GetCustomCommands() []CustomCommand
	GetTableFormat() string
}

type CustomCommand struct {
	Name        string
	Description string
	Cmd         string
	Args        []string
}
