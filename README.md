# covfmt

Utility for converting the go test coverage tools output into the lcov format.

## Dev Setup

    git clone git@github.com:ricallinson/covfmt.git
    cd ./covfmt
    export GOPATH=<CURRENT DIR>
    export PATH=$GOPATH/bin:$PATH

## Usage

    covfmt -f ./coverage/coverage.out -o ./coverage/lcov.info

## Testing

    go test -coverprofile=./coverage/coverage.out
