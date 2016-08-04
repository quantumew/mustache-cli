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
	doc := `mustache-cli

        Command line interface for rendering mustache templates.
        If not data path is given it will expect data from stdin.

        Usage:
            mustache.go [<data-path>] <template-path>

        Options:
            -h --help        Show this message.

        Arguments:
            <data-path>      Path to data file to populate template.
            <template-path>  Path to template file.
    `
	arguments, _ := docopt.Parse(doc, nil, true, "Mustache 0.1", false)
	dataPath := arguments["<data-path>"].(string)
	templatePath := arguments["<template-path>"].(string)

	var data interface{}
	var err error
	data, err = loadJson(dataPath)

	if err != nil {
		data, err = loadYaml(dataPath)
	}
	handleError(err)

	output, err := mustache.RenderFile(templatePath, data)
	handleError(err)
	fmt.Println(output)
}

func readFromFile(filePath string) []byte {
	raw, err := ioutil.ReadFile(filePath)
	handleError(err)

	return raw
}

func loadYaml(filePath string) (interface{}, error) {
	raw := readFromFile(filePath)

	var data interface{}
	err := yaml.Unmarshal(raw, &data)

	return data, err
}

func loadJson(filePath string) (interface{}, error) {
	raw := readFromFile(filePath)

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
