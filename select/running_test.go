package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestRunning(t *testing.T) {
	t.Run("returns the faster URL", func(t *testing.T) {
		serverA := makeDelayedServer(10 * time.Millisecond)
		serverB := makeDelayedServer(0 * time.Millisecond)

		// Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.
		defer serverA.Close()
		defer serverB.Close()

		slowURL := serverA.URL
		fastURL := serverB.URL

		expected := fastURL
		result, err := Running(slowURL, fastURL, 100*time.Millisecond)

		assert.NilError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(15 * time.Millisecond)

		// Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.
		defer server.Close()

		_, err := Running(server.URL, server.URL, 10*time.Millisecond)

		assert.ErrorContains(t, err, "timed out waiting for servers to respond")
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
