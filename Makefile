# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY=tsqlsh

default: build
dep: 
	dep ensure 

all: test build
build: build-linux
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -rf bin/*
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)   

# Cross compilation
build-linux: clean dep
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/linux/$(BINARY) -v

build-osx: clean dep
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/osx/$(BINARY) -v

build-win: clean dep
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/win/$(BINARY).exe -v