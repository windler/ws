package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/windler/ws/app/commands"
	yaml "gopkg.in/yaml.v2"
)

//Config represents all app configurations
type config struct {
	WsDir              string
	ParallelProcessing int
	CustomCommands     []customCommand
	TableFormat        string
}

type customCommand struct {
	Name        string
	Description string
	Cmd         string
	Args        []string
}

func (c customCommand) GetName() string {
	return c.Name
}

func (c customCommand) GetDescription() string {
	return c.Description
}

func (c customCommand) GetCmd() string {
	return c.Cmd
}

func (c customCommand) GetArgs() []string {
	return c.Args
}

func (c config) GetWsDir() string {
	return c.WsDir
}

func (c config) GetParallelProcessing() int {
	return c.ParallelProcessing
}

func (c config) GetCustomCommands() []commands.CustomCommand {
	ret := []commands.CustomCommand{}
	for _, c := range c.CustomCommands {
		ret = append(ret, c)
	}
	return ret
}

func (c config) GetTableFormat() string {
	return c.TableFormat
}

var (
	cfgFile string
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

//Repository returns the config for the app
func CreateYamlRepository() config {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error loading config.", r)
		}
	}()

	cfg := &config{
		ParallelProcessing: 3,
	}

	cfg.ensureCfgFile()

	d, err := ioutil.ReadFile(cfgFile)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(d, &cfg)
	if err != nil {
		panic(err)
	}

	return *cfg
}

func (c config) ensureCfgFile() {
	if os.Getenv("WS_CFG") == "" {
		usr, err := user.Current()

		if err != nil {
			log.Fatal("can not obtain user ", err)
		}
		cfgFile = usr.HomeDir + "/.wshero"
	} else {
		cfgFile = ensureFileFormat(os.Getenv("WS_CFG"))
	}

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		f, e := os.Create(cfgFile)
		if e != nil {
			log.Fatal("Cannot create file ", err)
		}
		defer f.Close()
	}
}

func ensureFileFormat(dir string) string {
	value := strings.TrimSpace(dir)
	value = ensureHomeDir(value)

	return value
}

func ensureHomeDir(dir string) string {
	result := dir
	if strings.HasPrefix(result, "~") {
		usr, _ := user.Current()
		result = strings.Replace(result, "~", usr.HomeDir, 1)
	}

	return result
}

func ensureDirSuffic(dir string) string {
	result := dir

	if !strings.HasSuffix(result, "/") {
		result = result + "/"
	}

	return result
}

func (c config) setupDefaultValues() {
	c.ParallelProcessing = 3
}
