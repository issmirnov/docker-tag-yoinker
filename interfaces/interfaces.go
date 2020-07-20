package interfaces

import (
	"net/http"

	"github.com/rs/zerolog"
)

// HTTPClient interface, used for mocks.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Tag is a Docker Hub JSON element.
type Tag struct {
	Name  string
	Layer string
}

// Context is the global context.
type Context struct {
	// Parsed config for app run.
	Config Config
	// We provide a HttpClient instance, so that we can inject mocks.
	HttpClient HTTPClient
	// Global loggeer instance.
	Logger zerolog.Logger
}

// Config is user provided config.toml
type Config struct {
	// Name of image: "sourcegraph/server"
	Image string
	// Optional: registry url. Provides default.
	Registry    string `default:"https://registry.hub.docker.com/v1/repositories/"`
	Blacklist   []string
	Whitelist   string
	StripPrefix string `toml:"strip_prefix"`
	StripSuffix string `toml:"strip_suffix"`
}
