package rules

import (
	"fmt"
	"time"

	"github.com/Woody1193/miletos-test/types"
)

func PaidOnTime(invoiceItem *types.InvoiceItem, receivablesItem *types.ReceivablesItem) (bool, error) {

	if invoiceItem.DueDate.After(*receivablesItem.Date) {
		return false, fmt.Errorf("Invoice due date of %s is after receivables date of %s",
			invoiceItem.DueDate, receivablesItem.Date)
	}

	return true, nil
}

func DateNotInFuture(invoiceItem *types.InvoiceItem, receivablesItem *types.ReceivablesItem) (bool, error) {

	if receivablesItem.Date.After(time.Now().AddDate(0, 1, 0)) {
		return false, fmt.Errorf("Receivables date of %s is more than one month in the future",
			receivablesItem.Date)
	}

	return true, nil
}
