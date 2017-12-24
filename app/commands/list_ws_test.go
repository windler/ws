package commands

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/windler/asd/app/common"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli"
	"github.com/windler/asd/config"
	"github.com/windler/asd/internal/test"
)

func TestLsWsCommand(t *testing.T) {
	f := new(ListWsFactory).CreateCommand()

	assert.Equal(t, "ls", f.Command)
	assert.Equal(t, []string{}, f.Aliases)
	assert.Equal(t, "List all workspaces with fancy information.", f.Description)
}

func TestListWsNoWsDefined(t *testing.T) {
	ui := test.MockUI()
	f := ListWsFactory{
		UserInterface: ui,
	}

	c, _ := test.CreateTestContext(config.ConfigFlag)
	config.Repository(c)

	f.CreateCommand().Action(&cli.Context{})

	ui.AssertCalled(t, "PrintString", " >> No workspaces defined to scan <<", color.FgRed)
}

func TestListNoDirs(t *testing.T) {
	ui := test.MockUI()

	f := ListWsFactory{
		UserInterface: ui,
	}

	c, _ := test.CreateTestContext(config.ConfigFlag)

	tmpWsDir, _ := ioutil.TempDir("", "projherotest")
	config.Repository(c).WsDir = tmpWsDir

	f.CreateCommand().Action(&cli.Context{})

	ui.AssertCalled(t, "PrintString", "No workspaces found!", color.FgRed)
}

type testInfoRetriever struct {
	mock.Mock
}

func (t testInfoRetriever) Status(ws string) string {
	args := t.Called(ws)
	return args.String(0)
}

func (t testInfoRetriever) CurrentBranch(ws string) string {
	args := t.Called(ws)
	return args.String(0)
}

func TestList(t *testing.T) {
	ui := test.MockUI()
	infoRetriever := new(testInfoRetriever)

	f := ListWsFactory{
		UserInterface: ui,
		InfoRetriever: infoRetriever,
	}

	c, _ := test.CreateTestContext(config.ConfigFlag)

	tmpWsDir, _ := ioutil.TempDir("", "wshero")
	tmpWsDir = common.EnsureDirFormat(tmpWsDir)
	ws1 := tmpWsDir + "ws1"
	ws2 := tmpWsDir + "ws2"

	os.MkdirAll(ws1, os.ModePerm)
	os.MkdirAll(ws2, os.ModePerm)

	config.Repository(c).WsDir = tmpWsDir

	infoRetriever.On("Status", ws1).Return("super")
	infoRetriever.On("Status", ws2).Return("bad")

	infoRetriever.On("CurrentBranch", ws1).Return("master")
	infoRetriever.On("CurrentBranch", ws2).Return("someBranch")

	ui.On("PrintTable", mock.Anything, mock.Anything).Return()

	f.CreateCommand().Action(&cli.Context{})

	ui.AssertCalled(t, "PrintTable", []string{"dir", "git status", "branch"}, [][]string{
		[]string{ws1, "super", "master"},
		[]string{ws2, "bad", "someBranch"},
	})
}
