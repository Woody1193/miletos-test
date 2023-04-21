package types

import (
	"time"

	"github.com/shopspring/decimal"
)

type InvoiceItem struct {
	ID      string           `json:"ID" csv:"ID"`
	DueDate *time.Time       `json:"Due Date" csv:"Due Date"`
	Amount  *decimal.Decimal `json:"Amount" csv:"Amount"`
}

func (i *InvoiceItem) Key() string {
	return i.ID
}

type ReceivablesItem struct {
	ID     string           `json:"ID" csv:"ID"`
	Date   *time.Time       `json:"Date" csv:"Date"`
	Amount *decimal.Decimal `json:"Amount" csv:"Amount"`
}

func (r *ReceivablesItem) Key() string {
	return r.ID
}
