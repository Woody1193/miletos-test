package io

import "encoding/csv"

type ErrorHandler struct {
	ParseErrors  []*csv.ParseError
	FormatErrors []error
}

func (eh *ErrorHandler) HandleParseError(inner *csv.ParseError) bool {
	eh.ParseErrors = append(eh.ParseErrors, inner)
	return true
}

func (eh *ErrorHandler) HandleFormatError(inner error) {
	eh.FormatErrors = append(eh.FormatErrors, inner)
}
