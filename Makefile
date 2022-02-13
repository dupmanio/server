# This file is part of the dupman/server project.
#
# (c) 2022. dupman <info@dupman.cloud>
#
# For the full copyright and license information, please view the LICENSE
# file that was distributed with this source code.
#
# Written by Temuri Takalandze <me@abgeo.dev>, February 2022

default: help

## help			:	Print commands help.
.PHONY: help
help : Makefile
	@sed -n 's/^##//p' $<

## download-tools		:	Download necessary development tools.
.PHONY: download-tools
download-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44
	go install github.com/go-swagger/go-swagger/cmd/swagger@v0.29

## test			:	Run Go tests.
.PHONY: test
test:
	go test ./...

## test-race		:	Run Go tests with -race flag.
.PHONY: test-race
test-race:
	go test -race ./...

## test-coverage		:	Run Go tests in coverage mode.
.PHONY: test-coverage
test-coverage:
	go test ./... -coverprofile=coverage.out -covermode=atomic

## coverage		:	Run Go tests and open display coverage.
.PHONY: coverage
coverage: test-coverage
	go tool cover -html=coverage.out

## lint			:	Run linter.
.PHONY: lint
lint:
	golangci-lint run

## update-swagger		:	Update Swagger file.
.PHONY: update-swagger
update-swagger:
	swagger generate spec --scan-models -o docs/swagger.yml
