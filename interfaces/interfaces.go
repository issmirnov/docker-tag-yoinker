package interfaces

import "net/http"

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Tag is a Docker Hub JSON element.
type Tag struct {
	Name  string
	Layer string
}

type Context struct {
	Config     Config
	HttpClient HTTPClient
}

type Config struct {
	Blacklist   []string
	Whitelist   string
	StripPrefix string `toml:"strip_prefix"`
	StripSuffix string `toml:"strip_suffix"`
}
