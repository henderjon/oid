BIN=uid
CMD=cmd
BINS=release
HEAD=$(shell git describe --dirty --long --tags 2> /dev/null  || git rev-parse --short HEAD)
TIMESTAMP=$(shell date '+%Y-%m-%dT%H:%M:%S %z %Z')
DEPLOYMENT_PATH=apps/$(BIN)-$(HEAD)

LDFLAGS="-X 'main.buildVersion=$(HEAD)' -X 'main.buildTimestamp=$(TIMESTAMP)' -X 'main.compiledBy=$(shell go version)'" # `-s -w` removes some debugging info that might not be necessary in production (smaller binaries)

all: build

################################################################################
#### HOUSE CLEANING
################################################################################

.PHONY: dep
dep:
	go mod vendor

.PHONY: clean
clean:
	rm -f $(BIN) $(BINS)/$(BIN)

################################################################################
#### ENV BUILDS
################################################################################

.PHONY: build
build: clean
	go build -ldflags $(LDFLAGS) -o $(BIN) ./$(CMD)/$(BIN)

.PHONY: install
install:
	go install -ldflags $(LDFLAGS) ./$(CMD)/$(BIN)
