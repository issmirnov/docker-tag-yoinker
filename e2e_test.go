package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/issmirnov/docker-updater/config"
	"github.com/issmirnov/docker-updater/interfaces"
	"github.com/issmirnov/docker-updater/mocks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestE2E(t *testing.T) {

	Convey("Setup", t, func() {
		// Send testing output to /dev/null
		// setupLogging(ioutil.Discard, false)

		ctx := interfaces.Context{
			HttpClient: &mocks.MockClient{},
			//Config:     internal.LoadValidTestConfig(),
		}
		log.Info("initialized mock client")

		Convey("Run main app", func() {

			Convey("Test on cached sourcegraph data", func() {

				tagsResp, err := ioutil.ReadFile("./testdata/sourcegraph/tags.json")
				So(err, ShouldBeNil)

				conf, err := config.LoadConfig("./testdata/sourcegraph/config.toml")
				So(conf.Image, ShouldEqual, "sourcegraph/server")
				ctx.Config = conf

				mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
					r := ioutil.NopCloser(bytes.NewReader([]byte(tagsResp)))
					return &http.Response{
						StatusCode: 200,
						Body:       r,
					}, nil
				}

				tag := run(ctx)

				So(tag, ShouldEqual, "3.17.3")

			})
		})

	})
}
