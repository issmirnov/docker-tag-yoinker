TARGET=docker-tag-yoinker

CGO_CFLAGS=-I/usr/local/include
CGO_LDFLAGS=-L/usr/local/lib


# count processors
NPROCS:=1
OS:=$(shell uname -s)
ifeq ($(OS),Linux)
	NPROCS:=$(shell grep -c ^processor /proc/cpuinfo)
endif
ifeq ($(OS),Darwin) # Assume Mac OS X
	NPROCS:=$(shell sysctl -n hw.ncpu)
endif

all: debug

## go build
build:
	go build -o "$(TARGET)" .

## debug with hot reload via modd
debug: build
	modd

## install go deps
deps:
	go get -t -v ./...

## run local godoc server on :8080
doc:
	godoc -http=localhost:8080 -index


## Build and run
run: build
	./$(TARGET) -c testdata/sourcegraph/config.toml

## Run go tests
test:
	go test ./... -v

## init project
init: deps build
	@echo Done! Run project with ./$(TARGET) or 'make' or 'make run'


# Fancy help message
# Source: https://gist.github.com/prwhite/8168133
# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
TARGET_MAX_CHAR_NUM=20

## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)