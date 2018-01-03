package commands

import (
	"bytes"
	"fmt"
	"html/template"
	"os/exec"

	"github.com/fatih/color"
	"github.com/windler/ws/app/common"

	"github.com/urfave/cli"
	"github.com/windler/ws/app/config"
)

//SetupAppFactory creates commands to list workspace information
type CustomCommandFactory struct {
	UserInterface UI
	Cmd           config.CustomCommand
}

//ensure interface
var _ BaseCommandFactory = &CustomCommandFactory{}

//CreateCommand creates a ListWsCommand
func (factory *CustomCommandFactory) CreateCommand() BaseCommand {
	return BaseCommand{
		Command:     factory.Cmd.Name,
		Description: factory.Cmd.Description,
		Action: func(c *cli.Context) error {
			return factory.action(c)
		},
	}
}

func (factory *CustomCommandFactory) UI() UI {
	return factory.UserInterface
}

func (factory *CustomCommandFactory) action(c *cli.Context) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Custom command is invalid. Check config.", r)
		}
	}()

	factory.UI().PrintString(factory.Cmd.Name+":", color.FgGreen)

	args := getArgs(factory.Cmd.Args)
	data, err := exec.Command(factory.Cmd.Cmd, args...).Output()

	if err != nil {
		panic(err)
	}

	factory.UI().PrintString(string(data))

	return err
}

type customCommandEnv struct {
	WSRoot string
}

func getArgs(original []string) []string {
	result := []string{}
	env := &customCommandEnv{
		WSRoot: common.GetWsDirs(config.Repository().WsDir, true)[0],
	}

	for _, arg := range original {
		t := template.New("args")

		_, err := t.Parse(arg)

		if err != nil {
			panic(err)
		}

		buf := new(bytes.Buffer)
		t.Execute(buf, env)

		result = append(result, buf.String())
	}

	return result
}
