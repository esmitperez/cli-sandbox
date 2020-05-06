
BINARY_NAME=esmit-cli-sandbox
BINARY_WIN=$(BINARY_NAME).exe
BINARY_LINUX=$(BINARY_NAME)
BINARY_RELEASE=2020.5.6
BINARY_RELEASE_NICKNAME=Margarita 🍸

# Go parameters
GOCMD=go
GOARCH := $(shell go env GOARCH)
GOOS := $(shell go env GOOS)

# See https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
# the path to pass -ldflag -X was determined with `go tool nm ./cli-sandbox | grep Version`
LDFLAGS=-X 'esmit.me/cli-sandbox/cmd.Version=$(BINARY_RELEASE) ($(BINARY_RELEASE_NICKNAME))' \
  -X 'esmit.me/cli-sandbox/cmd.CmdName=$(BINARY_NAME)'
GOBUILD=$(GOCMD) build -ldflags="$(LDFLAGS)"
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get


all: test build

# by default, built for your local arch. Specific architectures below
build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_LINUX)
	rm -f $(BINARY_WIN)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	/$(BINARY_NAME)

deps:
	$(GOGET) github.com/spf13/cobra

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_WIN) -v

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_WIN) -v
