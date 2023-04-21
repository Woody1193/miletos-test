package types

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

var (
	EmptyIDError = fmt.Errorf("ID was empty")
	NoDueDate    = fmt.Errorf("Due Date was empty")
	NoAmount     = fmt.Errorf("Amount was empty")
)

type InvoiceItem struct {
	ID      string           `json:"ID" csv:"ID"`
	DueDate *time.Time       `json:"Due Date" csv:"Due Date"`
	Amount  *decimal.Decimal `json:"Amount" csv:"Amount"`
}

func (i *InvoiceItem) Key() string {
	return i.ID
}

func (i *InvoiceItem) Verify() error {

	if i.ID == "" {
		return EmptyIDError
	}

	if i.DueDate == nil {
		return NoDueDate
	}

	if i.Amount == nil {
		return NoAmount
	} else if i.Amount.Exponent() < 0 {
		return fmt.Errorf("Amount of %s was invalid", i.Amount)
	}

	return nil
}

type ReceivablesItem struct {
	ID     string           `json:"ID" csv:"ID"`
	Date   *time.Time       `json:"Date" csv:"Date"`
	Amount *decimal.Decimal `json:"Amount" csv:"Amount"`
}

func (r *ReceivablesItem) Key() string {
	return r.ID
}

func (i *ReceivablesItem) Verify() error {

	if i.ID == "" {
		return EmptyIDError
	}

	if i.Date == nil {
		return NoDueDate
	}

	if i.Amount == nil {
		return NoAmount
	} else if i.Amount.Exponent() < 0 {
		return fmt.Errorf("Amount of %s was invalid", i.Amount)
	}

	return nil
}
