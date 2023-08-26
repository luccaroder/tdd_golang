package main

import (
	"bytes"
	"testing"

	"gotest.tools/v3/assert"
)

func TestHello(t *testing.T) {
	buffer := bytes.Buffer{}
	Hello(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"
	assert.Equal(t, got, want)
}
