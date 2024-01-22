BINDIR = ../../bin
APPNAME ?= agents
DEST = $(BINDIR)/$(APPNAME)

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags "-X main.version=$$VERSION" -o $(DEST) *.go

.PHONY: run
run: build
	@$(DEST)

.PHONY: clean
clean:
	find . -type d -name '*mocks' -exec rm -rf {} +