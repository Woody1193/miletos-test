package rules

import "github.com/Woody1193/miletos-test/types"

// Rule is a function that takes an invoice and receivables item and returns
// whether the rule passed and an error if the rule failed
type Rule func(*types.InvoiceItem, *types.ReceivablesItem) (bool, error)
