# Docker Tag Yoinker

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

