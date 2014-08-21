# lcov

## Dev Setup

    git clone git@github.com:ricallinson/covfmt.git
    cd ./covfmt
    export GOPATH=<CURRENT DIR>
    export PATH=$GOPATH/bin:$PATH

## Usage

    covfmt ./coverage/coverage.out

## Testing

    go test -coverprofile=./coverage/coverage.out
