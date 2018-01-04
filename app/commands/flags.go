package commands

type StringFlag struct {
	Name  string
	Usage string
}

func (f StringFlag) GetName() string {
	return f.Name
}

func (f StringFlag) GetUsage() string {
	return f.Usage
}

func (f StringFlag) GetType() string {
	return "string"
}
