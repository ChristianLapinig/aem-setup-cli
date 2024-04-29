# AEM Setup CLI

## Summary

`aem-setup-cli` generates a base environment to run AEM (Adobe Experience Manager)
locally. The tool will create the folders and add the files necessary for both `author`
and `publish` instances. It supports options, such as specifying where to setup
your local AEM environment, mount an existing `crx-quickstart` repository, and the use
of a different AEM quickstart JAR if given.

## Installation

As of April 29, 2024, [Go](https://go.dev/) is required to install the tool. After
you have installed Go and have your `GOPATH` defined, please follow the steps below to install `aem-setup-cli`:

```bash
# 1. Clone the repository
$ git clone https://github.com/ChristianLapinig/aem-setup-cli

# 2. cd into the repository
$ cd aem-setup-cli

# 3. Run the 'go install' command
$ go install
```

## Usage

The main command to setup a local AEM environment is the `aem-setup-cli setup` command.
A valid `license.properties` file and a path to a `cq-quickstart` JAR are required to properly run AEM and need to be
passed to the `setup` command:

```bash
$ aem-setup-cli setup /path/to/license.properties /path/to/cq-quickstart.jar
```

This will setup the environment in the current directory the tool is executed.

The `setup` command also has the following flags/options:

```
-h, --help                help for setup
-m, --mount string        Path to an existing AEM repository that will be mounted to your instances.
-p, --path string         Path where AEM should be setup. Default is the current directory you are in. (default ".")
```

Example:

```bash
$ aem-setup-cli setup /path/to/license.properties -p /path/to/setup/env -m /path/to/crx-quickstart
```
