package main

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "github.com/onsi/gomega/gexec"
    "github.com/onsi/gomega/gbytes"
    "os/exec"
)

var _ = Describe("Mustache", func() {
    var (
        successDecodeOutput map[string]interface{}
        expectedTemplate string
    )

    BeforeEach(func() {
        successDecodeOutput = map[string]interface{}{"key":"val"}
        expectedTemplate = "my name is Montoooorb69 but you can call me Monica Rachel Pheobe.\n" +
            "I come from Cloooooooooooooooo in the Florbatorb galaxy.\n" +
            "I would like to Drink earth liquids, Kill the producers of 3rd Rock From The Sun, Go to Disney World.\n\n" +
            "Thank you,\nMontoooorb69"
    })

    Describe("Cli", func() {
        Context("With data file", func() {
            It("should successfully output example template from JSON data source", func() {
                command := exec.Command("./examples/run-example-json")
                session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
                Expect(err).ShouldNot(HaveOccurred())
                Eventually(session.Out).Should(gbytes.Say(expectedTemplate))
                Eventually(session).Should(gexec.Exit(0))
            })

            It("should successfully output example template from YAML data source", func() {
                command := exec.Command("./examples/run-example-yaml")
                session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
                Expect(err).ShouldNot(HaveOccurred())
                Eventually(session.Out).Should(gbytes.Say(expectedTemplate))
                Eventually(session).Should(gexec.Exit(0))
            })
        })

        Context("From Stdin", func() {
            It("should successfully output example template", func() {
                command := exec.Command("./examples/run-example-stdin")
                session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
                Expect(err).ShouldNot(HaveOccurred())
                Eventually(session.Out).Should(gbytes.Say(expectedTemplate))
                Eventually(session).Should(gexec.Exit(0))
            })
        })
    })

    Describe("Testing mustache decode data", func() {
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
