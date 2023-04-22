package io

import (
	"io"

	"github.com/gocarina/gocsv"
)

// WriteCsv writes a slice of items to a CSV file.
func WriteCsv[TItem any](writer io.Writer, items ...TItem) error {
	csvWriter := gocsv.DefaultCSVWriter(writer)
	csvWriter.Comma = '\t'
	return gocsv.MarshalCSV(items, csvWriter)
}
