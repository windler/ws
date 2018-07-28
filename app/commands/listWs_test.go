package commands_test

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/windler/ws/app/commands"
)

func TestLsWsCommand(t *testing.T) {
	f := new(commands.ListWsFactory).CreateCommand()

	assert.Equal(t, "ls", f.Command)
	assert.Equal(t, []string{}, f.Aliases)
	assert.Equal(t, "List all workspaces with fancy information.", f.Description)
}

func TestListWsNoWsDefined(t *testing.T) {
	ui := MockUI()
	f := commands.ListWsFactory{
		UserInterface: ui,
	}

	c, _ := CreateTestContextWithWsDir("")

	f.CreateCommand().Action(c)

	ui.AssertCalled(t, "PrintString", "Panic!", "red")
	ui.AssertCalled(t, "PrintString", " >> No workspaces defined to scan <<")
}

func TestListNoDirs(t *testing.T) {
	ui := MockUI()

	f := commands.ListWsFactory{
		UserInterface: ui,
	}

	tmpWsDir, _ := ioutil.TempDir("", "projherotest")
	c, _ := CreateTestContextWithWsDir(tmpWsDir)

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
	ui := MockUI()
	infoRetriever := new(testInfoRetriever)

	f := commands.ListWsFactory{
		UserInterface: ui,
		InfoRetriever: infoRetriever,
	}

	tmpWsDir, _ := ioutil.TempDir("", "wshero")
	tmpWsDir = tmpWsDir + "/"
	ws1 := tmpWsDir + "ws1"
	ws2 := tmpWsDir + "ws2"

	c, _ := CreateTestContextWithWsDir(tmpWsDir)

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

func CreateTestContextWithWsDir(dir string) (TestContext, *os.File) {
	return createTestContext(dir, true)
}

func createTestContext(dir string, useDir bool) (TestContext, *os.File) {
	file, _ := ioutil.TempFile("", "projherotest")

	fs := flag.FlagSet{}
	fs.String("config", file.Name(), "")
	os.Setenv("WS_CFG", file.Name())

	wsdir := file.Name()
	if useDir {
		wsdir = dir
	}

	return TestContext{
		cfg: testCfg{
			wsdir: wsdir,
		},
	}, file
}

type TestContext struct {
	cfg testCfg
}

func (c TestContext) GetStringFlag(flag string) string {
	return ""
}

func (c TestContext) GetBoolFlag(flag string) bool {
	return false
}

func (c TestContext) GetIntFlag(flag string) int {
	return 0
}

func (c TestContext) GetArgs() []string {
	return []string{}
}

func (c TestContext) GetConfig() commands.Config {
	return c.cfg
}

type testCfg struct {
	wsdir string
}

func (c testCfg) GetWsDir() string {
	return c.wsdir
}
func (c testCfg) GetParallelProcessing() int {
	return 0
}
func (c testCfg) GetCustomCommands() []commands.CustomCommand {
	return []commands.CustomCommand{}
}
func (c testCfg) GetTableFormat() string {
	return ""
}

func MockUI() *UIMock {
	u := new(UIMock)

	u.On("PrintString", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return()
	u.On("PrintString", mock.AnythingOfType("string")).Return()

	return u
}

type UIMock struct {
	mock.Mock
}

func (u *UIMock) PrintTable(header []string, rows [][]string) {
	u.Called(header, rows)
}

func (u *UIMock) PrintString(s string, colorOrNil ...string) {
	if colorOrNil == nil {
		u.Called(s)
	} else {
		u.Called(s, colorOrNil[0])
	}
}
