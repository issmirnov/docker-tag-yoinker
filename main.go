package main

import (
	"fmt"
	"io"

	// "log" // TODO implement logging (like my other projects)

	"flag"

	"os"

	"github.com/issmirnov/docker-updater/config"
	"github.com/issmirnov/docker-updater/docker"
	"github.com/issmirnov/docker-updater/interfaces"
	"github.com/issmirnov/docker-updater/semver"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger(config.AppName)

var Commit = "xxxxxx"
var Version = "x.x.x"
var Branch = "x"

var (
	Ctx interfaces.Context
)

// Pass writer. Pass in ioutil.Discard to silence logs.
func setupLogging(logWriter io.Writer, debug bool) {
	var format = logging.MustStringFormatter(
		`%{color}%{shortfile} (%{shortfunc}) ▶ %{color:reset}%{message}`,
	)
	backend1 := logging.NewLogBackend(logWriter, "", 0)
	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	logging.SetBackend(backend1Formatter)

	if debug {
		logging.SetLevel(logging.DEBUG, config.AppName)
	} else {
		logging.SetLevel(logging.WARNING, config.AppName)
	}
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
		log.Fatalf("Problem loading config file: %s", err.Error())
		return
	}
	Ctx.Config = c

	if *version {
		fmt.Printf("%s version %s (%s-%s)\n", config.AppName, Version, Branch, Commit)
		os.Exit(0)
	}

	fmt.Println("Hello, starting up...")

	run(Ctx)

}

// separate function, so that we can test this outside of main.
func run(ctx interfaces.Context) (tag string) {
	tags, err := docker.GetDockerTags(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	tag = semver.MunchTags(tags, ctx).String()
	log.Debugf("final tag= %s", tag)
	return
}
