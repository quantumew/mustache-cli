//usr/bin/env go run $0 $@; exit;
package main

import (
	"encoding/json"
	"fmt"
	"github.com/cbroglie/mustache"
	"github.com/docopt/docopt-go"
	"io/ioutil"
	"os"
)

func main() {
	doc := `mustache
        Usage:
            mustache <data-path> <template-path>

        Options:
            -h --help        Show this message.

        Arguments:
            <data-path>      Path to data file to populate template.
            <template-path>  Path to template file.
    `
	arguments, _ := docopt.Parse(doc, nil, true, "Mustache 0.1", false)
	filePath := arguments["<data-path>"].(string)
	templatePath := arguments["<template-path>"].(string)
	raw, err := ioutil.ReadFile(filePath)
	handleError(err)

	var data interface{}
	err = json.Unmarshal(raw, &data)
	handleError(err)

	output, err := mustache.RenderFile(templatePath, data)
	handleError(err)
	fmt.Println(output)
}

func handleError(err interface{}) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
