package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFibonacci(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fibonacci Test Suite")
}

var _ = Describe("Fibonacci", func() {
	It("Tests", func() {
		Expect(Fibonacci(0)).To(Equal(0))
	})
})
