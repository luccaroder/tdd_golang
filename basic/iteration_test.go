package tddgo

import "testing"

func TestIteration(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("esperado '%s', resultado '%s'", expected, repeated)
	}
}
