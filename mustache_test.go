package main

import (
    "testing"
    "fmt"
)

func TestFailedJson(t *testing.T) {
    json := []byte("{\"key\": \"val\" \"otherKey\": \"noComma\"}")
    _, err := loadJson(json)
    expectedMsg := "invalid character '\"' after object key:value pair"

    if err == nil || err.Error() != expectedMsg {
        t.Errorf("got %s, want error", err)
    }
}

func ExampleLoadJson() {
    json := []byte("{ \"key\": \"val\"}")
    data, _ := loadJson(json)
    fmt.Println(data)
    // Output: map[key:val]
}

func ExampleLoadJsonFailed() {
    json := []byte("{\"key\": \"val\" \"otherKey\": \"noComma\"}")
    _, err := loadJson(json)
    fmt.Println(err)
    // Output: invalid character '"' after object key:value pair
}

func ExampleLoadUnknown() {
    input := []byte("key: val")
    data, _ := loadUnknown(input)
    fmt.Println(data)
    // Output: map[key:val]
}

func ExampleLoadUnknownFailed() {
    input := []byte("key: %val%")
    _, err := loadUnknown(input)
    fmt.Println(err)
    // Output:
    // Could not decode provided data.
    // Child Error: invalid character 'k' looking for beginning of value
    // Child Error: error converting YAML to JSON: yaml: [while scanning for the next token] found character that cannot start any token at line 1, column 6
}

func ExampleLoadYaml() {
    yaml := []byte("key: val")
    data, _ := loadYaml(yaml)
    fmt.Println(data)
    // Output: map[key:val]
}

func ExampleLoadYamlFailed() {
    yaml := []byte("key\": \"%val%\"}")
    _, err := loadYaml(yaml)
    fmt.Println(err)
    // Output:
    // error converting YAML to JSON: yaml: [while parsing a block mapping] did not find expected key at line 1, column 14
}
