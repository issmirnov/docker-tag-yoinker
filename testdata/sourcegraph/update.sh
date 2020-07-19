#!/usr/bin/env bash
rm tags.json
wget https://registry.hub.docker.com/v1/repositories/sourcegraph/server/tags -O tags.json