package rules

import (
	"fmt"
	"time"

	"github.com/Woody1193/miletos-test/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Date Tests", func() {

	// Tests the NotPastDue function under various conditions
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
		},
		Entry("Receivables is not nil", &types.InvoiceItem{DueDate: timePtr(time.Now())}, &types.ReceivablesItem{}, true, nil),
		Entry("Invoice date is in the past", &types.InvoiceItem{DueDate: timePtr(time.Date(2022, time.April, 22, 0, 0, 0, 0, time.UTC))}, nil, false,
			fmt.Errorf("Invoice due date of 2022-04-22 00:00:00 +0000 UTC has past")),
		Entry("Invoice date is in the future", &types.InvoiceItem{DueDate: timePtr(time.Now().AddDate(0, 0, 1))}, nil, true, nil))

	// Tests the PaidOnTime function under various conditions
	DescribeTable("PaidOnTime - Conditions",
		func(inv *types.InvoiceItem, rec *types.ReceivablesItem, okay bool, expected error) {

			// Run the function and collect the results
			f, err := PaidOnTime(inv, rec)

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
		Entry("Invoice due date is before receivables paid date",
			&types.InvoiceItem{DueDate: timePtr(time.Date(2022, time.April, 22, 0, 0, 0, 0, time.UTC))},
			&types.ReceivablesItem{Date: timePtr(time.Date(2022, time.April, 23, 0, 0, 0, 0, time.UTC))}, false,
			fmt.Errorf("Invoice due date of 2022-04-22 00:00:00 +0000 UTC is before receivables paid date of 2022-04-23 00:00:00 +0000 UTC")),
		Entry("Invoice due date is after receivables paid date",
			&types.InvoiceItem{DueDate: timePtr(time.Date(2022, time.April, 23, 0, 0, 0, 0, time.UTC))},
			&types.ReceivablesItem{Date: timePtr(time.Date(2022, time.April, 22, 0, 0, 0, 0, time.UTC))}, true, nil),
		Entry("Invoice due date is equal to receivables paid date",
			&types.InvoiceItem{DueDate: timePtr(time.Date(2022, time.April, 22, 0, 0, 0, 0, time.UTC))},
			&types.ReceivablesItem{Date: timePtr(time.Date(2022, time.April, 22, 0, 0, 0, 0, time.UTC))}, true, nil))

	// Tests the DateNotInFuture function under various conditions
	DescribeTable("DateNotInFuture - Conditions",
		func(rec *types.ReceivablesItem, okay bool, expected error) {

			// Run the function and collect the results
			f, err := DateNotInFuture(&types.InvoiceItem{DueDate: timePtr(time.Now())}, rec)

			// Verify that the results are correct
			Expect(f).To(Equal(okay))
			if okay {
				Expect(err).To(BeNil())
			} else {
				Expect(err).To(Equal(expected))
			}
		},
		Entry("Receivables is nil", nil, true, nil),
		Entry("Receivables date is in the future", &types.ReceivablesItem{Date: timePtr(today().AddDate(0, 2, 0))}, false,
			fmt.Errorf("Receivables date of %s is in the future", today().AddDate(0, 2, 0))),
		Entry("Receivables date is in the past",
			&types.ReceivablesItem{Date: timePtr(time.Date(2022, time.April, 22, 0, 0, 0, 0, time.UTC))}, true, nil))
})
