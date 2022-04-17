package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string

// 아래와 똑같이 동작
//var errNotFound = errors.New("not Found")
//var errWordExists = errors.New("that word already exists")
//var errCantUpdate = errors.New("cant update non-existing word")

var (
	errNotFound   = errors.New("not Found")
	errWordExists = errors.New("that word already exists")
	errCantUpdate = errors.New("cant update non-existing word")
)

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

// Update definition
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		d[word] = def
	}
	return nil
}

// Delete word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
