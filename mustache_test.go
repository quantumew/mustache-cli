package main

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.com/onsi/gomega/gexec"
    "github.com/onsi/gomega/gbytes"
    "os/exec"
    "time"
)

var _ = Describe("Mustache", func() {
    successDecodeOutput := map[string]interface{}{"key":"val"}
    expectedTemplate := "my name is Montoooorb69 but you can call me Monica Rachel Pheobe.\n" +
        "I come from Cloooooooooooooooo in the Florbatorb galaxy.\n" +
        "I would like to Drink earth liquids, Kill the producers of 3rd Rock From The Sun, Go to Disney World.\n\n" +
        "Thank you,\nMontoooorb69"

    assertExampleSuccess := func(path string) {
        command := exec.Command(path)
        session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
        Expect(err).ShouldNot(HaveOccurred())
        session.Wait(5 * time.Second)
        Expect(session.Out).Should(gbytes.Say(expectedTemplate))
        Expect(session).Should(gexec.Exit(0))
    }

    Describe("When Cli is executed", func() {
        // These tests leverage the example scripts for ease of testing.
        Context("With data file", func() {
            It("should successfully output example template from JSON data source", func() {
                assertExampleSuccess("./examples/run-example-json")
            })

            It("should successfully output example template from YAML data source", func() {
                assertExampleSuccess("./examples/run-example-yaml")
            })
        })

        Context("From Stdin", func() {
            It("should successfully output example template", func() {
                assertExampleSuccess("./examples/run-example-stdin")
            })
        })
    })

    Describe("decodeData", func() {
        Context("When called with valid JSON data", func() {
            It("should decode JSON to a map", func() {
                json := []byte("{ \"key\": \"val\"}")
                data, _ := decodeData(json)
                Expect(data).To(Equal(successDecodeOutput))
            })
        })

        Context("When called with valid Yaml data", func() {
            It("Should decode YAML to a map", func() {
                input := []byte("key: val")
                data, _ := decodeData(input)
                Expect(data).To(Equal(successDecodeOutput))
            })
        })

        Context("When called with data in invalid format", func() {
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
