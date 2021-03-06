# Mustache Cli

**DEPRECATED** [cbroglie/mustache](https://github.com/cbroglie/mustache), now has a CLI wrapper of its own.

Command line interface to mustache template engine in Go.
Basically a simple CLI wrapper for [cbroglie/mustache](https://github.com/cbroglie/mustache).
Works with YAML and JSON. Data can be piped in via stdin or passed in as a file name via an option.

See examples directory for a more in depth example of using JSON, YAML, and stdin.

[Build of latest release.](https://github.com/quantumew/mustache-cli/releases)

## Usage

    mustache [<data-file>] <template-path>
    mustache <template-path>

## Examples

    # Basic template usage
    mustache data.json template.mustache

    # Pull variables from environment
    mustache ENV template.mustache

    # Pull variables from environment with overrides. This will merge starting with env vars.
    # Think of order as priority.
    mustache ENV template.mustache --override data.json --override data1.json

    # get base data from stdin
    cat data-source.json | mustache template.mustache

## Arguments

    <data-file>      Path to data file. ENV is a special identifier to use environment variables.

    <template-path>  Path to template file.

# Options

    -h --help            Show help message.
    -o --override <file> Override data files. Overrides will be done in order.

See also: [EXAMPLES](examples/README.md)

## Build

If you need a build not in releases, you can either request it or build it yourself. Here is how you build it.

    go get github.com/quantumew/mustache-cli
    cd "$GOPATH/src/github.com/quantumew/mustache-cli"
    make
    mv mustache <in your path somewhere>
