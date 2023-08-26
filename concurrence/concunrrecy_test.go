package main

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func slowStubVerificationWebsite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkVerificationaWebsites(b *testing.B) {
	for i := 0; i < b.N; i++ {
		websites := []string{
			"http://google.com",
			"http://blog.gypsydave5.com",
			"waat://furhurterwe.geds",
		}

		expected := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    true,
		}

		out := VerificationWebsites(slowStubVerificationWebsite, websites)
		assert.DeepEqual(b, out, expected)
	}
}
