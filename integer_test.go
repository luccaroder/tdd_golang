package tddgo

import "testing"

func TestIntegerAdd(t *testing.T) {
	sum := IntegerAdd(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("esperado '%d', resultado '%d'", sum, expected)
	}
}
