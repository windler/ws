package config

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"github.com/windler/ws/app/appcontracts"
	"github.com/windler/ws/app/common"

	yaml "gopkg.in/yaml.v2"
)

//Config represents all app configurations
type config struct {
	WsDir              string
	ParallelProcessing int
	CustomCommands     []appcontracts.CustomCommand
	TableFormat        string
}

func (c config) GetWsDir() string {
	return c.WsDir
}

func (c config) GetParallelProcessing() int {
	return c.ParallelProcessing
}

func (c config) GetCustomCommands() []appcontracts.CustomCommand {
	return c.CustomCommands
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

//SetConfigFile sets the config to use (when not set default path will be user)
func SetConfigFile(file string) {
	cfgFile = file
}

//Repository returns the config for the app
func CreateYamlRepository() appcontracts.Config {
	cfg := &config{
		ParallelProcessing: 3,
	}

	cfg.ensureCfgFile()

	d, err := ioutil.ReadFile(cfgFile)

	if err != nil {
		log.Fatal("Cannot read file ", err)
	}

	yaml.Unmarshal(d, &cfg)

	return cfg
}

func (c config) ensureCfgFile() {

	if os.Getenv("WS_CFG") == "" {
		usr, err := user.Current()

		if err != nil {
			log.Fatal("can not obtain user ", err)
		}
		cfgFile = usr.HomeDir + "/.wshero"
	} else {
		cfgFile = common.EnsureFileFormat(os.Getenv("WS_CFG"))
	}

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		f, e := os.Create(cfgFile)
		if e != nil {
			log.Fatal("Cannot create file ", err)
		}
		defer f.Close()
	}
}

func (c config) setupDefaultValues() {
	c.ParallelProcessing = 3
}
