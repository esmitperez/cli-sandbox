# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=esmit-cli-sandbox
BINARY_WIN=$(BINARY_NAME).exe
BINARY_LINUX=$(BINARY_NAME)

all: test build

# by default, build for Mac OS :)
build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v

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
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_WIN) -v
