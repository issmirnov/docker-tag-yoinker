package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/issmirnov/docker-updater/interfaces"
)

func getDockerTags() (res []string, err error) {

	url := "https://registry.hub.docker.com/v1/repositories/sourcegraph/server/tags"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	resp, err := Context.HttpClient.Do(req)
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
