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
To run unit tests:
```
make test
```

To run unit and integration tests
```
export KALEIDO_API=https://xxxx
export KALEIDO_API_KEY=xxxxx
make all_tests
```

Optionally use `go test -v ./kaleido` to view
all test logs.

# Define a consortium using a configuration file

A consortium, can be defined with a YAML file. The following configuration will create a consoritum, an environment with 2 members where each member has 1 node.

```
kld create consortium -c /path/to/consortium.yaml
```


```
---
provider: quorum
consensus: raft
waitok: true
consortium:
  mode: single-org
  description: A description
  name: test-consortium
  environments:
  - name: dev
    description: environment description
    members:
    - name: org1
      nodes:
      - name: org1-node1
    - name: org2
      nodes: 
      - name: org2-node1
```

If *waitok: true* is set, it will wait until the nodes are ready so that the enpoints can be output. This is a sample output: 
```
{
	"consortium_id": "e0olj47vdk",
	"environments": [{
		"id": "e0wy1qwc1z",
		"members": [{
			"appcreds": {
				"id": "",
				"password": "TdVJJzmpKDkZ8_wW-pNfhMASiX_l9Ouha-juhDXk_OQ",
				"username": "e0dvjcg7kt"
			},
			"id": "e0wfnc3of3",
			"name": "org1",
			"nodes": [{
				"block_height": 0,
				"geth": {
					"public_address": "",
					"validators": null
				},
				"quorum": {
					"private_address": "L8QZW8+Z7eBFVTxgmGJ/P91X/NFxRtidFQSaqqJ9LEk=",
					"public_address": "0xcdb6a9185df7c1d3a9ffeb8e1e5f32f2a2569893"
				},
				"id": "8f30392ff18d681642560b123cd0b93c44f1ba8ac4830b23b54703e823e47e8dd79f822acbee7f0db6d81f7730776fcec7d62fdedd3753ff50a021e58a1067e4",
				"urls": {
					"rpc": "https://e0wy1qwc1z-e0q9a9cryg-rpc.eu-central-1.kaleido.io",
					"wss": "wss://e0wy1qwc1z-e0q9a9cryg-wss.eu-central-1.kaleido.io"
				},
				"user_accounts": ["0x54Fe04De0E46F5C1EEF023781d0a809F58228Aa1"]
			}]
		}, {
			"appcreds": {
				"id": "",
				"password": "ybP85DIK9mFuclVU0_znBkcmH7VEBRswQgT4mJbNInw",
				"username": "e0axwx9q11"
			},
			"id": "e0rsbvcgdi",
			"name": "org2",
			"nodes": [{
				"block_height": 0,
				"geth": {
					"public_address": "",
					"validators": null
				},
				"quorum": {
					"private_address": "18qbrglhneGmAGhcRNDBh+Q74Gjuuh+QlKctdMfQ5lE=",
					"public_address": "0x6e26c0a4f5ce887a74d84d4a848081d900ffc364"
				},
				"id": "7e2b4838687f2e46d42748bd1bb0fc8a869c16b2969944532686f889d3b3bc3047393d936e1f9ac4f56455da081b7b2502de411923e3dd2bc3708ba2582e0ab3",
				"urls": {
					"rpc": "https://e0wy1qwc1z-e0iyv0yy28-rpc.eu-central-1.kaleido.io",
					"wss": "wss://e0wy1qwc1z-e0iyv0yy28-wss.eu-central-1.kaleido.io"
				},
				"user_accounts": ["0x15C5D9bbc349eD0B0FfD5f056D74bA6ECd5E0ba9"]
			}]
		}]
	}]
}
```