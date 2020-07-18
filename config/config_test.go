package config

import (
	"testing"

	"github.com/issmirnov/docker-updater/internal"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidConfig(t *testing.T) {

	Convey("Setup", t, func() {
		// Send testing output to /dev/null
		// setupLogging(ioutil.Discard, false)

		c := internal.LoadValidTestConfig()
		blacklistArray := []string{"insiders", "rc"}

		Convey("Verify TOML parses correctly.", func() {
			So(c.Blacklist, ShouldResemble, blacklistArray)
			So(c.Whitelist, ShouldEqual, "alpine")
			So(c.StripPrefix, ShouldEqual, "linux-arm-")
			So(c.StripSuffix, ShouldEqual, "-alpine")
		})
	})
}
