package rules

import (
	"fmt"

	"github.com/Woody1193/miletos-test/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Exists Tests", func() {

	// Tests the Exists function under various conditions
	DescribeTable("InvoiceExists - Conditions",
		func(inv *types.InvoiceItem, rec *types.ReceivablesItem, okay bool, expected error) {

			// Run the function and collect the results
			f, err := InvoiceExists(inv, rec)

			// Verify that the results are correct
			Expect(f).To(Equal(okay))
			if okay {
				Expect(err).To(BeNil())
			} else {
				Expect(err).To(Equal(expected))
			}
		},
		Entry("Invoice is nil", nil, &types.ReceivablesItem{}, false, fmt.Errorf("Invoice does not exist")),
		Entry("Receivables is nil", &types.InvoiceItem{}, nil, true, nil),
		Entry("Invoice and receivables exist", &types.InvoiceItem{}, &types.ReceivablesItem{}, true, nil))
})
