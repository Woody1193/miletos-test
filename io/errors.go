package io

import "encoding/csv"

// ErrorHandler is a helper type intended to be used with gocsv.UnmarshalWithErrorHandler
// to collect all parse errors.
type ErrorHandler struct {
	ParseErrors []*csv.ParseError
}

// HandleParseError is called when gocsv.UnmarshalWithErrorHandler encounters a parse error.
// It returns true to continue parsing, regardless, and collects the error.
func (eh *ErrorHandler) HandleParseError(inner *csv.ParseError) bool {
	eh.ParseErrors = append(eh.ParseErrors, inner)
	return true
}
