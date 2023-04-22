package rules

import (
	"fmt"

	"github.com/Woody1193/miletos-test/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

var _ = Describe("Amount Tests", func() {

	// Tests the AmountsEqual function under various conditions
	DescribeTable("AmountsEqual - Conditions",
		func(inv *types.InvoiceItem, rec *types.ReceivablesItem, okay bool, expected error) {

			// Run the function and collect the results
			f, err := AmountsEqual(inv, rec)

			// Verify that the results are correct
			Expect(f).To(Equal(okay))
			if okay {
				Expect(err).To(BeNil())
			} else {
				Expect(err).To(Equal(expected))
			}
		},
		Entry("Invoice is nil", nil, &types.ReceivablesItem{}, true, nil),
		Entry("Receivables is nil", &types.InvoiceItem{}, nil, true, nil),
		Entry("Invoice and receivables amounts do not match", &types.InvoiceItem{Amount: decimal.New(100, 0)},
			&types.ReceivablesItem{Amount: decimal.New(200, 0)}, false,
			fmt.Errorf("Invoice amount of 100 does not match receivables amount of 200")),
		Entry("Invoice and receivables amounts match", &types.InvoiceItem{Amount: decimal.New(100, 0)},
			&types.ReceivablesItem{Amount: decimal.New(100, 0)}, true, nil))
})
