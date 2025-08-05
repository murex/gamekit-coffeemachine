ifeq ($(OS),Windows_NT)
	EXT := ".exe"
else
	EXT := ""
endif

CONFIG_PKG="github.com/murex/gamekit-coffeemachine/settings"
export CGO_ENABLED=0

.PHONY: default
default: help ## Shows this help message (default target)

.PHONY: help
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

# Convenience target for automating release preparation
.PHONY: prepare
prepare: deps tidy lint modernize build test ## Prepare for commit or release - runs deps, tidy, lint, modernize, build, test

.PHONY: deps
deps: ## Update all dependencies
	@go get -u -t tool ./...

.PHONY: lint
lint: ## Run linter (golangci-lint)
	@golangci-lint run -v

.PHONY: modernize
modernize: ## Run go modernize utility
	@go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@latest -test ./...

.PHONY: vet
vet: ## Run go vet
	@go vet ./...

.PHONY: tidy
tidy: ## Tidy go modules
	@go mod tidy

.PHONY: build
build: ## Build all binaries (cli, progress-runner, gotestsum, test2json)

cli_bin := bin/cli$(EXT)
build: $(cli_bin)
$(cli_bin):
	@mkdir -p bin
	@go build -o $@ -ldflags "-s -w \
    	-X ${CONFIG_PKG}.BuildVersion="`git describe --tags`" \
        -X ${CONFIG_PKG}.BuildOs="`go env GOOS`" \
        -X ${CONFIG_PKG}.BuildArch="`go env GOARCH`" \
        -X ${CONFIG_PKG}.BuildCommit="`git rev-list --max-count=1 --tags`" \
        -X ${CONFIG_PKG}.BuildDate="`date -u +%FT%TZ`" \
        -X ${CONFIG_PKG}.BuildAuthor="`id -un`" \
        "

progress_runner_bin := bin/progress-runner$(EXT)
build: $(progress_runner_bin)
$(progress_runner_bin):
	@mkdir -p bin
	@go test -c -o $@ ./progress -ldflags "-s -w \
 	  -X ${CONFIG_PKG}.BuildVersion="`git describe --tags`" \
      -X ${CONFIG_PKG}.BuildOs="`go env GOOS`" \
      -X ${CONFIG_PKG}.BuildArch="`go env GOARCH`" \
      -X ${CONFIG_PKG}.BuildCommit="`git rev-list --max-count=1 --tags`" \
      -X ${CONFIG_PKG}.BuildDate="`date -u +%FT%TZ`" \
      -X ${CONFIG_PKG}.BuildAuthor="`id -un`" \
      "

gotestsum_bin := bin/gotestsum$(EXT)
build: $(gotestsum_bin)
$(gotestsum_bin):
	@mkdir -p bin
	@go get -tool gotest.tools/gotestsum
	@go build -o $@ -ldflags="-s -w" gotest.tools/gotestsum

test2json_bin := bin/test2json$(EXT)
build: $(test2json_bin)
$(test2json_bin):
	@mkdir -p bin
	@go build -o $@ -ldflags="-s -w" cmd/test2json

.PHONY: test
test: ## Run all tests
	@go tool gotestsum ./... -test.count=1

.PHONY: clean
clean: ## Clean build artifacts and test results
	@rm -rf bin
	@rm -rf dist
	@rm -rf _test_results

.PHONY: download
download: ## Download go modules
	@go mod download

.PHONY: release
release: ## Create a release using goreleaser
	@go tool goreleaser $(GORELEASER_ARGS)

.PHONY: snapshot
snapshot: GORELEASER_ARGS= --clean --snapshot ## Create a snapshot release using goreleaser
snapshot: release
