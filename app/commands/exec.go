package commands

import (
	"bytes"
	"html/template"
	"io"
	"os"
	"os/exec"
)

func ExecCustomCommandInCurrentWs(cmd *CustomCommand, c *WSCommandContext) {
	execCustomCommand(cmd, "", c, os.Stdout)
}

func ExecCustomCommand(cmd *CustomCommand, ws string, c *WSCommandContext) {
	execCustomCommand(cmd, ws, c, os.Stdout)
}

func ExecCustomCommandToString(cmd *CustomCommand, ws string, c *WSCommandContext) string {
	buf := bytes.NewBufferString("")
	execCustomCommand(cmd, ws, c, buf)
	return buf.String()
}

func execCustomCommand(cmd *CustomCommand, forceRoot string, c *WSCommandContext, outStream io.Writer) {
	args := getArgs((*cmd).GetArgs(), forceRoot, c)
	execCmd := exec.Command((*cmd).GetCmd(), args...)

	execCmd.Stdin = os.Stdin
	execCmd.Stdout = outStream
	err := execCmd.Run()

	if err != nil {
		panic(err)
	}
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
