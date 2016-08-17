//usr/bin/env go run $0 $@; exit;
package main

import (
	"fmt"
	"github.com/cbroglie/mustache"
	"github.com/docopt/docopt-go"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	doc := `Mustache Cli

        Command line interface for rendering mustache templates.
        Data is either expected via data option with a file name or
        via stdin. If data option is given that will be used.

        Usage:
            mustache [<data-file>] <template-path>
            mustache <template-path>

        Options:
            -h --help        Show this message.

        Arguments:
            <data-file>      Path to data file.

            <template-path>  Path to template file.
    `
	arguments, _ := docopt.Parse(doc, nil, true, "Mustache 0.1", false)
	dataPath := arguments["<data-file>"]
	templatePath := arguments["<template-path>"].(string)
	var (
        err error
	    data interface{}
    )

	if dataPath == nil {
		data, err = loadFromStdin()
	} else {
		path := dataPath.(string)
		data, err = loadFromFile(path)
	}

	if err != nil {
		logError("Error occurred loading data", err)
		os.Exit(1)
	}

	output, err := mustache.RenderFile(templatePath, data)

	if err != nil {
		logError("Error occurred rendering template", err)
	}
	fmt.Println(output)
}

func loadFromFile(path string) (interface{}, error) {
	raw, readErr := ioutil.ReadFile(path)

	if readErr != nil {
		return nil, readErr
	}

    return decodeData(raw)
}

func loadFromStdin() (interface{}, error) {
	raw, readErr := ioutil.ReadAll(os.Stdin)

	if readErr != nil {
		return nil, readErr
	}

	return decodeData(raw)
}

func decodeData(raw []byte) (interface{}, error) {
	var data interface{}
	err := yaml.Unmarshal(raw, &data)

	return data, err
}

func logError(msg string, err error) {
	log := log.New(os.Stderr, "", 0)
	log.Println(msg)
	log.Println(err.Error())
}
