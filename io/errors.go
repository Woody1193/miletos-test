package io

import "encoding/csv"

type ErrorHandler struct {
	ParseErrors []*csv.ParseError
}

func (eh *ErrorHandler) HandleParseError(inner *csv.ParseError) bool {
	eh.ParseErrors = append(eh.ParseErrors, inner)
	return true
}
