//usr/bin/env go run $0 $@; exit;
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/cbroglie/mustache"
	"github.com/docopt/docopt-go"
	"github.com/ghodss/yaml"
)

var logger = log.New(os.Stderr, "", 0)

func main() {
	doc := `Mustache Cli

        Command line interface for rendering mustache templates.
        Data is either expected via data option with a file name or
		via stdin. If data option is given that will be used.

        Examples:
            # Basic template usage
            mustache data.json template.mustache

            # Pull variables from environment
            mustache ENV template.mustache

            # Pull variables from environment with overrides. This will merge starting with env vars.
            # Think of order as priority.
			mustache ENV template.mustache --override data.json --override data1.json

			# get base data from stdin
    		cat data-source.json | mustache template.mustache

        Usage:
            mustache [<data-file>] <template-path> [--override=<data-file>]...
            mustache <template-path> [--override=<data-file>]...


        Arguments:
            <data-file>      Path to data file. ENV is a special identifier to use environment variables.

			<template-path>  Path to template file.

        Options:
            -h --help            Show help message.

            -o --override <file> Override data files. Overrides will be done in order.
	`

	arguments, _ := docopt.Parse(doc, nil, true, "Mustache 1.0.0", false)
	dataPath := arguments["<data-file>"]
	templatePath := arguments["<template-path>"].(string)
	overrideList := arguments["--override"].([]string)

	var (
		err     error
		data    interface{}
		context = make([]interface{}, 1+len(overrideList))
	)

	if dataPath == nil {
		data, err = loadFromStdin()
	} else {
		path := dataPath.(string)
		data = loadFromEnvOrFile(path)
	}
	context[0] = data

	for i, override := range overrideList {
		newData := loadFromEnvOrFile(override)
		checkErr(err)
		context[i+1] = newData
	}

	checkErr(err)
	output, err := mustache.RenderFile(templatePath, context...)
	checkErr(err)
	fmt.Println(output)
}

func checkErr(err error) {
	if err != nil {
		logError("Error occurred rendering template", err)
		os.Exit(1)
	}
}

func loadFromEnvOrFile(path string) interface{} {
	if path == "ENV" {
		return loadFromEnv()
	}

	d, err := loadFromFile(path)
	checkErr(err)
	return d
}

func loadFromEnv() interface{} {
	m := map[string]string{}

	for _, item := range os.Environ() {
		splits := strings.Split(item, "=")
		m[splits[0]] = os.Getenv(splits[0])
	}

	return m
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
	logger.Println(msg)
	logger.Println(err.Error())
}
