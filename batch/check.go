package batch

import (
	"strings"

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

func (cb *CheckBatch) Check() []*types.CheckResult {

	results := make([]*types.CheckResult, 0)

	cb.receivablesData.ForEach(func(key string, value *types.ReceivablesItem) bool {

		invoiceItem, _ := cb.invoiceData.Get(key)

		errs := make([]error, 0)
		for _, rule := range cb.rules {

			_, err := rule.Apply(invoiceItem, value)
			if err != nil {
				errs = append(errs, err)
			}
		}

		if len(errs) > 0 {

			result := types.CheckResult{
				ID:          value.ID,
				Description: strings.Join(collections.Convert(func(err error) string { return err.Error() }, errs...), "; "),
			}

			invIndex, ok := cb.invoiceData.Index(key)
			if ok {
				result.InvoicesFileLine = invIndex + 1
			}

			recIndex, ok := cb.receivablesData.Index(key)
			if ok {
				result.ReceivablesFileLine = recIndex + 1
			}

			results = append(results, &result)
		}

		return true
	})

	return results
}
