package io

import (
	"bytes"
	"time"

	"github.com/Woody1193/miletos-test/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

var _ = Describe("Writer Tests", func() {

	// Tests that the WriteCsv function works as expected
	It("WriteCsv - Works", func() {

		// First, create some data to write
		data := []*types.InvoiceItem{
			types.NewInvoiceItem("123", time.Date(2022, time.May, 4, 0, 0, 0, 0, time.UTC), decimal.New(400, 0)),
			types.NewInvoiceItem("124", time.Date(2022, time.June, 4, 0, 0, 0, 0, time.UTC), decimal.New(45099, 0)),
		}

		// Next, attempt to write to a nonexistent CSV file; this should not fail
		buffer := new(bytes.Buffer)
		err := WriteCsv(buffer, data...)

		// Finally, verify that the data was written
		Expect(err).ShouldNot(HaveOccurred())
		Expect(string(buffer.Bytes())).Should(Equal("ID,Amount,Due Date\n" +
			"123,400,2022-05-04T00:00:00Z\n124,450.99,2022-06-04T00:00:00Z\n"))
	})
})
