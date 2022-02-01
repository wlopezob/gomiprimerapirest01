  # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINARY_NAME=shopping_service
GOMAIN=main.go
    
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	$(GORUN) $(GOMAIN)