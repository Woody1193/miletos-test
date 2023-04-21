package io

import (
	"fmt"
	"os"

	"github.com/Woody1193/miletos-test/types"
	"github.com/gocarina/gocsv"
	"github.com/xefino/goutils/collections"
)

type Keyer[TKey comparable] interface {
	Key() TKey

	Verify() error
}

func ReadCSV[TKey comparable, TItem Keyer[TKey]](path string) (*collections.IndexedMap[TKey, TItem], []*types.ErrorResult, error) {
	eh := new(ErrorHandler)

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	data := make([]TItem, 0)
	if err := gocsv.UnmarshalWithErrorHandler(file, eh.HandleParseError, &data); err != nil {
		return nil, nil, err
	}

	results := make([]*types.ErrorResult, len(eh.ParseErrors))
	for i, err := range eh.ParseErrors {
		results[i] = types.NewErrorResult(path, uint(err.Line), err)
	}

	mapping := collections.NewIndexedMap[TKey, TItem]()
	for i, item := range data {

		key := item.Key()

		copy := item
		if err := copy.Verify(); err != nil {
			results[i] = types.NewErrorResult(path, uint(i+1), err)
			continue
		}

		mapping.AddIf(key, copy, func(existing TItem, newItem TItem) bool {
			results[i] = types.NewErrorResult(path, uint(i+1),
				fmt.Errorf("Duplicate invoice ID of %v detected", key))
			return false
		})
	}

	return mapping, results, nil
}
