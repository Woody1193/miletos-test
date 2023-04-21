package rules

import (
	"github.com/Woody1193/miletos-test/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Date Tests", func() {

	DescribeTable("NotPastDue - Conditions",
		func(inv *types.InvoiceItem, rec *types.ReceivablesItem, okay bool, expected error) {

			// Run the function and collect the results
			f, err := NotPastDue(inv, rec)

			// Verify that the results are correct
			Expect(f).To(Equal(okay))
			if okay {
				Expect(err).To(BeNil())
			} else {
				Expect(err).To(Equal(expected))
			}
		})
})
