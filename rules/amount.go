package rules

import (
	"fmt"

	"github.com/Woody1193/miletos-test/types"
)

func AmountsEqual(invoiceItem *types.InvoiceItem, receivablesItem *types.ReceivablesItem) (bool, error) {

	if invoiceItem.Amount.Cmp(*receivablesItem.Amount) != 0 {
		return false, fmt.Errorf("Invoice amount of %s does not match receivables amount of %s",
			invoiceItem.Amount, receivablesItem.Amount)
	}

	return true, nil
}
