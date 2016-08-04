Mustache Cli
============

Command line interface to mustache template engine in Go.
Basically a simple CLI wrapper for (cbroglie/go-mustache)[https://github.com/cbroglie/go-mustache].
Works with YAML and JSON. Data can be piped in via stdin or passed in as a file name via an option.

Usage:

	mustache.go <template-path> [options]

Options:

	-d --data FILE   - Path to data to use in template.

	-h --help        - Show this message.

Arguments:

	<template-path>  - Path to template file.

Examples:

	mustache.go template.mustache --data data-source.json

	mustache.go template.mustache --data data-source.yaml

	jq '.' data-source.json | mustache.go template.mustache
