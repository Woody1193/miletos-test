package rules

import (
	"fmt"
	"time"

	"github.com/Woody1193/miletos-test/types"
)

// NotPastDue compares the due date of the invoice and returns an error if the
// due date has passed. This function expects that invoiceItem is not nil and that
// the due date is not nil.
func NotPastDue(invoiceItem *types.InvoiceItem, receivablesItem *types.ReceivablesItem) (bool, error) {
	if receivablesItem == nil && invoiceItem.DueDate.Before(time.Now()) {
		return false, fmt.Errorf("Invoice due date of %s has past", invoiceItem.DueDate)
	}

	return true, nil
}

// PaidOnTime compares the due date of the invoice and receivables items and
// returns an error if the receivables date is after the invoice due date. This
// function expects that, if invoiceItem and receivablesItem are not nil, then
// the due date is not nil and the receivables date is not nil.
func PaidOnTime(invoiceItem *types.InvoiceItem, receivablesItem *types.ReceivablesItem) (bool, error) {
	if receivablesItem != nil && invoiceItem != nil && invoiceItem.DueDate.Before(*receivablesItem.Date) {
		return false, fmt.Errorf("Invoice due date of %s is before receivables date of %s",
			invoiceItem.DueDate, receivablesItem.Date)
	}

	return true, nil
}

// DateNotInFuture compares the date of the receivables item and returns an error
// if the date is more than one month in the future. This function expects that,
// if receivablesItem is not nil, then the date is not nil.
func DateNotInFuture(invoiceItem *types.InvoiceItem, receivablesItem *types.ReceivablesItem) (bool, error) {
	if receivablesItem != nil && receivablesItem.Date.After(time.Now().AddDate(0, 1, 0)) {
		return false, fmt.Errorf("Receivables date of %s is more than one month in the future",
			receivablesItem.Date)
	}

	return true, nil
}
