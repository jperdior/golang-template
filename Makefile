# Makefile for golang-template
# vim: set ft=make ts=8 noet
# Licence MIT

# Variables
# UNAME		:= $(shell uname -s)
PWD = $(shell pwd)

.EXPORT_ALL_VARIABLES:

# this is godly
# https://news.ycombinator.com/item?id=11939200
.PHONY: help
help:
ifeq ($(UNAME), Linux)
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
else
	@# this is not tested, but prepared in advance for you, Mac drivers
	@awk -F ':.*###' '$$0 ~ FS {printf "%15s%s\n", $$1 ":", $$2}' \
		$(MAKEFILE_LIST) | grep -v '@awk' | sort
endif

tests:
	go test -v ./...

start: initialize build run

restart : stop start

initialize:
	./ops/scripts/initialize.sh

build:
	docker-compose -f ops/docker/docker-compose.yml build

run:
	docker-compose -f ops/docker/docker-compose.yml up -d

stop:
	docker-compose -f ops/docker/docker-compose.yml down

analysis: ### Run static analysis and linter
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:latest golangci-lint run

openapi-init:
	swag init -g cmd/api/main.go -g internal/platform/server/routes.go -o docs
