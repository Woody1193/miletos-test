package types

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// Contains a list of errors that can be returned by the Verify function
var (
	EmptyIDError  = fmt.Errorf("ID was empty")
	NoDateError   = fmt.Errorf("Date was empty")
	NoAmountError = fmt.Errorf("Amount was empty")
)

// InvoiceItem is a struct that contains the data for an invoice item
type InvoiceItem struct {
	ID      string          `json:"ID" csv:"ID"`
	DueDate time.Time       `json:"Due Date" csv:"Due Date"`
	Amount  decimal.Decimal `json:"Amount" csv:"Amount"`
}

// Key returns the ID of the invoice item which uniquely identifies it
func (i *InvoiceItem) Key() string {
	return i.ID
}

// Verify checks that the invoice item is valid and returns an error if it is not
func (i *InvoiceItem) Verify() error {

	// First, check that the ID is not empty; if it is, return an error
	if i.ID == "" {
		return EmptyIDError
	}

	// Next, check that the due date is not empty; if it is, return an error
	if i.DueDate == (time.Time{}) {
		return NoDateError
	}

	// Finally, check that the amount is not empty and that it is an integer
	// value; if it is not, return an error
	if i.Amount.IsZero() {
		return NoAmountError
	} else if i.Amount.Exponent() < 0 {
		return fmt.Errorf("Amount of %s was invalid", i.Amount)
	}

	return nil
}

// ReceivablesItem is a struct that contains the data for a receivables item
type ReceivablesItem struct {
	ID     string          `json:"ID" csv:"ID"`
	Date   time.Time       `json:"Date" csv:"Date"`
	Amount decimal.Decimal `json:"Amount" csv:"Amount"`
}

// Key returns the ID of the receivables item which uniquely identifies it
func (r *ReceivablesItem) Key() string {
	return r.ID
}

// Verify checks that the receivables item is valid and returns an error if it is
func (i *ReceivablesItem) Verify() error {

	// First, check that the ID is not empty; if it is, return an error
	if i.ID == "" {
		return EmptyIDError
	}

	// Next, check that the date is not empty; if it is, return an error
	if i.Date == (time.Time{}) {
		return NoDateError
	}

	// Finally, check that the amount is not empty and that it is an integer
	// value; if it is not, return an error
	if i.Amount.IsZero() {
		return NoAmountError
	} else if i.Amount.Exponent() < 0 {
		return fmt.Errorf("Amount of %s was invalid", i.Amount)
	}

	return nil
}
