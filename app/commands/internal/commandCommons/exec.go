package commandCommons

import (
	"bytes"
	"html/template"
	"os/exec"

	"github.com/windler/ws/app/config"
)

func ExecCustomCommandInCurrentWs(cmd *config.CustomCommand) string {
	return execCustomCommand(cmd, "")
}

func ExecCustomCommand(cmd *config.CustomCommand, ws string) string {
	return execCustomCommand(cmd, ws)
}

func execCustomCommand(cmd *config.CustomCommand, forceRoot string) string {
	args := getArgs(cmd.Args, forceRoot)
	data, err := exec.Command(cmd.Cmd, args...).Output()

	if err != nil {
		panic(err)
	}

	return string(data)
}

type customCommandEnv struct {
	WSRoot string
}

func getArgs(original []string, forceRoot string) []string {
	result := []string{}
	env := &customCommandEnv{}
	if forceRoot != "" {
		env.WSRoot = forceRoot
	} else {
		env.WSRoot = GetCurrentWorkspace(config.Repository().WsDir)
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
