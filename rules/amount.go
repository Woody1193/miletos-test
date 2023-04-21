package rules

import (
	"fmt"

	"github.com/Woody1193/miletos-test/types"
)

// AmountsEqual compares the amount of the invoice and receivables items and
// returns an error if they do not match
func AmountsEqual(invoiceItem *types.InvoiceItem, receivablesItem *types.ReceivablesItem) (bool, error) {
	if receivablesItem != nil && receivablesItem.Amount.Cmp(*invoiceItem.Amount) != 0 {
		return false, fmt.Errorf("Invoice amount of %s does not match receivables amount of %s",
			invoiceItem.Amount, receivablesItem.Amount)
	}

	return true, nil
}
