PROGRAM=app
ARCH=arm64
# .DEFAULT_GOAL := help

# ifneq (,$(wildcard ./.env))
#     include .env
#     export
# endif

.PHONY: serve
serve:
	go run main.go serve

.PHONY: build
build: ## build application
	# go build -v -o $(PROGRAM)
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) go build "-ldflags=-s -w -buildid=" -trimpath -o $(PROGRAM)

.PHONY: lint
lint:
	@# golangci-lint
	govulncheck ./...
	@# govulncheck -show verbose ./...
