#!/usr/bin/env bash
rm tags.json
wget https://registry.hub.docker.com/v1/repositories/ghost/tags -O tags.json
