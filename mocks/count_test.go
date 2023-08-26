package main

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

type SpyOperationCount struct {
	Calls []string
}

func (s *SpyOperationCount) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyOperationCount) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCount(t *testing.T) {
	t.Run("countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeperSpy := &SpyOperationCount{}

		Count(buffer, sleeperSpy)

		out := buffer.String()
		want := `3
2
1
Go!`
		assert.Equal(t, out, want)
	})

	t.Run("write before every sleep", func(t *testing.T) {
		SpyOperationCount := &SpyOperationCount{}
		Count(SpyOperationCount, SpyOperationCount)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		assert.DeepEqual(t, SpyOperationCount.Calls, want)
	})
}

const (
	write = "write"
	sleep = "sleep"
)
