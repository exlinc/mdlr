### Prerequisites

1. Ubuntu 16.04 (Server or Desktop) operating system -- other similar systems and OS X might work, but aren't guaranteed to...
2. GoLang 1.9.^ installed
3. GoLang dependency manager installed [(Install guide)](https://github.com/golang/dep#setup)

### Get the project

1. `go get git.exlhub.io/exlinc/tools-mdlr`

### Setting up `dep`

1. `cd $GOPATH/src/git.exlhub.io/exlinc/tools-mdlr`
2. `dep ensure -v # Set to verbose to track the progress as this might take a while...`

### Build the code (for dev platform)

1. `go build # Optionally, replace tools-mdlr with go run main.go and your code will compile every time you run with the hottest changes`

### Run the code

1. `./tools-mdlr # Runs the latest output of go build` OR `go run main.go # Compiles a temp binary from the latest changes and runs it all in one command`

### Install the dev binary on your system

1. Run `go install` which (if your `$GOPATH/bin` is in your `$PATH`) will create a globally-accessible `tools-mdlr` binary that you can use to easily test a 'dev' version of your code anywhere on your dev system with `tools-mdlr`

### Build the distribution binaries

1. `./xplatform-build.sh`
2. See the output binaries in the `./build` directory!