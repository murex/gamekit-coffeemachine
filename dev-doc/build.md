# Building the gamification kit on your machine

This section provides information related to the gamification kit development environment setup for those who wish to
build it tool locally.

## Clone gamekit repository - `Required`

```shell
git clone https://github.com/murex/gamekit-coffeemachine.git
cd gamekit-coffeemachine
```

## Install Go SDK - `Required`

The gamification kit is written in Go. This implies having Go compiler and tools installed on your machine.

Simply follow the instructions provided [here](https://go.dev/). Make sure to install **Go version 1.24** or higher.

## Install additional Go tools and utility packages

### Go IDE - `Optional`

You can check this [link](https://www.tabnine.com/blog/top-7-golang-ides-for-go-developers/)
for a list of recommended IDEs supporting Go language.

### GoReleaser utility - `Optional`

New versions of the gamification kit are released through [GoReleaser](https://goreleaser.com/).

You do not need it as long as you don't plan to release a new gamification kit version.

If you do, you can refer to [GoReleaser Installation Instructions](https://goreleaser.com/install/)
for installing it locally on your machine.

In most cases you will not even have to install it locally as the gamification kit new releases are built through
a [GoReleaser GitHub action](../.github/workflows/go_releaser.yml).

### golangci-lint package - `Optional`

We use the Go Linter aggregator [golangci-lint](https://golangci-lint.run/) to perform various static checks on the
gamification kit
code.

A [dedicated GitHub action](../.github/workflows/go_releaser.yml) triggers execution of golangci-lint every time a new
gamification kit version is being released.

Although not mandatory, we advise you to install it locally on your machine to check that your changes comply with
golangci-lint rules. Refer to [golangci-lint installation instructions](https://golangci-lint.run/welcome/install/)
for installation.

Once golangci-lint is installed, you can run it from the root directory:

```shell
make lint
```

Both local run and GitHub Action use [this configuration file](../.golangci.yml)

### gotestsum utility - `Optional`

We use [gotestsum](https://github.com/gotestyourself/gotestsum) for running tests
with the possibility to generate a xunit-compatible test report.

Although not mandatory, we advise you to install it locally on your machine as it greatly improves
readability of test results.
Refer to [gotestsum's Install section](https://github.com/gotestyourself/gotestsum#install)
for installation.

Once gotestsum is installed, you can run make's test target from the root directory:

- For running unit tests:

  ```shell
  make test
  ```

## Build gamification kit executables

To build the gamification kit executables locally on your machine, simply type the following from the root directory:

```shell
make
```

This command generates all the gamification kit executables in the [bin](../bin) directory.

| Executable                 | Description                                                                        |
|----------------------------|------------------------------------------------------------------------------------|
| `bin/progress-tests[.exe]` | Progress Runner executable. This is the test executable running all progress tests |
| `bin/gotestsum[.exe]`      | gotestsum executable used for running and rendering go test results (*)            |
| `bin/test2json[.exe]`      | go tool used for converting raw go test results into json (*)                      |
| `bin/cli[.exe]`            | Command Line Interface utility                                                     |

(*) These executables are built here so that they can be re-distributed along with progress-tests executable
without the need to have Go SDK installed on the end user machine.

## Run gamification kit executables

Both the Progress Runner and the Command Line Interface utility need a coffee machine implementation to interact with.

In the examples provided below, we assume that the coffee machine implementation is available
at the following relative path: `../kata-coffeemachine/java`. Replace this path with the actual path to your
coffee machine implementation.

### Run the progress runner

#### Default mode (one dot per test case)

Either:

```shell
env LANG_IMPL_PATH=../kata-coffeemachine/java make run-progress
```

Or:

```shell
./run-progress-runner.sh ../kata-coffeemachine/java
```

#### Verbose mode (one line per test case with test description)

```shell
./run-progress-runner.sh ../kata-coffeemachine/java -v
```

#### Very verbose mode (one line per test case with test description + failed tests output details)

```shell
./run-progress-runner.sh ../kata-coffeemachine/java -vv
```

### Run the command line interface utility

Either:

```shell
env LANG_IMPL_PATH=../kata-coffeemachine/java make run-cli
```

Or:

```shell
./bin/cli ../kata-coffeemachine/java
```
