package rules

import (
	"fmt"

	"github.com/Woody1193/miletos-test/types"
)

// InvoiceExists checks if the invoice item is nil and returns an error if it is nil
func InvoiceExists(invoiceItem *types.InvoiceItem, receivablesItem *types.ReceivablesItem) (bool, error) {
	if invoiceItem == nil {
		return false, fmt.Errorf("Invoice does not exist")
	}

	return true, nil
}
