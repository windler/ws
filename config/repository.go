package config

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

//Config represents all app configurations
type Config struct {
	WsDir string
}

var (
	cfg     Config
	cfgFile string
)

//Prepare ensures that at least an empty config exists and inits the repo
func Prepare(filename string) {
	cfgFile = filename

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		f, e := os.Create(filename)
		if e != nil {
			log.Fatal("Cannot create file ", err)
		}
		defer f.Close()
	}

	d, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal("Cannot read file ", err)
	}

	yaml.Unmarshal(d, &cfg)
}

//Get retrievs the whole config as a struct
func Get() Config {
	return cfg
}

//Save persists the configuration
func (newCfg Config) Save() {
	cfg = newCfg

	data, err := yaml.Marshal(newCfg)
	if err != nil {
		log.Fatal("could not marshall configuration ", err)
	}

	err = ioutil.WriteFile(cfgFile, data, 0644)
	if err != nil {
		log.Fatal("could not write configuration ", err)
	}
}
