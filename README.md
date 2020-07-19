# Docker Tag Yoinker

[![Travis](https://travis-ci.com/issmirnov/docker-tag-yoinker.svg?branch=master)](https://travis-ci.com/issmirnov/docker-tag-yoinker)
[![Release](https://img.shields.io/github/release/issmirnov/docker-tag-yoinker.svg?style=flat-square)](https://github.com/issmirnov/docker-tag-yoinker/releases/latest)
![Total Downloads](https://img.shields.io/github/downloads/issmirnov/docker-tag-yoinker/total.svg)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/issmirnov/docker-tag-yoinker?style=flat-square)](https://goreportcard.com/report/github.com/issmirnov/docker-tag-yoinker)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

Sometimes, you just want to yoink the latest tag of a docker image, but `:latest` is too Basic for you.

`dty` (docker-tag-yoinker) is here to help.

Simply set up a config file, and the binary will auto-yoink the latest tag that satisfies your rules.

## Example

[TOML](https://github.com/toml-lang/toml) Config:

```
image = "sourcegraph/server"
blacklist = ["insiders", "rc"]
```

This will pull the latest tag from https://hub.docker.com/r/sourcegraph/server/tags that doesn't have "insiders" or "rc" in it and print it to STDOUT.


## Usage in shell scripts.

```
export TAG=$(dty -config sourcegraph.toml)
docker run --name sourcegraph \
      --restart=always -d --publish 8220:7080 --publish 2633:2633 \
      --volume ~/.sourcegraph/config:/etc/sourcegraph  \
      --volume ~/.sourcegraph/data:/var/opt/sourcegraph  \
      sourcegraph/server:$TAG \
```

This can be run inside a cron job, wired up to https://healthchecks.io/ or equivalent.
