package config

import (
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/issmirnov/docker-updater/interfaces"
	"github.com/op/go-logging"
	"github.com/pelletier/go-toml"
)

const AppName = "docker-updater"

var log = logging.MustGetLogger(AppName)

func LoadConfig(fname string) (config interfaces.Config, err error) {
	doc, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Errorf("Could not read config file: %s", err.Error())
		return
	}

	config = interfaces.Config{}
	toml.Unmarshal(doc, &config)
	log.Debug(spew.Print(config))
	return config, nil
}
