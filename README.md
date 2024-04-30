# AEM Setup CLI

## Summary

`aem-setup-cli` generates a base environment to run AEM (Adobe Experience Manager)
locally. The tool will create the folders and add the files necessary for both `author`
and `publish` instances. It supports options, such as specifying where to setup
your local AEM environment, and mounting an existing `crx-quickstart` repository to a specified instance (in-progress).

## Installation

As of April 29, 2024, [Go](https://go.dev/) is required to install the tool. After
you have installed Go and have your `GOPATH` defined, run the following command to install `aem-setup-cli`:

```bash
$ go install github.com/ChristianLapinig/aem-setup-cli@latest
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
-h, --help              help for setup
-i, --instance string   Instance where the given repository from 'mount' should be mounted.
-m, --mount string      Path to an existing AEM repository that will be mounted to your instances. Must be used with the instance flag.
-p, --path string       Path where AEM should be setup. Default is the current directory you are in. (default ".")
```

Example:

```bash
$ aem-setup-cli setup /path/to/license.properties -p /path/to/setup/env -m /path/to/crx-quickstart
```

## Contributing

To contribute to the project, please fork the repository and clone it locally:

```bash
$ git clone https://github.com/YourUsename/aem-setup-cli
```

If there are any issues or bugs, please open an issue in the issues tab.

### Creating a Pull Request

Use the following steps below to create a pull request:

1. Create a branch under your fork
2. Push the changes to that branch
3. Add tests if applicable
4. Create a pull request. If there is an issue associated with your PR, please add it
   to the description.

## Contributors

- [Chrisitan Lapinig](https://github.com/ChristianLapinig)
