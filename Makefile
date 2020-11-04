################################################################################
#### INSTALLATION VARS
################################################################################
PREFIX=/usr/local

################################################################################
#### BUILD VARS
################################################################################
BIN=oid
BINS_DIR=bin
CMD_DIR=cmd
HEAD=$(shell git describe --dirty --long --tags 2> /dev/null  || git rev-parse --short HEAD)
TIMESTAMP=$(shell TZ=UTC date '+%FT%T %Z')
DEPLOYMENT_PATH=apps/$(BIN)-$(HEAD)
SSH_ALIAS=

LDFLAGS="-X 'main.buildVersion=$(HEAD)' -X 'main.buildTimestamp=$(TIMESTAMP)' -X 'main.compiledBy=$(shell go version)'" # `-s -w` removes some debugging info that might not be necessary in production (smaller binaries)

all: local

################################################################################
#### HOUSE CLEANING
################################################################################

.PHONY: dep
dep:
	go mod vendor

.PHONY: check
check:
	golint
	goimports -w ./
	gofmt -w ./
	go vet

.PHONY: clean
clean:
	rm -f $(BIN) $(BIN)-* $(BINS_DIR)/$(BIN) $(BINS_DIR)/$(BIN)-*

################################################################################
#### INSTALL
################################################################################

.PHONY: install
install:
	mkdir -p $(PREFIX)/bin
	cp $(BINS_DIR)/$(BIN) $(PREFIX)/bin/$(BIN)

.PHONY: uninstall
uninstall:
	rm -f $(PREFIX)/bin/$(BIN)

################################################################################
#### ENV BUILDS
################################################################################

.PHONY: local
local: clean
	go build -ldflags $(LDFLAGS) -o $(BINS_DIR)/$(BIN) ./$(CMD_DIR)/$(BIN)



