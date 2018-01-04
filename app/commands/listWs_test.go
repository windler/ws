package commands

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/windler/ws/app/apptest"
	"github.com/windler/ws/app/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLsWsCommand(t *testing.T) {
	f := new(ListWsFactory).CreateCommand()

	assert.Equal(t, "ls", f.Command)
	assert.Equal(t, []string{}, f.Aliases)
	assert.Equal(t, "List all workspaces with fancy information.", f.Description)
}

func TestListWsNoWsDefined(t *testing.T) {
	ui := apptest.MockUI()
	f := ListWsFactory{
		UserInterface: ui,
	}

	c, _ := apptest.CreateTestContextWithWsDir("")

	f.CreateCommand().Action(c)

	ui.AssertCalled(t, "PrintString", " >> No workspaces defined to scan <<", "red")
}

func TestListNoDirs(t *testing.T) {
	ui := apptest.MockUI()

	f := ListWsFactory{
		UserInterface: ui,
	}

	tmpWsDir, _ := ioutil.TempDir("", "projherotest")
	c, _ := apptest.CreateTestContextWithWsDir(tmpWsDir)

	f.CreateCommand().Action(c)

	ui.AssertCalled(t, "PrintString", "No workspaces found!", "red")
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
	ui := apptest.MockUI()
	infoRetriever := new(testInfoRetriever)

	f := ListWsFactory{
		UserInterface: ui,
		InfoRetriever: infoRetriever,
	}

	tmpWsDir, _ := ioutil.TempDir("", "wshero")
	tmpWsDir = common.EnsureDirFormat(tmpWsDir)
	ws1 := tmpWsDir + "ws1"
	ws2 := tmpWsDir + "ws2"

	c, _ := apptest.CreateTestContextWithWsDir(tmpWsDir)

	os.MkdirAll(ws1, os.ModePerm)
	os.MkdirAll(ws2, os.ModePerm)

	infoRetriever.On("Status", ws1).Return("super")
	infoRetriever.On("Status", ws2).Return("bad")

	infoRetriever.On("CurrentBranch", ws1).Return("master")
	infoRetriever.On("CurrentBranch", ws2).Return("someBranch")

	ui.On("PrintTable", mock.Anything, mock.Anything).Return()

	f.CreateCommand().Action(c)

	ui.AssertCalled(t, "PrintTable", []string{"ws", "git status", "git branch"}, [][]string{
		[]string{ws1, "super", "master"},
		[]string{ws2, "bad", "someBranch"},
	})
}
