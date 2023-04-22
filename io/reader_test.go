package io

import (
	"sort"
	"time"

	"github.com/Woody1193/miletos-test/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reader Tests", func() {

	// Tests that the ReadCSV function returns an error when
	// it is passed a nonexistent CSV file
	It("ReadCSV - Open fails - Error", func() {

		// Attempt to read from a nonexistent CSV file; this should fail
		data, results, err := ReadCSV[string, *types.InvoiceItem]("nonexistent.csv")

		// Verify that we retrieved no data and no results and that an error was returned
		Expect(data).Should(BeNil())
		Expect(results).Should(BeNil())
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("open nonexistent.csv: no such file or directory"))
	})

	// Tests that the ReadCSV function returns an error when
	// it is passed a CSV file that is not properly formatted
	It("ReadCSV - UnmarshalWithErrorHandler fails - Error", func() {

		// Attempt to read from a CSV file that will fail to unmarshal; this should fail
		data, results, err := ReadCSV[string, *types.InvoiceItem]("failure.csv")

		// Verify that we retrieved no data and no results and that an error was returned
		Expect(data).Should(BeNil())
		Expect(results).Should(BeNil())
		Expect(err).Should(HaveOccurred())
		Expect(err.Error()).Should(Equal("found unmatched struct field with tags [Due Date]"))
	})

	// Tests that the ReadCSV function returns the expected data
	It("ReadCSV - Success", func() {

		// Attempt to read from a CSV file that will fail to unmarshal; this should not fail
		data, results, err := ReadCSV[string, *types.InvoiceItem]("success.csv")
		Expect(err).ShouldNot(HaveOccurred())

		// Verify that we retrieved the expected data
		Expect(data.Length()).Should(Equal(3))
		verifyItem(data.At(0), "128", "700", time.Date(2022, time.May, 4, 0, 0, 0, 0, time.UTC))
		verifyItem(data.At(1), "129", "1000", time.Date(2022, time.June, 4, 0, 0, 0, 0, time.UTC))
		verifyItem(data.At(2), "130", "250", time.Date(2022, time.May, 2, 0, 0, 0, 0, time.UTC))

		// Before we verify the results, we need to sort them by file and line number so that
		// we can ensure the test either passes or fails consistently
		sort.Slice(results, func(i, j int) bool {
			return results[i].File < results[j].File || results[i].Line < results[j].Line
		})

		// Verify that we retrieved the expected results
		Expect(results).Should(HaveLen(7))
		verifyErrorResult(results[0], "success.csv", 2, "ID was empty")
		verifyErrorResult(results[1], "success.csv", 3, "record on line 0; parse error on line 3, column 2: "+
			"error decoding string '': can't convert  to decimal; Amount was empty")
		verifyErrorResult(results[2], "success.csv", 4, "record on line 0; parse error on line 4, column 2: "+
			"error decoding string 'derp': can't convert derp to decimal: exponent is not numeric; Amount was empty")
		verifyErrorResult(results[3], "success.csv", 5, "Amount of 450.99 was invalid")
		verifyErrorResult(results[4], "success.csv", 6, "record on line 0; parse error on line 6, column 3: "+
			"parsing time \"\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"\" as \"2006\"; Date was empty")
		verifyErrorResult(results[5], "success.csv", 7, "record on line 0; parse error on line 7, column 3: "+
			"parsing time \"derp\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"derp\" as \"2006\"; Date was empty")
		verifyErrorResult(results[6], "success.csv", 9, "Duplicate invoice ID of 128 detected")
	})
})

// Helper function to verify that an invoice item is correct
func verifyItem(item *types.InvoiceItem, id string, amount string, date time.Time) {
	Expect(item).ShouldNot(BeNil())
	Expect(item.ID).Should(Equal(id))
	Expect(item.Amount).ShouldNot(BeNil())
	Expect(item.Amount.String()).Should(Equal(amount))
	Expect(item.DueDate).ShouldNot(BeNil())
	Expect(item.DueDate).Should(Equal(date))
}

// Helper function to verify that an error result is correct
func verifyErrorResult(result *types.ErrorResult, file string, line uint, err string) {
	Expect(result).ShouldNot(BeNil())
	Expect(result.File).Should(Equal(file))
	Expect(result.Line).Should(Equal(line))
	Expect(result.Error).Should(Equal(err))
}
