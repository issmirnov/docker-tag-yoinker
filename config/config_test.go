package config

import (
	"testing"

	"github.com/issmirnov/docker-updater/interfaces"
	"github.com/issmirnov/docker-updater/internal"
	"github.com/rs/zerolog"
	. "github.com/smartystreets/goconvey/convey"
)

func TestValidConfig(t *testing.T) {

	// Send testing output to /dev/null
	zerolog.SetGlobalLevel(zerolog.FatalLevel)

	Convey("Setup", t, func() {

		c := internal.LoadValidTestConfig()

		Convey("Verify TOML parses correctly.", func() {
			So(c.Blacklist, ShouldResemble, []string{"insiders", "rc"})
			So(c.Whitelist, ShouldEqual, "alpine")
			So(c.StripPrefix, ShouldEqual, "linux-arm-")
			So(c.StripSuffix, ShouldEqual, "-alpine")
		})
	})

	Convey("Load sample config from repo", t, func() {
		ctx := &interfaces.Context{}

		Convey("Verify TOML parses correctly.", func() {
			err := LoadConfig("../config.toml", ctx)
			So(err, ShouldBeNil)

			So(ctx.Config.Blacklist, ShouldResemble, []string{"insiders", "rc"})
			So(ctx.Config.Whitelist, ShouldEqual, "alpine")
			So(ctx.Config.StripPrefix, ShouldEqual, "linux-arm-")
			So(ctx.Config.StripSuffix, ShouldEqual, "-alpine")
		})

		Convey("Loading non existent file should fail.", func() {
			err := LoadConfig("../config-does-not-exist.toml", ctx)
			So(err, ShouldNotBeNil)
		})

	})
}
