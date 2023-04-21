package types

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Results Tests", func() {

	// Tests that the NewErrorResult function works correctly
	It("NewErrorResult - Works", func() {

		// Create a new error result
		result := NewErrorResult("test", 1, EmptyIDError)

		// Verify that the result is correct
		Expect(result.File).To(Equal("test"))
		Expect(result.Line).To(Equal(uint(1)))
		Expect(result.Error).To(Equal(EmptyIDError.Error()))
	})
})
