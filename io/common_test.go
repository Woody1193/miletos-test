package io

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Create a new test runner we'll use to test all the
// modules in the io package
func TestIO(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IO Suite")
}
