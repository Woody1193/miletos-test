package rules

import (
	"fmt"

	"github.com/Woody1193/miletos-test/types"
)

func InvoiceExists(invoiceItem *types.InvoiceItem, receivablesItem *types.ReceivablesItem) (bool, error) {

	if invoiceItem == nil {
		return false, fmt.Errorf("Invoice does not exist")
	}

	return true, nil
}
