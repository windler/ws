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
	commandString := parseCmd((*cmd).GetCmd(), forceRoot, c)
	execCmd := exec.Command("/bin/sh", "-c", commandString)

	execCmd.Stdin = os.Stdin
	execCmd.Stdout = outStream
	err := execCmd.Run()

	if err != nil {
		panic(err)
	}
}

type customCommandEnv struct {
	WSRoot string
	Args   []string
}

func parseCmd(original string, forceRoot string, c *WSCommandContext) string {
	env := &customCommandEnv{}
	if forceRoot != "" {
		env.WSRoot = forceRoot
	} else {
		env.WSRoot = GetCurrentWorkspace((*c).GetConfig().GetWsDir())
	}

	env.Args = (*c).GetArgs()

	t := template.New("args")

	_, err := t.Parse(original)

	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	t.Execute(buf, env)

	return buf.String()
}
