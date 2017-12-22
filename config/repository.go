package config

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

const (
	ParallelProcessing string = "parallelProcessing"
	WsDir              string = "wsDir"
)

//Config represents all app configurations
type config struct {
	items map[string]string
	mu    sync.RWMutex
}

//repo impl. inspired by http://blog.ralch.com/tutorial/design-patterns/golang-singleton/
var (
	cfg     *config
	cfgFile string
	once    sync.Once
)

//Prepare ensures that at least an empty config exists and inits the repo
func Prepare(filename string) {
	repo := Repository()
	cfgFile = filename

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		f, e := os.Create(filename)
		if e != nil {
			log.Fatal("Cannot create file ", err)
		}
		defer f.Close()

		repo.Save()
	}

	d, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal("Cannot read file ", err)
	}

	yaml.Unmarshal(d, &cfg.items)
}

func (r *config) Set(key, data string) {
	cfg.mu.Lock()
	defer r.mu.Unlock()
	cfg.items[key] = data
}

func (r *config) Get(key string) string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := cfg.items[key]
	if !ok {
		return ""
	}
	return item
}

func Repository() *config {
	once.Do(func() {
		if cfg == nil {
			cfg = &config{
				items: make(map[string]string),
			}
		}
	})
	return cfg
}

//Save persists the configuration
func (newCfg config) Save() {
	cfg.mu.RLock()
	defer cfg.mu.RUnlock()

	data, err := yaml.Marshal(cfg.items)
	if err != nil {
		log.Fatal("could not marshall configuration ", err)
	}

	err = ioutil.WriteFile(cfgFile, data, 0644)
	if err != nil {
		log.Fatal("could not write configuration ", err)
	}
}
