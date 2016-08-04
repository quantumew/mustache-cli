Mustache Cli
============

Command line interface to mustache template engine in Go.
Basically a simple CLI wrapper for [cbroglie/go-mustache](https://github.com/cbroglie/go-mustache).
Works with YAML and JSON. Data can be piped in via stdin or passed in as a file name via an option.

See examples directory for a more in depth example of using JSON, YAML, and stdin.

Installation:

    go get github.com/quantumew/mustache-cli && make

Usage:

	mustache <template-path> [options]

Options:

	-d --data FILE   - Path to data to use in template.

	-h --help        - Show this message.

Arguments:

	<template-path>  - Path to template file.

Examples:

	mustache template.mustache --data data-source.json

	mustache template.mustache --data data-source.yaml

	cat data-source.json | mustache template.mustache
