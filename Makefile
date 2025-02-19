ifeq ($(OS),Windows_NT)
	EXT := ".exe"
else
	EXT := ""
endif

# Run with java by default when environment variable is not set
LANG_IMPL_PATH ?= java

CONFIG_PKG="github.com/murex/gamekit-coffeemachine/settings"
export CGO_ENABLED=0

.PHONY: default
default: build ;

# Convenience target for automating release preparation
.PHONY: prepare
prepare: deps install-tools tidy lint build test

.PHONY: deps
deps:
	@go get -u -t ./...

.PHONY: lint
lint:
	@golangci-lint run -v

.PHONY: vet
vet:
	@go vet ./...

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build

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

.PHONY: run-cli
run-cli: $(cli_bin)
	@$< $(LANG_IMPL_PATH)

progress_tests_bin := bin/progress-tests$(EXT)
build: $(progress_tests_bin)
$(progress_tests_bin):
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
	@go get gotest.tools/gotestsum
	@go build -o $@ -ldflags="-s -w" gotest.tools/gotestsum

test2json_bin := bin/test2json$(EXT)
build: $(test2json_bin)
$(test2json_bin):
	@mkdir -p bin
	@go build -o $@ -ldflags="-s -w" cmd/test2json

define RUN_PROGRESS_TESTS
mkdir -p _test_results
export LANG_IMPL_PATH=$(1)
export LANGUAGE=$(basename $(1))
$(gotestsum_bin) \
  --format testdox \
  --junitfile _test_results/progress-tests-$(LANGUAGE).xml \
  --hide-summary=all \
  --raw-command \
  -- $(test2json_bin) -t -p progress-tests-$(LANGUAGE) bin/progress-tests$(EXT) -test.v=test2json
endef

.PHONY: run-progress
run-progress: $(progress_tests_bin) $(test2json_bin) $(gotestsum_bin)
	@$(call RUN_PROGRESS_TESTS,$(LANG_IMPL_PATH))

.PHONY: test
test:
	@env LANG_IMPL_PATH=java gotestsum ./...

.PHONY: clean
clean:
	@rm -rf bin
	@rm -rf dist
	@rm -rf _test_results


.PHONY: download
download:
	@go mod download

.PHONY: install-tools
install-tools: download
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %@latest

.PHONY: release
release:
	@goreleaser $(GORELEASER_ARGS)

.PHONY: snapshot
snapshot: GORELEASER_ARGS= --clean --snapshot
snapshot: release
