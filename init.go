package main

import (
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	Ctx.HttpClient = &http.Client{}
	log.Logger = log.Output(zerolog.ConsoleWriter{
		//TimeFormat: time.RFC822
	})
	zerolog.SetGlobalLevel(zerolog.WarnLevel)

}
