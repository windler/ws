package commandCommons

import (
	"bytes"
	"html/template"
	"os/exec"

	"github.com/windler/ws/app/appcontracts"
)

func ExecCustomCommandInCurrentWs(cmd *appcontracts.CustomCommand, c *appcontracts.WSCommandContext) string {
	return execCustomCommand(cmd, "", c)
}

func ExecCustomCommand(cmd *appcontracts.CustomCommand, ws string, c *appcontracts.WSCommandContext) string {
	return execCustomCommand(cmd, ws, c)
}

func execCustomCommand(cmd *appcontracts.CustomCommand, forceRoot string, c *appcontracts.WSCommandContext) string {
	args := getArgs(cmd.Args, forceRoot, c)
	data, err := exec.Command(cmd.Cmd, args...).Output()

	if err != nil {
		panic(err)
	}

	return string(data)
}

type customCommandEnv struct {
	WSRoot string
}

func getArgs(original []string, forceRoot string, c *appcontracts.WSCommandContext) []string {
	result := []string{}
	env := &customCommandEnv{}
	if forceRoot != "" {
		env.WSRoot = forceRoot
	} else {
		env.WSRoot = GetCurrentWorkspace((*c).GetConfig().GetWsDir())
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
