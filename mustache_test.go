package main

import (
    . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mustache unit tests", func() {
    var successDecodeOutput map[string]interface{}

    BeforeEach(func() {
        successDecodeOutput = map[string]interface{}{"key":"val"}
    })

    Describe("Testing mustach functions", func() {
        Context("Decode data with valid data", func() {
            It("should decode JSON to a map", func() {
                json := []byte("{ \"key\": \"val\"}")
                data, _ := decodeData(json)
                Expect(data).To(Equal(successDecodeOutput))
            })

            It("Should decode YAML to a map", func() {
                input := []byte("key: val")
                data, _ := decodeData(input)
                Expect(data).To(Equal(successDecodeOutput))
            })
        })

        Context("Decoding data with invalid format", func() {
            It("should fail to load malformatted JSON", func() {
                json := []byte("{ \"key\": %val%\"}")
                _, err := decodeData(json)
                expectedErr := "error converting YAML to JSON: yaml: " +
                    "[while scanning for the next token] found character that cannot start any token at line 1, column 10"
                Expect(err.Error()).To(Equal(expectedErr))
            })

            It("should error with invalid format", func() {
                yaml := []byte("key\": \"%val%\"}")
                _, err := decodeData(yaml)
                expectedErr := "error converting YAML to JSON: yaml: " +
                    "[while parsing a block mapping] did not find expected key at line 1, column 14"
                Expect(err.Error()).To(Equal(expectedErr))
            })
        })
    })
})
