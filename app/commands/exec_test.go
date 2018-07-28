package commands_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/windler/ws/app/commands"
	"github.com/windler/ws/app/commands/testfiles/mocks"
)

func TestExec(t *testing.T) {
	customCmd := createCustomCommand("echo printThis")
	cmdContext := createContext([]string{}, "root")

	res := commands.ExecCustomCommandToString(&customCmd, "/usr/home/workspaces/", &cmdContext)

	assert.Equal(t, "printThis\n", res)
}

func TestExecWSRootForce(t *testing.T) {
	customCmd := createCustomCommand("echo {{.WSRoot}}")
	cmdContext := createContext([]string{"my"}, "/usr/home/workspaces")

	res := commands.ExecCustomCommandToString(&customCmd, "/usr/home/workspaces/ws1", &cmdContext)

	assert.Equal(t, "/usr/home/workspaces/ws1\n", res)
}

func TestExecWSArgs(t *testing.T) {
	customCmd := createCustomCommand("echo {{index .Args 0}}")
	cmdContext := createContext([]string{"my"}, "/usr/home/workspaces")

	res := commands.ExecCustomCommandToString(&customCmd, "", &cmdContext)

	assert.Equal(t, "my\n", res)
}

func createCustomCommand(cmd string) commands.CustomCommand {
	customCommand := &mocks.CustomCommand{}
	customCommand.On("GetCmd").Return(cmd)

	return customCommand
}

func createContext(args []string, wsRoot string) commands.WSCommandContext {
	cfg := createConfig(wsRoot)

	cmdContext := &mocks.WSCommandContext{}
	cmdContext.On("GetArgs").Return(args)
	cmdContext.On("GetConfig").Return(cfg)

	return cmdContext
}

func createConfig(wsRoot string) commands.Config {
	cfg := &mocks.Config{}
	cfg.On("GetWsDir").Return(wsRoot)

	return cfg
}
