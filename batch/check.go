package batch

import (
	"github.com/Woody1193/miletos-test/rules"
	"github.com/Woody1193/miletos-test/types"
	"github.com/xefino/goutils/collections"
)

type CheckBatch struct {
	invoiceData     *collections.IndexedMap[string, *types.InvoiceItem]
	receivablesData *collections.IndexedMap[string, *types.ReceivablesItem]
	rules           []rules.Rule
}

func NewCheckBatch(invoiceData *collections.IndexedMap[string, *types.InvoiceItem],
	receivablesData *collections.IndexedMap[string, *types.ReceivablesItem], rules ...rules.Rule) *CheckBatch {
	return &CheckBatch{
		invoiceData:     invoiceData,
		receivablesData: receivablesData,
		rules:           rules,
	}
}

func (cb *CheckBatch) Check() ([]*types.CheckResult, error) {

}
