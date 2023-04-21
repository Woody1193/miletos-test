package types

type CheckResult struct {
	InvoicesFileLine    uint   `json:"Invoices File Line" csv:"Invoices File Line"`
	ReceivablesFileLine uint   `json:"Receivables File Line" csv:"Receivables File Line"`
	ID                  string `json:"ID" csv:"ID"`
	Description         string `json:"Description" csv:"Description"`
}
