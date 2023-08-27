package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestIteration(t *testing.T) {
	scenarios := []struct {
		Name     string
		input    interface{}
		expected []string
	}{
		{
			Name:     "test iteration one string field",
			input:    struct{ Name string }{Name: "Lucca"},
			expected: []string{"Lucca"},
		},
		{
			Name: "test iteration more string fields",
			input: struct {
				Name string
				City string
			}{Name: "Lucca", City: "São Paulo"},
			expected: []string{"Lucca", "São Paulo"},
		},
		{
			Name: "test iteration with int and string fields",
			input: struct {
				Name string
				Age  int
			}{Name: "Lucca", Age: 33},
			expected: []string{"Lucca"},
		},
		{
			Name: "test iteration with int and string fields",
			input: struct {
				Name    string
				Age     int
				Profile struct {
					Email string
				}
			}{Name: "Lucca", Age: 33, Profile: struct{ Email string }{Email: "lucca@test.com"}},
			expected: []string{"Lucca", "lucca@test.com"},
		},
	}

	for _, s := range scenarios {
		t.Run(s.Name, func(t *testing.T) {
			var result []string

			Iteration(s.input, func(param string) {
				result = append(result, param)
			})

			assert.DeepEqual(t, s.expected, result)
		})
	}
}
