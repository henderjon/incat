################################################################################
#### INSTALLATION VARS
################################################################################
PREFIX=/usr/local

################################################################################
#### BUILD VARS
################################################################################
BIN=incat
BINS=.
HEAD=$(shell git describe --dirty --long --tags 2> /dev/null  || git rev-parse --short HEAD)
TIMESTAMP=$(shell TZ=UTC date '+%FT%T %Z')

LDFLAGS="-X 'main.buildVersion=$(HEAD)' -X 'main.buildTimestamp=$(TIMESTAMP)' -X 'main.compiledBy=$(shell go version)'" # `-s -w` removes some debugging info that might not be necessary in production (smaller binaries)

all: local

################################################################################
#### HOUSE CLEANING
################################################################################

clean:
	rm -f $(BIN) $(BIN)-* $(BINS)/$(BIN) $(BINS)/$(BIN)-*

.PHONY: check
check:
	golint
	goimports -w ./
	gofmt -w ./
	go vet

################################################################################
#### INSTALL
################################################################################

.PHONY: install
install:
	mkdir -p $(PREFIX)/bin
	cp $(BINS)/$(BIN) $(PREFIX)/bin/$(BIN)

.PHONY: uninstall
uninstall:
	rm -f $(PREFIX)/bin/$(BIN)

################################################################################
#### ENV BUILDS
################################################################################

.PHONY: local
local: check
	go build -ldflags $(LDFLAGS) -o $(BINS)/$(BIN)

.PHONY: localv
localv: check
	go build -mod=vendor -ldflags $(LDFLAGS) -o $(BINS)/$(BIN)

.PHONY: test
test:
	go mod vendor
	go mod tidy
	go test -mod=vendor -coverprofile=coverage.out -covermode=count ./...
