package main

import (
	"fmt"
	"io"

	// "log" // TODO implement logging (like my other projects)

	"flag"

	"os"

	"github.com/issmirnov/docker-updater/config"
	"github.com/issmirnov/docker-updater/interfaces"
	"github.com/op/go-logging"
)

const appName = "docker-updater"

var log = logging.MustGetLogger(appName)

var Commit = "xxxxxx"
var Version = "x.x.x"
var Branch = "x"

var (
	Context interfaces.Context
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
		logging.SetLevel(logging.DEBUG, appName)
	} else {
		logging.SetLevel(logging.WARNING, appName)
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
	Context.Config = config.LoadConfig(*configPath)

	if *version {
		fmt.Printf("%s version %s (%s-%s)\n", appName, Version, Branch, Commit)
		os.Exit(0)
	}

	fmt.Println("Hello, starting up...")

	tags, err := getDockerTags()
	if err != nil {
		log.Fatal(err)
		return
	}

	tag := MunchTags(tags, Context)
	log.Debugf("final tag= %s", tag)

}
