package rules

import "github.com/Woody1193/miletos-test/types"

type Rule interface {
	Apply(*types.InvoiceItem, *types.ReceivablesItem) (bool, error)
}
