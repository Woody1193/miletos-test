package types

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

var _ = Describe("Data Tests", func() {

	// Tests that the Key function works correctly
	It("InvoiceItem - Key - Works", func() {

		// Create a new invoice item
		item := InvoiceItem{
			ID: "123",
		}

		// Verify that the key is correct
		Expect(item.Key()).To(Equal("123"))
	})

	// Tests that the Verify function works correctly under various conditions
	DescribeTable("InvoiceItem - Verify - Conditions",
		func(id string, dueDate time.Time, amount decimal.Decimal, expected error) {

			// Create a new invoice item
			item := InvoiceItem{
				ID:      id,
				DueDate: dueDate,
				Amount:  amount,
			}

			// Verify that the item is invalid
			if expected == nil {
				Expect(item.Verify()).To(BeNil())
			} else {
				Expect(item.Verify()).To(Equal(expected))
			}
		},
		Entry("Empty ID", "", time.Now(), decimal.New(100, 0), EmptyIDError),
		Entry("Missing Due Date", "123", time.Time{}, decimal.New(100, 0), NoDateError),
		Entry("Missing Amount", "123", time.Now(), decimal.Zero, NoAmountError),
		Entry("Invalid Amount", "123", time.Now(), decimal.New(1001, -1),
			fmt.Errorf("Amount of 100.1 was invalid")),
		Entry("No Errors", "123", time.Now(), decimal.New(100, 0), nil))

	// Tests that the Key function works correctly
	It("ReceivablesItem - Key - Works", func() {

		// Create a new receivables item
		item := ReceivablesItem{
			ID: "123",
		}

		// Verify that the key is correct
		Expect(item.Key()).To(Equal("123"))
	})

	// Tests that the Verify function works correctly under various conditions
	DescribeTable("ReceivablesItem - Verify - Conditions",
		func(id string, date *time.Time, amount *decimal.Decimal, expected error) {

			// Create a new receivables item
			item := ReceivablesItem{
				ID:     id,
				Date:   date,
				Amount: amount,
			}

			// Verify that the item is invalid
			if expected == nil {
				Expect(item.Verify()).To(BeNil())
			} else {
				Expect(item.Verify()).To(Equal(expected))
			}
		},
		Entry("Empty ID", "", timePtr(time.Now()), decimalPtr(decimal.New(100, 0)), EmptyIDError),
		Entry("Missing Date", "123", nil, decimalPtr(decimal.New(100, 0)), NoDateError),
		Entry("Missing Amount", "123", timePtr(time.Now()), nil, NoAmountError),
		Entry("Invalid Amount", "123", timePtr(time.Now()), decimalPtr(decimal.New(1001, -1)),
			fmt.Errorf("Amount of 100.1 was invalid")),
		Entry("No Errors", "123", timePtr(time.Now()), decimalPtr(decimal.New(100, 0)), nil))
})

// Helper functions to create pointers to time values
func timePtr(time time.Time) *time.Time {
	return &time
}

// Helper functions to create pointers to decimal values
func decimalPtr(decimal decimal.Decimal) *decimal.Decimal {
	return &decimal
}
