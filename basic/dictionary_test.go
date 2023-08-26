package tddgo

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		want := "this is just a test"
		got, _ := dictionary.Search("test")
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})

	t.Run("add word", func(t *testing.T) {
		want := "other test"
		err := dictionary.Add("other", want)
		got, _ := dictionary.Search("other")
		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("add existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "this is just a test")
		assertError(t, err, ErrWordExists)
	})

	t.Run("update word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		want := "this is just a test updated"
		err := dictionary.Update("test", want)
		got, _ := dictionary.Search("test")
		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("update not existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Update("other", "this is just a test")
		assertError(t, err, ErrNotExists)
	})

	t.Run("delete word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Delete("test")
		assertError(t, err, nil)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	assert.Equal(t, got, want)
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	assert.Equal(t, got, want)
}
