package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("not Found")
var errWordExists = errors.New("that word already exists")

// Search foir a word in Dictionary
func (d Dictionary) Search(word string) (string, error) {

	value, exists := d[word]

	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word into a dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}
