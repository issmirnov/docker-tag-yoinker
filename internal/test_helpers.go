package internal

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/issmirnov/docker-updater/interfaces"
)

var validConfig = []byte(`
blacklist = ["insiders", "rc"]
whitelist = "alpine"
strip_prefix = "linux-arm-"
strip_suffix = "-alpine"
image = "bar/baz"
registry = "foo"
`)

func LoadValidTestConfig() interfaces.Config {
	config := interfaces.Config{}
	err := toml.Unmarshal(validConfig, &config) // ignore error.
	if err != nil {
		log.Fatalf("programmer error. Fix the test: %s", err.Error())
	}
	return config
}
