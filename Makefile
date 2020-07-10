# Go parameters
default: build

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY=tsqlsh

build: build-osx

test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -rf bin/*

run: build
		./bin/osx/$(BINARY)

# Cross compilation
build-linux: clean 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/linux/$(BINARY) -v

build-osx: clean 
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/osx/$(BINARY) -v

build-win: clean 
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/win/$(BINARY).exe -v