package io

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/xefino/goutils/collections"
)

type Keyer[TKey comparable, TItem any] interface {
	Key() (TKey, *TItem)

	DuplicateKeyError(*TItem) error
}

func ReadCSV[TKey comparable, TItem any](path string, handler *ErrorHandler,
	keyer Keyer[TKey, TItem]) (*collections.IndexedMap[TKey, *TItem], error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data := make([]Keyer[TKey, TItem], 0)
	if err := gocsv.UnmarshalWithErrorHandler(file, handler.HandleParseError, &data); err != nil {
		return nil, err
	}

	mapping := collections.NewIndexedMap[TKey, *TItem]()
	for _, keyer := range data {

		key, item := keyer.Key()

		mapping.AddIf(key, item, func(existing *TItem, newItem *TItem) bool {
			handler.HandleFormatError(keyer.DuplicateKeyError(newItem))
			return false
		})
	}

	return mapping, nil
}
