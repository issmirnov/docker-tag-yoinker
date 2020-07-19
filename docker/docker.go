package docker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"strings"

	"github.com/issmirnov/docker-updater/config"
	"github.com/issmirnov/docker-updater/interfaces"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger(config.AppName)

func buildUrl(ctx interfaces.Context) string {
	//if ctx.Config.Registry
	if !strings.HasSuffix(ctx.Config.Registry, "/") {
		ctx.Config.Registry += "/"
		log.Warning("Registry url did not have trailing slash, adding automatically. Please fix config.")
	}
	return ctx.Config.Registry + ctx.Config.Image + "/tags"
}
func GetDockerTags(ctx interfaces.Context) (res []string, err error) {

	url := buildUrl(ctx)

	// url := "https://registry.hub.docker.com/v1/repositories/sourcegraph/server/tags"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	resp, err := ctx.HttpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	//res = string(body)
	log.Debug(string(body))

	tags := []interfaces.Tag{}
	err = json.Unmarshal(body, &tags)
	for _, t := range tags {
		res = append(res, t.Name)
	}
	return
}
