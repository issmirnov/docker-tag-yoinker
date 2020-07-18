package internal

import (
	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	"github.com/issmirnov/docker-updater/interfaces"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("docker-updater")

var validConfig = []byte(`
blacklist = ["insiders", "rc"]
whitelist = "alpine"
strip_prefix = "linux-arm-"
strip_suffix = "-alpine"
`)

func LoadValidTestConfig() interfaces.Config {
	config := interfaces.Config{}
	_ = toml.Unmarshal(validConfig, &config) // ignore error.
	log.Debug(spew.Print(config))
	return config
}
