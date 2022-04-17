package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search foir a word in Dictionary
func (d Dictionary) Search(word string) (string, error) {

	value, exists := d[word]

	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d *Dictionary) SearchForStudy(word string) (string, error) {
	// https://stackoverflow.com/questions/36463608/go-invalid-operation-type-mapkeyvalue-does-not-support-indexing
	value, exists := (*d)[word]

	if exists {
		return value, nil
	}
	return "", errNotFound
}
