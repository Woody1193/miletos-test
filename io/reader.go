package io

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/xefino/goutils/collections"
)

type Keyer[TKey comparable] interface {
	Key() TKey
}

func ReadCSV[TKey comparable, TItem Keyer[TKey]](path string,
	handler *ErrorHandler) (*collections.IndexedMap[TKey, TItem], error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data := make([]TItem, 0)
	if err := gocsv.UnmarshalWithErrorHandler(file, handler.HandleParseError, &data); err != nil {
		return nil, err
	}

	mapping := collections.NewIndexedMap[TKey, TItem]()
	for _, item := range data {

		key := item.Key()

		copy := item
		mapping.AddIf(key, copy, func(existing TItem, newItem TItem) bool {
			handler.HandleFormatError(fmt.Errorf("Duplicate invoice ID of %v detected", key))
			return false
		})
	}

	return mapping, nil
}
