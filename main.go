package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"time"

	"github.com/issmirnov/docker-updater/config"
	"github.com/issmirnov/docker-updater/docker"
	"github.com/issmirnov/docker-updater/interfaces"
	"github.com/issmirnov/docker-updater/semver"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Commit = "xxxxxx"
var Version = "x.x.x"
var Branch = "x"

var (
	Ctx interfaces.Context
)

// Pass writer. Pass in ioutil.Discard to silence logs.
func setupLogging(logWriter io.Writer, debug bool) {
	output := zerolog.ConsoleWriter{Out: logWriter, TimeFormat: time.RFC822}

	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	Ctx.Logger = zerolog.New(output).With().Timestamp().Logger()
}

func main() {

	var (
		debug      = flag.Bool("debug", false, "enable debug logs")
		configPath = flag.String("config", "config.toml", "config file")
		version    = flag.Bool("v", false, "print version info")
	)
	flag.Parse()

	// set up logging.
	setupLogging(os.Stdout, *debug)

	//_ = config // FIXME
	c, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal().Msgf("Problem loading config file: %s", err.Error())
		return
	}
	Ctx.Config = c

	if *version {
		fmt.Printf("version %s (%s-%s)\n", Version, Branch, Commit)
		os.Exit(0)
	}

	fmt.Println("Hello, starting up...")

	tag := run(Ctx)
	fmt.Print(tag)

}

// separate function, so that we can test this outside of main.
func run(ctx interfaces.Context) (tag string) {
	tags, err := docker.GetDockerTags(ctx)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	tag = semver.MunchTags(tags, ctx).String()
	log.Debug().Msgf("final tag= %s", tag)
	return
}
