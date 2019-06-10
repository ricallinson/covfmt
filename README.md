# covfmt

[![Build Status](https://travis-ci.org/ricallinson/covfmt.svg?branch=master)](https://travis-ci.org/ricallinson/covfmt)

_WORK IN PROGRESS_: Utility for converting the go test coverage output into the lcov format.

## Usage

    go test -coverprofile=coverage.out; cat coverage.out | covfmt > ./lcov.info

    go test -coverprofile=coverage.out; cat coverage.out | covfmt $(pwd) > ./lcov.info

## Testing

    go test -cover

## Coverage

    go test -coverprofile=coverage.out; go tool cover -html=coverage.out -o=index.html