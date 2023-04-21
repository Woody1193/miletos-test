package types

type CheckResult struct {
	InvoicesFileLine    uint   `json:"Invoices File Line" csv:"Invoices File Line"`
	ReceivablesFileLine uint   `json:"Receivables File Line" csv:"Receivables File Line"`
	ID                  string `json:"ID" csv:"ID"`
	Description         string `json:"Description" csv:"Description"`
}

type ErrorResult struct {
	File  string `json:"File" csv:"File"`
	Line  uint   `json:"Line" csv:"Line"`
	Error string `json:"Error" csv:"Error"`
}

func NewErrorResult(file string, line uint, err error) *ErrorResult {
	return &ErrorResult{
		File:  file,
		Line:  line,
		Error: err.Error(),
	}
}
