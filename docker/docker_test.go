package docker

import (
	"bytes"
	"io/ioutil"
	"testing"

	"net/http"

	"github.com/issmirnov/docker-updater/interfaces"
	"github.com/issmirnov/docker-updater/mocks"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDocker(t *testing.T) {

	Convey("Setup", t, func() {
		// Send testing output to /dev/null
		// setupLogging(ioutil.Discard, false)

		ctx := interfaces.Context{
			HttpClient: &mocks.MockClient{},
		}
		ctx.Logger.Info().Msg("initialized mock client")

		Convey("Test Get Docker Tags", func() {
			Convey("Succeeds", func() {
				// build response JSON: TODO make this sample as docker json results.
				json := `[{"name":"v2.9.4","layer":""}]`
				// create a new reader with that JSON

				mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
					r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
					return &http.Response{
						StatusCode: 200,
						Body:       r,
					}, nil
				}
				resp, err := GetDockerTags(ctx)

				So(err, ShouldBeNil)
				So(resp, ShouldNotBeNil)
				So(resp[0], ShouldEqual, "v2.9.4")

			})
		})
		Convey("Test building tags URL", func() {
			ctx := interfaces.Context{
				Config: interfaces.Config{
					Registry: "foo",
					Image:    "bar/baz",
				},
			}
			url := buildUrl(ctx)

			So(url, ShouldEqual, "foo/bar/baz/tags")

		})
	})
}
