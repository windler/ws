package commands

import (
	"bytes"
	"html/template"
	"os/exec"
)

func ExecCustomCommandInCurrentWs(cmd *CustomCommand, c *WSCommandContext) string {
	return execCustomCommand(cmd, "", c)
}

func ExecCustomCommand(cmd *CustomCommand, ws string, c *WSCommandContext) string {
	return execCustomCommand(cmd, ws, c)
}

func execCustomCommand(cmd *CustomCommand, forceRoot string, c *WSCommandContext) string {
	args := getArgs((*cmd).GetArgs(), forceRoot, c)
	data, err := exec.Command((*cmd).GetCmd(), args...).Output()

	if err != nil {
		panic(err)
	}

	return string(data)
}

type customCommandEnv struct {
	WSRoot string
}

func getArgs(original []string, forceRoot string, c *WSCommandContext) []string {
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
