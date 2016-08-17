package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMustacheCli(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MustacheCli Suite")
}
