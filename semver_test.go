package main

import (
	"testing"

	"github.com/issmirnov/docker-updater/internal"
	. "github.com/smartystreets/goconvey/convey"
)

const targetVersion = "linux-arm-v3.11-alpine"

func TestSemver(t *testing.T) {

	Convey("Setup", t, func() {
		// Send testing output to /dev/null
		// setupLogging(ioutil.Discard, false)
		Context.Config = internal.LoadValidTestConfig()
		tags := []string{"v3.13.4-insiders", "v3.12.0-rc", targetVersion}

		Convey("Test Filter Results", func() {
			Convey("filterResults returns correct tags", func() {
				filtered := filterResults(tags, Context)

				So(filtered, ShouldNotBeNil)
				So(filtered, ShouldResemble, []string{targetVersion})

			})

			Convey("stripPrefixAndSuffix munches correctly", func() {
				stripped := stripPrefixAndSuffix([]string{targetVersion}, Context)

				So(stripped, ShouldNotBeNil)
				So(stripped, ShouldResemble, []string{"v3.11"})
			})

			Convey("processTags evals semver correctly", func() {
				semvers := processTags([]string{"v3.11"})

				So(semvers, ShouldNotBeNil)
				So(semvers[0].String(), ShouldEqual, "3.11.0")
			})

			Convey("End to end flow works", func() {
				semvers := MunchTags(tags, Context)

				So(semvers, ShouldNotBeNil)
				So(semvers.String(), ShouldEqual, "3.11.0")
			})

			Convey("End to end flow works with sorting logic", func() {
				tags := append(tags, "linux-arm-v3.12-alpine")
				semvers := MunchTags(tags, Context)

				So(semvers, ShouldNotBeNil)
				So(semvers.String(), ShouldEqual, "3.12.0")
			})
		})
	})
}
