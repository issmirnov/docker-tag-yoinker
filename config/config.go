package config

import (
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/issmirnov/docker-updater/interfaces"
	"github.com/pelletier/go-toml"
	"github.com/rs/zerolog/log"
)

func LoadConfig(fname string) (config interfaces.Config, err error) {
	doc, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Error().Msgf("Could not read config file: %s", err.Error())
		return
	}

	config = interfaces.Config{}
	toml.Unmarshal(doc, &config)
	log.Debug().Msg(spew.Sprint(config))
	return config, nil
}
