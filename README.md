# Command Line tool for managing Kaleido resources

This tool allows Kaleido users to provision and manager the resources such as consortiums, memberships, environments and Quorum/Geth nodes.

## Dependencies

Install Go: [https://golang.org/doc/install](https://golang.org/doc/install)

The project uses a GO dependency management tool called "dep".

To install:
```
brew install dep
```

## Dev Environment Setup

The follow are regular steps to set up a typical golang dev environment. If you are already familiar with this language you can skip to the next section.

1. decide a folder as your container for all golang projects, and set `GOPATH` environment variable to point to it
2. create a folder `$GOPATH/src/github.com/consensys` and `cd` to it
3. clone the repository in this folder, or if you have already cloned it somewhere else, move the project to this folder

You should have the following directory holding the project source:
```
$GOPATH/src/github.com/consensys/photic-sdk-go
```

## Install Dependencies

Use `dep` to install the dependent golang packages, from the root of the project:
```
cd $GOPATH/src/github.com/consensys/photic-sdk-go
dep ensure
```

## Running Tests

```
export KALEIDO_API=https://xxxx
export KALEIDO_API_KEY=xxxxx
go test ./kaleido
```

Optionally use `go test -v ./kaleido` to view
all test logs.
