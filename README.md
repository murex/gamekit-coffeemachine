[![Go build and test](https://github.com/murex/gamekit-coffeemachine/actions/workflows/go.yml/badge.svg)](https://github.com/murex/gamekit-coffeemachine/actions/workflows/go.yml)
[![Go lint](https://github.com/murex/gamekit-coffeemachine/actions/workflows/golangci_lint.yml/badge.svg)](https://github.com/murex/gamekit-coffeemachine/actions/workflows/golangci_lint.yml)
[![Go release](https://github.com/murex/gamekit-coffeemachine/actions/workflows/go_releaser.yml/badge.svg)](https://github.com/murex/gamekit-coffeemachine/actions/workflows/go_releaser.yml)
[![Dependabot Updates](https://github.com/murex/gamekit-coffeemachine/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/murex/gamekit-coffeemachine/actions/workflows/dependabot/dependabot-updates)

# Gamification kit for the coffee machine kata

This repository provides items focused on gamifying the coffee machine kata.

These items are by design "programming language agnostic",
e.g. they may be used with the coffee machine kata in any language.

## Available Gamification Tools

### Progress Runner

This tool allows to test the progress of a kata implementation.

Tests that are run by the progress runner are "black box" tests.

Their order of execution follows the iterations described in the kata.

### Command Line Interface utility

This tool allows to interact with a kata implementation from a terminal,
through sending commands and displaying the response sent by the implementation.

This a low-level tool which main purpose is to help tune the communication protocol
bootstrap files when adding support for a new programming language.

It's not intended to be used directly by kata participants.

## Overall Architecture

Both the Progress Runner and the Command Line Interface utility are built
on top of a common text protocol.
Refer to [Inter-Process Text Protocol](./dev-doc/inter-process-text-protocol.md) for further details.

```mermaid
graph LR
    subgraph runner [Progress Runner]
        PROGRESS_RUNNER(Progress<br>Runner)
        CLI_DRIVER_1(Command Line<br>Driver)
    end

    subgraph cli [Command Line Interface Utility]
        CLI_RUNNER(CLI<br>Runner)
        CLI_DRIVER_2(Command Line<br>Driver)
    end

    PROTOCOL{{"Inter-Process<br>Text<br>Protocol"}}

    subgraph java [Java]
        JAVA_RUNNER(Command Line<br>Runner)
        JAVA_FACADE(Facade)
        JAVA_IMPL(Implementation)
    end

    subgraph cpp [C++]
        CPP_RUNNER(Command Line<br>Runner)
        CPP_FACADE(Facade)
        CPP_IMPL(Implementation)
    end

    subgraph python [Python]
        PYTHON_RUNNER(Command Line<br>Runner)
        PYTHON_FACADE(Facade)
        PYTHON_IMPL(Implementation)
    end

    PROGRESS_RUNNER --> CLI_DRIVER_1 --> PROTOCOL
    CLI_RUNNER --> CLI_DRIVER_2 --> PROTOCOL
    PROTOCOL --> JAVA_RUNNER --> JAVA_FACADE --> JAVA_IMPL
    PROTOCOL --> CPP_RUNNER --> CPP_FACADE --> CPP_IMPL
    PROTOCOL --> PYTHON_RUNNER --> PYTHON_FACADE --> PYTHON_IMPL
    classDef testModule fill:#369;
    classDef implModule fill:#693;
    class PROGRESS_RUNNER testModule;
    class CLI_DRIVER_1 testModule;
    class CLI_RUNNER testModule;
    class CLI_DRIVER_2 testModule;
    class JAVA_RUNNER testModule;
    class JAVA_FACADE testModule;
    class JAVA_IMPL implModule;
    class CPP_RUNNER testModule;
    class CPP_FACADE testModule;
    class CPP_IMPL implModule;
    class PYTHON_RUNNER testModule;
    class PYTHON_FACADE testModule;
    class PYTHON_IMPL implModule;
```

## Repository Breakdown

This repository (`test-coffeemachine`) provides the client tool runners:

- Progress Runner
- Command Line Interface Runner

The kata repository (`kata-coffeemachine`) contains the implementation of the coffee machine kata in different
languages.
For each supported language:

- The command line runner is fully implemented.
- The facade skeleton is provided.

The parts remaining to be implemented by kata participants are

- the actual implementation of kata.
- the facade implementation, wiring the implementation to the command line runner.

## TODO

- [ ] Add a diagram with the 2 repositories
- [x] Add instructions on how to run the tools
- [x] Move protocol details to a separate page
- [ ] Add instructions on how to implement the facade in a new language
- [x] Add development instructions
- [x] Migrate this repository to murex (public)

## Building, testing and releasing coaching helpers for the coffee machine kata

Refer to [development documentation](./dev-doc/README.md) for details.

## How to Contribute?

These tools are still at an early stage of development,
and there are plenty of features that we would like to add in the future.

Refer to [CONTRIBUTING.md](./CONTRIBUTING.md) for general contribution agreement and guidelines.

## License

Contents from this repository are made available under the terms of the [MIT License](LICENSE.md)
which accompanies this distribution, and is available at the
[Open Source site](https://opensource.org/licenses/MIT).
