package io

import (
	"os"

	"github.com/gocarina/gocsv"
)

// WriteCsv writes a slice of items to a CSV file.
func WriteCsv[TItem any](path string, items ...TItem) error {

	// First, attempt to create the file; if this fails, return the error.
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	// Next, ensure that the file is closed when the function exits so we don't
	// leak file descriptors.
	defer file.Close()

	// Now, create a CSV writer that uses tabs as the delimiter.
	writer := gocsv.DefaultCSVWriter(file)
	writer.Comma = '\t'

	// Finally, attempt to write the items to the file; if this fails, return the
	// error.
	return gocsv.MarshalCSV(items, writer)
}
