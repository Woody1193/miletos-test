package types

// CheckResult is the result of a single check of an invoice and receivables item
// that fails a rule
type CheckResult struct {
	InvoicesFileLine    uint   `json:"Invoices File Line" csv:"Invoices File Line"`
	ReceivablesFileLine uint   `json:"Receivables File Line" csv:"Receivables File Line"`
	ID                  string `json:"ID" csv:"ID"`
	Description         string `json:"Description" csv:"Description"`
}

// ErrorResult is the result of a single error that occurred while reading a file
type ErrorResult struct {
	File  string `json:"File" csv:"File"`
	Line  uint   `json:"Line" csv:"Line"`
	Error string `json:"Error" csv:"Error"`
}

// NewErrorResult creates a new ErrorResult from a file name, line number, and error
func NewErrorResult(file string, line uint, err error) *ErrorResult {
	return &ErrorResult{
		File:  file,
		Line:  line,
		Error: err.Error(),
	}
}
