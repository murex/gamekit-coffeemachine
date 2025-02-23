ifeq ($(OS),Windows_NT)
	EXT := ".exe"
else
	EXT := ""
endif

CONFIG_PKG="github.com/murex/gamekit-coffeemachine/settings"
export CGO_ENABLED=0

.PHONY: default
default: build ;

# Convenience target for automating release preparation
.PHONY: prepare
prepare: deps tidy lint build test

.PHONY: deps
deps:
	@go get -u -t tool ./...

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
test:
	@go tool gotestsum ./... -test.count=1

.PHONY: clean
clean:
	@rm -rf bin
	@rm -rf dist
	@rm -rf _test_results


.PHONY: download
download:
	@go mod download

.PHONY: release
release:
	@go tool goreleaser $(GORELEASER_ARGS)

.PHONY: snapshot
snapshot: GORELEASER_ARGS= --clean --snapshot
snapshot: release
