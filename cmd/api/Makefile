BINDIR = ../../bin
APPNAME ?= gophernomoredev_api
DEST = $(BINDIR)/$(APPNAME)

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags "-X main.version=$$VERSION" -o $(DEST) *.go

.PHONY: run
run: build
	@$(DEST)

.PHONY: openapi
openapi:
	@mkdir -p restapi/openapi/openapiclient
	oapi-codegen --config oapi-codegen.yaml "restapi/openapi.yml"
	oapi-codegen --config oapi-codegen-client.yaml "restapi/openapi.yml"

.PHONY: clean
clean:
	find . -type d -name '*mocks' -exec rm -rf {} +