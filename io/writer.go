package io

import (
	"os"

	"github.com/gocarina/gocsv"
)

func WriteCsv[TItem any](path string, items ...TItem) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return gocsv.MarshalFile(items, file)
}
