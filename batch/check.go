package batch

import (
	"strings"

	"github.com/Woody1193/miletos-test/rules"
	"github.com/Woody1193/miletos-test/types"
	"github.com/xefino/goutils/collections"
)

// CheckBatch checks that the invoice and receivables items match the rules provided
type CheckBatch struct {
	invoiceData     *collections.IndexedMap[string, *types.InvoiceItem]
	receivablesData *collections.IndexedMap[string, *types.ReceivablesItem]
	rules           []rules.Rule
}

// NewCheckBatch creates a new CheckBatch with the provided invoice and receivables data and rules
func NewCheckBatch(invoiceData *collections.IndexedMap[string, *types.InvoiceItem],
	receivablesData *collections.IndexedMap[string, *types.ReceivablesItem], rules ...rules.Rule) *CheckBatch {
	return &CheckBatch{
		invoiceData:     invoiceData,
		receivablesData: receivablesData,
		rules:           rules,
	}
}

// Check runs the rules on the invoice and receivables data and returns the results
func (cb *CheckBatch) Check() []*types.CheckResult {
	results := make([]*types.CheckResult, 0)

	// Iterate over the receivables data and check the rules against the invoice data
	cb.receivablesData.ForEach(func(key string, value *types.ReceivablesItem) bool {

		// Find the invoice item associated with the receivables item
		invoiceItem, _ := cb.invoiceData.Get(key)

		// Next, run the rules against the invoice and receivables item
		// If any errors are returned, aggregate them for later
		errs := make([]error, 0)
		for _, rule := range cb.rules {
			_, err := rule(invoiceItem, value)
			if err != nil {
				errs = append(errs, err)
			}
		}

		// Finally, if any errors were returned, create a CheckResult and add it to the results
		if len(errs) > 0 {

			// First, create a CheckResult with the ID and description
			result := types.CheckResult{
				ID:          value.ID,
				Description: strings.Join(collections.Convert(func(err error) string { return err.Error() }, errs...), "; "),
			}

			// Next, find the line numbers of the invoice; if we can't find it then do nothing
			if invoiceItem != nil {
				result.InvoicesFileLine = invoiceItem.Line
			}

			// Finally, find the line numbers of the receivables and add the check result to the
			// list of results
			result.ReceivablesFileLine = value.Line
			results = append(results, &result)
		}

		return true
	})

	// Iterate over all the invoice data that does not have a receivables item and check the rules
	cb.invoiceData.ForEach(func(key string, value *types.InvoiceItem) bool {

		// If the invoice item has a receivables item, skip it as we already checked it
		if cb.receivablesData.Exists(key) {
			return true
		}

		// Iterate over the rules and check the invoice item against them
		for _, rule := range cb.rules {
			if _, err := rule(value, nil); err != nil {
				results = append(results, &types.CheckResult{
					ID:               value.ID,
					Description:      err.Error(),
					InvoicesFileLine: value.Line,
				})
			}
		}

		return true
	})

	return results
}
