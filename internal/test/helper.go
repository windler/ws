package test

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

//CreateTestRepo create a repo you can use in tests
func CreateTestContext(configFlag string) (*cli.Context, *os.File) {
	file, _ := ioutil.TempFile("", "projherotest")

	fs := flag.FlagSet{}
	fs.String(configFlag, file.Name(), "")
	c := cli.NewContext(nil, &fs, nil)

	return c, file
}
