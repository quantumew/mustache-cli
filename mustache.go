//usr/bin/env go run $0 $@; exit;
package main

import (
	"encoding/json"
	"fmt"
	"github.com/cbroglie/mustache"
	"github.com/docopt/docopt-go"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

func main() {
	doc := `Mustache Cli

        Command line interface for rendering mustache templates.
        Data is either expected via data option with a file name or
        via stdin. If data option is given that will be used.

        Usage:
            mustache <template-path> [options]

        Options:
            -d --data FILE   Path to data to use in template.

            -h --help        Show this message.

        Arguments:
            <template-path>  Path to template file.
    `
	arguments, _ := docopt.Parse(doc, nil, true, "Mustache 0.1", false)
	dataPath := arguments["--data"]
	templatePath := arguments["<template-path>"].(string)

	var readErr error
	var raw []byte

	if dataPath == nil {
		raw, readErr = ioutil.ReadAll(os.Stdin)
	} else {
		path := dataPath.(string)
		raw, readErr = ioutil.ReadFile(path)
	}
	handleError(readErr)

	var data interface{}
	var err error
	data, err = loadJson(raw)

	if err != nil {
		data, err = loadYaml(raw)
	}
	handleError(err)

	output, err := mustache.RenderFile(templatePath, data)
	handleError(err)
	fmt.Println(output)
}

func loadYaml(raw []byte) (interface{}, error) {
	var data interface{}
	err := yaml.Unmarshal(raw, &data)

	return data, err
}

func loadJson(raw []byte) (interface{}, error) {
	var data interface{}
	err := json.Unmarshal(raw, &data)

	return data, err
}

func handleError(err interface{}) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
