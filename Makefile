################################################################################
#### INSTALLATION VARS
################################################################################
PREFIX=$(HOME)

################################################################################
#### BUILD VARS
################################################################################
BIN=oid
BINDIR=bin
CMDDIR=cmd
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
	rm -f $(BIN) $(BIN)-* $(BINDIR)/$(BIN) $(BINDIR)/$(BIN)-*

################################################################################
#### INSTALL
################################################################################

.PHONY: install
install: local
	mkdir -p $(PREFIX)/$(BINDIR)
	mv $(BINDIR)/$(BIN) $(PREFIX)/$(BINDIR)/$(BIN)
	@echo "\ninstalled $(BIN) to $(PREFIX)/$(BINDIR)\n"


.PHONY: uninstall
uninstall:
	rm -f $(PREFIX)/$(BINDIR)/$(BIN)

################################################################################
#### ENV BUILDS
################################################################################

.PHONY: local
local: clean
	go build -ldflags $(LDFLAGS) -o $(BINDIR)/$(BIN) ./$(CMDDIR)/$(BIN)



