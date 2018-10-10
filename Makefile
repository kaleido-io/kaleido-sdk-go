 # Go parameters
XGOCMD=xgo

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
DEPCMD=dep
DEPENSURE=$(DEPCMD) ensure
BINARY_NAME=kld

LDFLAGS="-X main.buildDate=`date -u +\"%Y-%m-%dT%H:%M:%SZ\"` -X main.buildVersion=$(BUILD_VERSION)"
DEPS=https://gmplib.org/download/gmp/gmp-6.0.0a.tar.bz2
TARGETS="windows-10.0/*,darwin-10.10/*"

all: deps build test
build: 
		$(XGOCMD) -ldflags=$(LDFLAGS) -tags=prod --deps=$(DEPS) --out=$(BINARY_NAME)-$(BUILD_VERSION) --targets=$(TARGETS) -v -x .
test:
		$(GOTEST)  ./... -cover -coverprofile=coverage.txt -covermode=atomic
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)-$(BUILD_VERSION)*
deps:
	$(GOGET) github.com/karalabe/xgo
	$(DEPENSURE)

build-linux:
build-mac:
		GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_MAC) -v
build-win:
		