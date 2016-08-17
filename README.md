Mustache Cli
============

Command line interface to mustache template engine in Go.
Basically a simple CLI wrapper for [cbroglie/mustache](https://github.com/cbroglie/mustache).
Works with YAML and JSON. Data can be piped in via stdin or passed in as a file name via an option.

See examples directory for a more in depth example of using JSON, YAML, and stdin.

[Build of latest release.](https://github.com/quantumew/mustache-cli/releases/download/v0.2/mustache-cli)

Usage:

    mustache-cli [<data-file>] <template-path>
    mustache-cli <template-path>

Options:

    -h --help        - Show this message.

Arguments:
    <data-file>      - Path to data file.

    <template-path>  - Path to template file.

Examples:

    mustache-cli data-source.json template.mustache

    mustache-cli data-source.yaml template.mustache

    cat data-source.json | mustache-cli template.mustache
