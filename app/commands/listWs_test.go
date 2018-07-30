package commands_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/windler/ws/app/commands"
	"github.com/windler/ws/app/commands/testfiles/mocks"
)

func TestLsWsCommand(t *testing.T) {
	f := new(commands.ListWsFactory).CreateCommand()

	assert.Equal(t, "ls", f.Command)
	assert.Equal(t, []string{}, f.Aliases)
	assert.Equal(t, "List all workspaces with fancy information.", f.Description)
}

func TestListWsNoWsDefined(t *testing.T) {
	ui := createMockUI()
	f := commands.ListWsFactory{
		UserInterface: ui,
	}

	contextMock := createContextMock("", "")

	f.CreateCommand().Action(contextMock)

	ui.AssertCalled(t, "PrintString", "Panic!", "red")
	ui.AssertCalled(t, "PrintString", " >> No workspaces defined to scan <<")
}

func TestListNoDirs(t *testing.T) {
	ui := createMockUI()

	wsRetrieverMock := &mocks.WorkspaceRetriever{}
	wsRetrieverMock.On("GetWorkspacesIn", "/usr/home/workspaces/").Return([]string{})

	f := commands.ListWsFactory{
		UserInterface: ui,
		WSRetriever:   wsRetrieverMock,
	}

	contextMock := createContextMock("/usr/home/workspaces/", "")
	f.CreateCommand().Action(contextMock)

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
	ui := createMockUI()
	infoRetriever := new(testInfoRetriever)

	ws1 := "/usr/home/workspaces/ws1"
	ws2 := "/usr/home/workspaces/ws2"

	infoRetriever.On("Status", ws1).Return("super")
	infoRetriever.On("Status", ws2).Return("bad")

	infoRetriever.On("CurrentBranch", ws1).Return("master")
	infoRetriever.On("CurrentBranch", ws2).Return("someBranch")

	ui.On("PrintTable", mock.Anything, mock.Anything).Return()

	wsRetrieverMock := &mocks.WorkspaceRetriever{}
	wsRetrieverMock.On("GetWorkspacesIn", "/usr/home/workspaces/").Return([]string{
		ws1, ws2,
	})

	f := commands.ListWsFactory{
		UserInterface: ui,
		InfoRetriever: infoRetriever,
		WSRetriever:   wsRetrieverMock,
	}

	contextMock := createContextMock("/usr/home/workspaces/", "")
	f.CreateCommand().Action(contextMock)

	ui.AssertCalled(t, "PrintTable", []string{"ws", "git status", "git branch"}, [][]string{
		[]string{ws1, "super", "master"},
		[]string{ws2, "bad", "someBranch"},
	})
}

func createMockUI() *mocks.UI {
	u := &mocks.UI{}

	u.On("PrintString", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return()
	u.On("PrintString", mock.AnythingOfType("string")).Return()

	return u
}

func createContextMock(workspaceRoot, tableFormat string) *mocks.WSCommandContext {
	ctxMock := &mocks.WSCommandContext{}
	configMock := &mocks.Config{}

	configMock.On("GetWsDir").Return(workspaceRoot)
	configMock.On("GetTableFormat").Return(tableFormat)
	configMock.On("GetParallelProcessing").Return(1)

	ctxMock.On("GetStringFlag", "table").Return("")
	ctxMock.On("GetConfig").Return(configMock)

	return ctxMock
}
