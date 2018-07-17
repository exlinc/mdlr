# Mdlr Tool for Git Dependencies

Declarative and easy git 'submodules without the pain of submodules.'

## In action

[![asciicast](https://asciinema.org/a/192116.png)](https://asciinema.org/a/192116)

## Installing

### Linux / OS X

To install for linux or OS X systems, please run `curl https://s3-us-west-2.amazonaws.com/mdlr-dist/v1/install.sh | bash`.

### Dockerfile-compatible Automated Installer (assumes *nix system, root, non-interactive)

For docker-ized installs or systems where you are in the root shell, please run `curl https://s3-us-west-2.amazonaws.com/mdlr-dist/v1/install-root.sh | bash`

## Windows (Beta support)

Download the `.exe` binary for your platform:

- [Windows x64](https://s3-us-west-2.amazonaws.com/mdlr-dist/v1/windows-amd64-mdlr.exe)

- [Windows x86](https://s3-us-west-2.amazonaws.com/mdlr-dist/v1/windows-386-mdlr.exe)

From there, you can put the `.exe` file where you like and run it from the command line/powershell.

If you have any issues, with the Windows install/functionality, please report them [here]().

## Updating

Please run the install script for your platform above to get the latest updates. For windows users, replace the `.exe` binary

## Using Mdlr

### Example workflow

1.  `cd project/directory/ # Enter the project directory`
2.  `mdlr init`
3.  `mdlr add --name depname --path deps/mydep --url https://github/org/mydep.git`
4.  `mdlr import -f # Reset the module forcefully (wipe changes, if any) and then import it at the version in the mdlr.yml file. This is the recommended comand`
5.  `mdlr list # List the modules`
6.  `mdlr status # Get the status overview`
7.  `mdlr update -f # Reset the module forcefully (wipe changes, if any) and then update it and write the new update to the mdlr.yml file`
8.  `vim mdlr.yml # View/edit the mdlr.yml file`

### Create a new mdlr project

In the project directory, run `mdlr init`

### Import modules for a mdlr project

In the project directory, run `mdlr import -f`

### Commands overview

- `help`: get a help overview
- `init`: generate a mdlr.yml file in the directory
- `list`: list the current modules
- `add`: add a module to the mdlr.yml file
- `remove`: remove a module
- `import`: import a module
- `update`: update a module
- `status`: get the status for the mdlr.yml or invidual modules

## Installing from Source / Developing

### Prerequisites

1. Ubuntu 16.04 (Server or Desktop) operating system -- other similar systems and OS X might work, but aren't guaranteed to...
2. GoLang 1.9.^ installed
3. GoLang dependency manager installed [(Install guide)](https://github.com/golang/dep#setup)

### Get the project

1. `go get github.com/exlinc/mdlr`

### Setting up `dep`

1. `cd $GOPATH/src/github.com/exlinc/mdlr`
2. `dep ensure -v # Set to verbose to track the progress as this might take a while...`

### Build the code (for dev platform)

1. `go build # Optionally, use go run main.go instead of the build+run flow and your code will compile every time you run with the latest changes`

### Run the code

1. `./mdlr # Runs the latest output of go build` OR `go run main.go # Compiles a temp binary from the latest changes and runs it all in one command`

### Install the dev binary on your system

1. Run `go install` which (if your `$GOPATH/bin` is in your `$PATH`) will create a globally-accessible `tools-mdlr` binary that you can use to easily test a 'dev' version of your code anywhere on your dev system with `tools-mdlr`

## Distribution

### Build the distribution binaries

1. `./xplatform-build.sh`
2. See the output binaries in the `./build` directory!

### Uploading to S3 for distribution

The S3 upload is done by the maintainer (EXL Inc.) for new releases to the bucket used in the install scripts.

## Contributing

Contributions are welcome and appreciated in the form of issues and pull requests in this repo!

## License

[MIT](LICENSE)
