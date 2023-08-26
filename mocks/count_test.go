package main

import (
	"bytes"
	"testing"
	"time"

	"gotest.tools/assert"
)

type SpyOperationCount struct {
	Calls []string
}

// stub to grant that curtom sleeper is called
type TimeSpy struct {
	sleepDuration time.Duration
}

func (s *SpyOperationCount) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyOperationCount) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// sleep has function to simulate time.Sleep
func (t *TimeSpy) Sleep(duration time.Duration) {
	t.sleepDuration = duration
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

func TestCustomSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &TimeSpy{}
	sleeper := &SleeperCustom{sleepTime, spyTime.Sleep}
	sleeper.Sleep()
	assert.Equal(t, spyTime.sleepDuration, sleepTime)
}

const (
	write = "write"
	sleep = "sleep"
)
