package docker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"strings"

	"github.com/issmirnov/docker-tag-yoinker/interfaces"
)

// buildUrl is a tiny helper function, extracted for better test coverage
func buildUrl(ctx interfaces.Context) string {
	if !strings.HasSuffix(ctx.Config.Registry, "/") {
		ctx.Config.Registry += "/"
		ctx.Logger.Warn().Msg("Registry url did not have trailing slash, adding automatically. Please fix config.")
	}
	return ctx.Config.Registry + ctx.Config.Image + "/tags"
}

// GetDockerTags will query the provided registry and return a list of tags as strings.
func GetDockerTags(ctx interfaces.Context) (res []string, err error) {
	req, err := http.NewRequest(http.MethodGet, buildUrl(ctx), nil)
	if err != nil {
		ctx.Logger.Fatal().Msg(err.Error())
		return
	}

	resp, err := ctx.HttpClient.Do(req)
	if err != nil {
		ctx.Logger.Fatal().Msg(err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.Logger.Fatal().Msg(err.Error())
		return
	}

	ctx.Logger.Debug().Msg(string(body))

	tags := []interfaces.Tag{}
	err = json.Unmarshal(body, &tags)
	if err != nil {
		ctx.Logger.Fatal().Msg(err.Error())
		return
	}

	for _, t := range tags {
		res = append(res, t.Name)
	}
	return
}
