Examples
========

This directory contains examples of using mustache with different data sources.
These scripts are also used as part of integration tests.

All examples use the same template `template.mustache`. There are three example scripts:

Runs mustache-cli with a JSON data file. Technically the contents are also valid YAML.

    ./run-example-json

Runs mustache-cli with a YAML data file.

    ./run-example-yaml

Runs mustache-cli with JSON data from stdin.

    ./run-example-stdin
