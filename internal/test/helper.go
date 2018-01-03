package test

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/stretchr/testify/mock"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

//CreateTestRepo create a repo you can use in tests
func CreateTestContext(configFlag string) (*cli.Context, *os.File) {
	file, _ := ioutil.TempFile("", "projherotest")

	fs := flag.FlagSet{}
	fs.String(configFlag, file.Name(), "")
	c := cli.NewContext(nil, &fs, nil)

	os.Setenv("WS_CFG", file.Name())

	return c, file
}

//MockUI creates a mock for the ui interface and calls SetUI. It also mocks every 'anyOfType' expectations
func MockUI() *UIMock {
	u := new(UIMock)

	u.On("PrintHeader", mock.AnythingOfType("string")).Return()
	u.On("PrintString", mock.AnythingOfType("string"), mock.AnythingOfType("Attribute")).Return()
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
func (u *UIMock) PrintString(s string, colorOrNil ...color.Attribute) {
	if colorOrNil == nil {
		u.Called(s)
	} else {
		u.Called(s, colorOrNil[0])
	}
}
