package config

import (
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/issmirnov/docker-updater/interfaces"
	"github.com/op/go-logging"
	"github.com/pelletier/go-toml"
)

var log = logging.MustGetLogger("docker-updater")

func LoadConfig(fname string) interfaces.Config {
	doc, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("Could not read config file: %s", err.Error())
	}

	config := interfaces.Config{}
	toml.Unmarshal(doc, &config)
	log.Debug(spew.Print(config))
	return config
}
