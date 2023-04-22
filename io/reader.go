package io

import (
	"fmt"
	"io"

	"github.com/Woody1193/miletos-test/types"
	"github.com/gocarina/gocsv"
	"github.com/xefino/goutils/collections"
)

// Keyer is an interface that must be implemented by all types that are to be
// read from a CSV file. The Key() method is used to determine if the item is
// a duplicate, and the Verify() method is used to determine if the item is
// valid.
type Keyer[TKey comparable] interface {

	// Key returns the key that is used to determine if the item is a duplicate.
	Key() TKey

	// Verify returns an error if the item is invalid.
	Verify() error

	// SetLine sets the line number associated with the item.
	SetLine(uint)
}

// Setup the CSV reader to fail if there are duplicate headers or if there are
// unmatched struct tags.
func init() {
	gocsv.FailIfDoubleHeaderNames = true
	gocsv.FailIfUnmatchedStructTags = true
}

// ReadCSV reads a CSV file and returns a map of the items in the file. The
// path parameter is the path to the CSV file. This function requires two type
// parameters: TKey and TItem. TKey is the type of the key that is used to map
// the items in the file. TItem is the type of the items in the file. TItem must
// implement the Keyer[TKey] interface. The first return value is a map of the
// items in the file. The second return value is a slice of errors that were
// encountered while parsing the file. The third return value is an error that
// was encountered while reading the file.
func ReadCSV[TKey comparable, TItem Keyer[TKey]](reader io.Reader,
	path string) (*collections.IndexedMap[TKey, TItem], []*types.ErrorResult, error) {
	eh := new(ErrorHandler)

	// First, attempt to parse the file; if this fails, return the error. Although we
	// use an error handler, we still need to check the error return value because
	// gocsv can still return an error if the file is empty or the headers could not be read
	data := make([]TItem, 0)
	if err := gocsv.UnmarshalWithErrorHandler(reader, eh.HandleParseError, &data); err != nil {
		return nil, nil, err
	}

	// Next, iterate over the errors that were encountered while parsing the file and
	// convert them to ErrorResult objects.
	results := collections.NewIndexedMap[uint, *types.ErrorResult]()
	for _, err := range eh.ParseErrors {
		results.Add(uint(err.Line), types.NewErrorResult(path, uint(err.Line), err), false)
	}

	// Now, iterate over the items in the file and add them to the map. If an item
	// is a duplicate, record an error. Otherwise, if the item is invalid, record
	// an error.
	mapping := collections.NewIndexedMap[TKey, TItem]()
	for i, item := range data {

		// First, get the key for the item
		key := item.Key()

		// Next, get a copy of the item and then attempt to verify it. If the item
		// is invalid, record an error and continue to the next item.
		copy := item
		if err := copy.Verify(); err != nil {
			results.AddIf(uint(i+2), types.NewErrorResult(path, uint(i+2), err),
				func(existing *types.ErrorResult, newItem *types.ErrorResult) bool {
					existing.Error += "; " + newItem.Error
					return false
				})

			continue
		}

		// Now, set the line associated with this item
		copy.SetLine(uint(i + 2))

		// Finally, attempt to add the item to the map. If the item is a duplicate,
		// record an error and continue to the next item.
		mapping.AddIf(key, copy, func(existing TItem, newItem TItem) bool {
			results.Add(uint(i+2), types.NewErrorResult(path, uint(i+2),
				fmt.Errorf("Duplicate invoice ID of %v detected", key)), false)
			return false
		})
	}

	// Finally, return the map of items and the errors that were encountered.
	return mapping, results.Data(), nil
}
