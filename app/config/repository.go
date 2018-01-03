package config

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"sync"

	"github.com/windler/ws/app/common"

	yaml "gopkg.in/yaml.v2"
)

//Config represents all app configurations
type Config struct {
	WsDir              string
	ParallelProcessing int
	CustomCommands     []CustomCommand
	TableFormat        string
}

type CustomCommand struct {
	Name        string
	Description string
	Cmd         string
	Args        []string
}

var (
	cfg     *Config
	cfgFile string
	once    sync.Once
)

const (
	ConfigFlag string = "config"
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
func Repository() *Config {
	once.Do(func() {
		if cfg == nil {
			createCfg()
		}
	})
	return cfg
}

func createCfg() {
	cfg = &Config{
		ParallelProcessing: 3,
	}

	ensureCfgFile()

	d, err := ioutil.ReadFile(cfgFile)

	if err != nil {
		log.Fatal("Cannot read file ", err)
	}

	yaml.Unmarshal(d, &cfg)
}

func ensureCfgFile() {

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

		setupDefaultValues()
		cfg.Save()
	}
}

func setupDefaultValues() {
	cfg.ParallelProcessing = 3
}

//Save persists the current configuration
func (c Config) Save() {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatal("could not marshall configuration ", err)
	}

	err = ioutil.WriteFile(cfgFile, data, 0644)
	if err != nil {
		log.Fatal("could not write configuration ", err)
	}
}
