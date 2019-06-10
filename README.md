# covfmt

[![Build Status](https://travis-ci.org/ricallinson/covfmt.svg?branch=master)](https://travis-ci.org/ricallinson/covfmt) [![Coverage Status](https://coveralls.io/repos/github/ricallinson/covfmt/badge.svg)](https://coveralls.io/github/ricallinson/covfmt)

Utility for converting the go test coverage output into the lcov format.

## Usage

    go test -coverprofile=coverage.out; cat coverage.out | covfmt > ./lcov.info

    go test -coverprofile=coverage.out; cat coverage.out | covfmt -prefix $(pwd) > ./lcov.info

    go test -coverprofile=coverage.out; cat coverage.out | covfmt -trim github.com > ./lcov.info

## Testing

    go test -cover

## Coverage

    go test -coverprofile=coverage.out; go tool cover -html=coverage.out -o=index.html