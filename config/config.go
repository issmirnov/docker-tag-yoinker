package config

import (
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/issmirnov/docker-tag-yoinker/interfaces"
	"github.com/pelletier/go-toml"
	"github.com/rs/zerolog/log"
)

func LoadConfig(fname string, ctx *interfaces.Context) (err error) {
	doc, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Error().Msgf("Could not read config file: %s", err.Error())
		return
	}

	//config := interfaces.Config{}
	toml.Unmarshal(doc, &ctx.Config)
	ctx.Logger.Debug().Msg(spew.Sprint(ctx.Config))

	return nil
}
