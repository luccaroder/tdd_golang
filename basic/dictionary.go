package tddgo

import "fmt"

type Dictionary map[string]string

var (
	ErrNotFound   = fmt.Errorf("could not find the word you were looking for")
	ErrWordExists = fmt.Errorf("cannot add word because it already exists")
	ErrNotExists  = fmt.Errorf("cannot update word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	if _, ok := d[word]; !ok {
		return "", ErrNotFound
	}
	return d[word], nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrNotExists
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
