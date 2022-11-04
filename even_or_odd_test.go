package main_test

import (
	"testing"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

// This is the entry point for Ginkgo - the go test runner will run this function when you run 'go test' or 'ginkgo'.
func TestEvenOrOdd(t *testing.T) {
	RegisterFailHandler(Fail) // RegisterFailHandler connects Ginkgo to Gomega.
	RunSpecs(t, "EvenOrOdd Test Suite") // RunSpecs is the entry point for the Ginkgo spec runner.
}