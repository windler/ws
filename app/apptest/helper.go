package apptest

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/stretchr/testify/mock"
	"github.com/windler/ws/app/appcontracts"
)

func CreateTestContext() (TestContext, *os.File) {
	return createTestContext("", false)
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

func (c TestContext) GetFirstArg() string {
	return ""
}

func (c TestContext) GetConfig() appcontracts.Config {
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
func (c testCfg) GetCustomCommands() []appcontracts.CustomCommand {
	return []appcontracts.CustomCommand{}
}
func (c testCfg) GetTableFormat() string {
	return ""
}

func MockUI() *UIMock {
	u := new(UIMock)

	u.On("PrintHeader", mock.AnythingOfType("string")).Return()
	u.On("PrintString", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return()
	u.On("PrintString", mock.AnythingOfType("string")).Return()

	return u
}

type UIMock struct {
	mock.Mock
}

func (u *UIMock) PrintHeader(s string) {
	u.Called(s)
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
