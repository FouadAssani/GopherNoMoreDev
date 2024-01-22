BINDIR=./bin
BASE_API_FOLDER := cmd
DOCKER_BUILDKIT ?= 0

API_FOLDER := api
API_PORT ?= 8080
API_PUBLIC_URL ?= http://localhost:$(API_PORT)

.PHONY: build
build: build-api

.PHONY: build-api
build-api:
	cd $(BASE_API_FOLDER)/api/ && $(MAKE) build --no-print-directory
	cd $(BASE_API_FOLDER)/agents/ && $(MAKE) build --no-print-directory

.PHONY: test
test:
	go clean -testcache
	go test -count=1 ./cmd/... ./internal/... -coverprofile cover.out

.PHONY: test-integration
test-integration:
	go clean -testcache
	go test -timeout=60m -count=1 -p 1 ./test/integration/...

.PHONY: test-e2e
test-e2e:
	go clean -testcache
	go test -timeout=60m -count=1 -p 1 ./test/e2e/...

.PHONY: download
download:
	go mod download

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: format
format:
	golangci-lint run --config .golangci.yml --fix

.PHONY: lint
lint:
	golangci-lint run --timeout 15m0s --config .golangci.yml

.PHONY: deps
deps: download deps-lint deps-test deps-openapi

.PHONY: deps-lint
deps-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.55.2

.PHONY: deps-test
deps-test:
	go install github.com/golang/mock/mockgen@v1.6.0

.PHONY: deps-openapi
deps-openapi:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4

.PHONY: generate
generate:
	go generate -run golang.org/x/tools/cmd/stringer ./...
	go generate ./...

.PHONY: openapi
openapi:
	@cd cmd/$(API_FOLDER) && $(MAKE) openapi --no-print-directory

.PHONY: init
init: deps generate openapi

.PHONY: clean
clean:
	go clean -testcache
	rm -rf ./bin
	rm -f ./cover.out
	find . -type d -name '*mocks' -exec rm -rf {} +
	find . -type f -name '*_string.go' -exec rm {} +
	find . -type f -name '*.gen.go' -exec rm {} +

.PHONY: docker-build
docker-build:
	DOCKER_BUILDKIT=1 docker build -t gophernomoredev .

.PHONY: docker-compose-run
docker-compose-run:
	docker-compose up --quiet-pull -d gophernomoredev-api