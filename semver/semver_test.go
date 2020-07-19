package semver

import (
	"testing"

	"github.com/issmirnov/docker-tag-yoinker/interfaces"
	"github.com/issmirnov/docker-tag-yoinker/internal"
	"github.com/rs/zerolog"
	. "github.com/smartystreets/goconvey/convey"
)

const targetVersion = "linux-arm-v3.11-alpine"

func TestSemver(t *testing.T) {

	// Send testing output to /dev/null
	zerolog.SetGlobalLevel(zerolog.FatalLevel)

	Convey("Setup", t, func() {
		// Send testing output to /dev/null
		// setupLogging(ioutil.Discard, false)
		ctx := interfaces.Context{
			Config: internal.LoadValidTestConfig(),
		}

		tags := []string{"v3.13.4-insiders", "v3.12.0-rc", targetVersion}

		Convey("Test Filter Results", func() {
			Convey("filterResults returns correct tags", func() {
				filtered := filterResults(tags, ctx)

				So(filtered, ShouldNotBeNil)
				So(filtered, ShouldResemble, []string{targetVersion})

			})

			Convey("stripPrefixAndSuffix munches correctly", func() {
				stripped := stripPrefixAndSuffix([]string{targetVersion}, ctx)

				So(stripped, ShouldNotBeNil)
				So(stripped, ShouldResemble, []string{"v3.11"})
			})

			Convey("processTags evals semver correctly", func() {
				semvers := processTags([]string{"v3.11"}, ctx)

				So(semvers, ShouldNotBeNil)
				So(semvers[0].String(), ShouldEqual, "3.11.0")
			})

			Convey("End to end flow works", func() {
				semvers, err := MunchTags(tags, ctx)
				So(err, ShouldBeNil)
				So(semvers, ShouldNotBeNil)
				So(semvers.String(), ShouldEqual, "3.11.0")
			})

			Convey("End to end flow works with sorting logic", func() {
				tags := append(tags, "linux-arm-v3.12-alpine")
				semvers, err := MunchTags(tags, ctx)
				So(err, ShouldBeNil)
				So(semvers, ShouldNotBeNil)
				So(semvers.String(), ShouldEqual, "3.12.0")
			})
		})
	})
}
