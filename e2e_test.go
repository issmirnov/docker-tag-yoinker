package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"fmt"

	"github.com/issmirnov/docker-tag-yoinker/config"
	"github.com/issmirnov/docker-tag-yoinker/interfaces"
	"github.com/issmirnov/docker-tag-yoinker/mocks"
	"github.com/rs/zerolog"
	. "github.com/smartystreets/goconvey/convey"
)

func TestE2E(t *testing.T) {

	Convey("Setup", t, func() {
		// Send testing output to /dev/null
		zerolog.SetGlobalLevel(zerolog.FatalLevel)

		ctx := interfaces.Context{
			HttpClient: &mocks.MockClient{},
		}
		ctx.Logger.Info().Msg("initialized mock client")

		Convey("Run main app end to end", func() {

			Convey("Test on cached data", func() {

				// use a test table to avoid repeating test boilerplate
				cases := map[string]string{
					"sourcegraph": "3.17.3",
					"portainer":   "1.24.0",
					"minecraft":   "1.6.0",
					"ghost":       "3.25.0",
				}

				for target, version := range cases {
					Convey(fmt.Sprintf("Test on cached  '%s' data", target), func() {
						tagsResp, err := ioutil.ReadFile(fmt.Sprintf("./testdata/%s/tags.json", target))
						So(err, ShouldBeNil)

						err = config.LoadConfig(fmt.Sprintf("./testdata/%s/config.toml", target), &ctx)
						So(err, ShouldBeNil)
						So(ctx.Config.Image, ShouldContainSubstring, target)

						mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
							r := ioutil.NopCloser(bytes.NewReader([]byte(tagsResp)))
							return &http.Response{
								StatusCode: 200,
								Body:       r,
							}, nil
						}

						tag := run(ctx)

						So(tag, ShouldEqual, version)
					})
				}

			})

		})

	})
}
