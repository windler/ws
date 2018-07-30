package commands

import (
	"bytes"
	"html/template"
	"io"
	"os"
	"os/exec"
)

type CustomCommandExecutor interface {
	ExecInCurrentWs(cmd *CustomCommand, c *WSCommandContext)
	Exec(cmd *CustomCommand, ws string, c *WSCommandContext)
	ExecToString(cmd *CustomCommand, ws string, c *WSCommandContext) string
}

type SHExecutor struct {
	WSRetriever WorkspaceRetriever
}

func (e SHExecutor) ExecInCurrentWs(cmd *CustomCommand, c *WSCommandContext) {
	e.execCustomCommand(cmd, "", c, os.Stdout)
}

func (e SHExecutor) Exec(cmd *CustomCommand, ws string, c *WSCommandContext) {
	e.execCustomCommand(cmd, ws, c, os.Stdout)
}

func (e SHExecutor) ExecToString(cmd *CustomCommand, ws string, c *WSCommandContext) string {
	buf := bytes.NewBufferString("")
	e.execCustomCommand(cmd, ws, c, buf)
	return buf.String()
}

func (e SHExecutor) execCustomCommand(cmd *CustomCommand, forceRoot string, c *WSCommandContext, outStream io.Writer) {
	commandString := e.parseCmd((*cmd).GetCmd(), forceRoot, c)
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

func (e SHExecutor) parseCmd(original string, forceRoot string, c *WSCommandContext) string {
	env := &customCommandEnv{}
	if forceRoot != "" {
		env.WSRoot = forceRoot
	} else {
		workspaceRoot := (*c).GetConfig().GetWsDir()
		env.WSRoot = e.WSRetriever.GetCurrentWorkspace(workspaceRoot)
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
