package config

import (
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/issmirnov/docker-tag-yoinker/interfaces"
	"github.com/pelletier/go-toml"
	"github.com/rs/zerolog/log"
)

// LoadConfig takes a file name and the global context. It will save the parsed
// config to the provided global context.
func LoadConfig(fname string, ctx *interfaces.Context) (err error) {
	doc, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Error().Msgf("Could not read config file: %s", err.Error())
		return
	}

	toml.Unmarshal(doc, &ctx.Config)
	ctx.Logger.Debug().Msg(spew.Sprint(ctx.Config))

	return nil
}
