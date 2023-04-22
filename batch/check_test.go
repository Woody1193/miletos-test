package batch

import (
	"testing"
	"time"

	"github.com/Woody1193/miletos-test/rules"
	"github.com/Woody1193/miletos-test/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"github.com/xefino/goutils/collections"
	xtime "github.com/xefino/goutils/time"
)

// Create a new test runner we'll use to test all the
// modules in the batch package
func TestBatch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Batch Suite")
}

var _ = Describe("Check Tests", func() {

	// Tests that the Check function works as expected
	It("Check - Works", func() {

		// First, create some invoice data to check
		invoiceData := collections.NewIndexedMap[string, *types.InvoiceItem]()
		invoiceData.Add("123", types.NewInvoiceItem("123",
			time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC), decimal.New(400, 0), 2), false)
		invoiceData.Add("124", types.NewInvoiceItem("124",
			time.Date(2022, 6, 4, 0, 0, 0, 0, time.UTC), decimal.New(45099, 0), 3), false)
		invoiceData.Add("125", types.NewInvoiceItem("125",
			xtime.Today().AddDate(0, 3, 0), decimal.New(30000, 0), 4), false)
		invoiceData.Add("126", types.NewInvoiceItem("126",
			time.Date(2022, 8, 4, 0, 0, 0, 0, time.UTC), decimal.New(30000, 0), 5), false)
		invoiceData.Add("128", types.NewInvoiceItem("128",
			time.Date(2022, 8, 4, 0, 0, 0, 0, time.UTC), decimal.New(98100, 0), 6), false)

		// Create some receivables data to check
		receivablesData := collections.NewIndexedMap[string, *types.ReceivablesItem]()
		receivablesData.Add("123", types.NewReceivablesItem("123",
			time.Date(2022, 7, 4, 0, 0, 0, 0, time.UTC), decimal.New(400, 0), 2), false)
		receivablesData.Add("124", types.NewReceivablesItem("124",
			time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC), decimal.New(15099, 0), 3), false)
		receivablesData.Add("125", types.NewReceivablesItem("125",
			xtime.Today().AddDate(0, 2, 0), decimal.New(30000, 0), 4), false)
		receivablesData.Add("127", types.NewReceivablesItem("127",
			time.Date(2022, 8, 4, 0, 0, 0, 0, time.UTC), decimal.New(30000, 0), 5), false)
		receivablesData.Add("128", types.NewReceivablesItem("128",
			time.Date(2022, 8, 2, 0, 0, 0, 0, time.UTC), decimal.New(98100, 0), 6), false)

		// Next, attempt to check the data
		results := NewCheckBatch(invoiceData, receivablesData,
			rules.InvoiceExists,
			rules.AmountsEqual,
			rules.PaidOnTime,
			rules.DateNotInFuture,
			rules.NotPastDue,
		).Check()

		// Finally, verify that the results are as expected
		Expect(results).Should(HaveLen(5))
		verifyCheckResult(results[0], "123", 2, 2,
			"Invoice due date of 2022-05-04 00:00:00 +0000 UTC is before receivables date "+
				"of 2022-07-04 00:00:00 +0000 UTC")
		verifyCheckResult(results[1], "124", 3, 3,
			"Invoice amount of 45099 does not match receivables amount of 15099")
		verifyCheckResult(results[2], "125", 4, 4,
			"Receivables date of 2023-06-22 00:00:00 +0000 UTC is more than one month in the future")
		verifyCheckResult(results[3], "127", 0, 5, "Invoice does not exist")
		verifyCheckResult(results[4], "126", 5, 0,
			"Invoice due date of 2022-08-04 00:00:00 +0000 UTC has past")
	})
})

// Helper function to verify the results of a check
func verifyCheckResult(result *types.CheckResult, id string, invLine uint, recLine uint, reason string) {
	Expect(result.ID).Should(Equal(id))
	Expect(result.InvoicesFileLine).Should(Equal(invLine))
	Expect(result.ReceivablesFileLine).Should(Equal(recLine))
	Expect(result.Description).Should(Equal(reason))
}
