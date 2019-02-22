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
2. create a folder `$GOPATH/src/github.com/kaleido-io` and `cd` to it
3. clone the repository in this folder, or if you have already cloned it somewhere else, move the project to this folder

You should have the following directory holding the project source:
```
$GOPATH/src/github.com/kaleido-io/kaleido-sdk-go
```

## Install Dependencies

Use `dep` to install the dependent golang packages, from the root of the project:
```
cd $GOPATH/src/github.com/kaleido-io/kaleido-sdk-go
dep ensure
```

## Build and Run

```
go build -o kld
```

The `kld` command works similarly to the popular CLI tools like `kubectl` which multiple levels of commands with help. Try these commands:
```
./kld
./kld -h
./kld create -h
./kld create consortium -h
```

To use the CLI to do something, such as create a consortium, you first need to provide the two pieces of required information, API URL and API Key. The tool supports configurations through files (yaml, json, toml formats are all supported), or environment variables.

### Configuration in files and environment variables

Create a `.kld.yaml` (or .kld.json or .kld.toml based on your preferences) file in the home directory and give these values:
```
api:
  url: https://console.kaleido.io/api/v1
  key: <your API Key>
```
To establish your organizational identity in the registry service, you can run the following command to easily configure your yaml file. You may have to run ```chmod +x configureYaml``` before you run the following command:
```
./configureYaml "https://console.kaleido.io/api/v1" "<your API key>" "<your on-chain registry service ID(last portion of the dashboard url)>" "<a fully-qualified node endpoint>"
```

The configuration file can be at any location and called by any name, you can tell the command where it is:
```
./kld --config /path/to/your/config.toml create consortium ...
```

You can also use environment variables. Use all capital letters with `KLD_` as the prefix, for instance:
```
KLD_API_URL=http://console.kaleido.io/api/v1
KLD_API_KEY=blahblah
```

If the same configurations are specified in multiple places, this is the precedence order:
> **environment variables** overrides **config files specified in --config** overrides **.kld.<yaml/json/toml> in home directory**

### Create a consortium

```
jimzhang$ ./kld create consortium -m single-org -n testConsortium234 -d "this is a test consortium" |jq
{
  "name": "testConsortium234",
  "description": "this is a test consortium",
  "mode": "single-org",
  "owner": "z4a5lhio",
  "_id": "xmvsgnlg",
  "state": "setup",
  "_revision": "0",
  "created_at": "2018-04-13T15:14:20.930Z",
  "environment": {
    "name": "Auto-generated Environment",
    "description": "Initial Environment for testConsortium234",
    "state": "live",
    "region": "us-east",
    "provider": "quorum",
    "consensus_type": "raft",
    "mode": "single-org",
    "_id": "wgv4mdl4",
    "node_list": [],
    "release_id": "i05cs743",
    "_revision": "0",
    "created_at": "2018-04-13T15:14:20.956Z"
  }
}
```

## Running Tests

```
export KALEIDO_API=https://xxxx
export KALEIDO_API_KEY=xxxxx
go test ./kaleido
```

Optionally use `go test -v ./kaleido` to view
all test logs.
