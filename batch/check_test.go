package batch

import (
	"testing"
	"time"

	"github.com/Woody1193/miletos-test/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"github.com/xefino/goutils/collections"
)

// Create a new test runner we'll use to test all the
// modules in the batch package
func TestBatch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Batch Suite")
}

var _ = Describe("Check Tests", func() {

	It("Check - Works", func() {

		invoiceData := collections.NewIndexedMap[string, *types.InvoiceItem]()
		invoiceData.Add("123", types.NewInvoiceItem("123",
			time.Date(2022, 5, 4, 0, 0, 0, 0, time.UTC), decimal.New(400, 0)))
	})
})
