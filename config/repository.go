package config

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"sync"

	"github.com/urfave/cli"
	"github.com/windler/workspacehero/app/common"

	yaml "gopkg.in/yaml.v2"
)

//Config represents all app configurations
type Config struct {
	WsDir              string
	ParallelProcessing int
}

var (
	cfg     *Config
	cfgFile string
	once    sync.Once
	dirVals = []string{}
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
func Repository(c *cli.Context) *Config {
	once.Do(func() {
		if cfg == nil {
			createCfg(c)
		}
	})
	return cfg
}

func createCfg(c *cli.Context) {
	cfg = &Config{}

	ensureCfgFile(c)

	d, err := ioutil.ReadFile(cfgFile)

	if err != nil {
		log.Fatal("Cannot read file ", err)
	}

	yaml.Unmarshal(d, &cfg)
}

func ensureCfgFile(c *cli.Context) {

	if c.String(ConfigFlag) == "" {
		usr, err := user.Current()

		if err != nil {
			log.Fatal("can not obtain user ", err)
		}
		cfgFile = usr.HomeDir + "/.projherocfg"
	} else {
		cfgFile = common.EnsureFileFormat(c.String(ConfigFlag))
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
