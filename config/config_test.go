package config

import (
	"testing"

	"github.com/issmirnov/docker-updater/internal"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidConfig(t *testing.T) {

	// Send testing output to /dev/null
	// setupLogging(ioutil.Discard, false)

	Convey("Setup", t, func() {

		c := internal.LoadValidTestConfig()
		blacklistArray := []string{"insiders", "rc"}

		Convey("Verify TOML parses correctly.", func() {
			So(c.Blacklist, ShouldResemble, blacklistArray)
			So(c.Whitelist, ShouldEqual, "alpine")
			So(c.StripPrefix, ShouldEqual, "linux-arm-")
			So(c.StripSuffix, ShouldEqual, "-alpine")
		})
	})

	Convey("Load sample config from repo", t, func() {

		Convey("Verify TOML parses correctly.", func() {
			c, err := LoadConfig("../config.toml")
			blacklistArray := []string{"insiders", "rc"}

			So(err, ShouldBeNil)
			So(c.Blacklist, ShouldResemble, blacklistArray)
			So(c.Whitelist, ShouldEqual, "alpine")
			So(c.StripPrefix, ShouldEqual, "linux-arm-")
			So(c.StripSuffix, ShouldEqual, "-alpine")
		})

		Convey("Loading non existent file should fail.", func() {
			_, err := LoadConfig("../config-does-not-exist.toml")
			So(err, ShouldNotBeNil)
		})

	})
}
