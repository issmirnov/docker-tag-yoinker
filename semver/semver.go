package semver

import (
	"sort"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/davecgh/go-spew/spew"
	"github.com/issmirnov/docker-updater/filters"
	"github.com/issmirnov/docker-updater/interfaces"
)

func filterResults(tags []string, ctx interfaces.Context) (res []string) {

	var blacklistFilter filters.Filter = func(s string) bool {
		for _, b := range ctx.Config.Blacklist {
			if strings.Contains(s, b) {
				return false
			}
		}
		return true
	}

	var whitelistFilter filters.Filter = func(s string) bool {
		return strings.Contains(s, ctx.Config.Whitelist)
	}

	var prefixFilter filters.Filter = func(s string) bool {
		return strings.HasPrefix(s, ctx.Config.StripPrefix)
	}

	var suffixFilter filters.Filter = func(s string) bool {
		return strings.HasSuffix(s, ctx.Config.StripSuffix)
	}

	fs := []filters.Filter{blacklistFilter,
		whitelistFilter,
		prefixFilter,
		suffixFilter}

	res = filters.ApplyFilters(tags, fs...)
	return
}

func stripPrefixAndSuffix(tags []string, ctx interfaces.Context) (res []string) {
	for _, tag := range tags {
		tag = strings.TrimPrefix(tag, ctx.Config.StripPrefix)
		tag = strings.TrimSuffix(tag, ctx.Config.StripSuffix)
		res = append(res, tag)
	}
	return
}

func processTags(tags []string, ctx interfaces.Context) (res []*semver.Version) {
	for _, tag := range tags {
		v, err := semver.NewVersion(tag)
		if err != nil {
			ctx.Logger.Debug().Msg(err.Error())
		} else {
			res = append(res, v)
		}
	}

	return
}

func MunchTags(tags []string, ctx interfaces.Context) *semver.Version {
	filtered := filterResults(tags, ctx)
	stripped := stripPrefixAndSuffix(filtered, ctx)
	parsed := processTags(stripped, ctx)
	sort.Sort(semver.Collection(parsed)) // Note: sort is ascending.
	ctx.Logger.Debug().Msg(spew.Sprint(parsed))
	return parsed[len(parsed)-1] // return last element

}
