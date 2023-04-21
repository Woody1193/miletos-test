package rules

import "github.com/Woody1193/miletos-test/types"

type Rule func(*types.InvoiceItem, *types.ReceivablesItem) (bool, error)
