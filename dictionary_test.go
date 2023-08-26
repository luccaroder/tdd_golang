package tddgo

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	type errorTestCases struct {
		description string
		word        string
		want        string
		wantErr     error
	}

	for _, tt := range []errorTestCases{
		{
			description: "known word",
			word:        "test",
			want:        "this is just a test",
			wantErr:     nil,
		},
		{
			description: "unknown word",
			word:        "unknown",
			want:        "",
			wantErr:     ErrNotFound,
		},
	} {
		t.Run(tt.description, func(t *testing.T) {
			got, err := dictionary.Search(tt.word)
			assertStrings(t, got, tt.want)
			assertError(t, err, tt.wantErr)
		})
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	assert.Equal(t, got, want)
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	assert.Equal(t, got, want)
}
