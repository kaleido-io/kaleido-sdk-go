EXE = kld-test
PACKAGE = github.com/kaleido-io/kaleido-sdk-go
BASE 		= $(GOPATH)/src/$(PACKAGE)
VERSION = 0.1.0

.PHONY : all clean fmt test test-junit build

all : clean fmt test build

build : test
	@GOOS=darwin GOARCH=amd64 go build  -o $(GOPATH)/bin/$(EXE)-$(VERSION)-osx $(PACKAGE)
	@GOOS=linux GOARCH=amd64 go build  -o $(GOPATH)/bin/$(EXE)-$(VERSION)-linux $(PACKAGE)

test : fmt
	@go test -v -cover ./...

all_tests : fmt
	@go test -v -tags=integration -cover ./...
	
fmt :
	@gofmt -w $(BASE)/*.go

clean :
	@-rm $(GOPATH)/bin/$(EXE)-*
	@-rm $(GOPATH)/bin/rice
