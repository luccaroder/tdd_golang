package main

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func VerificationWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		/** can cause race conditions because the goroutine is not synchronized and try write in the same memory address/object
			It's necessary create channels to synchronize the goroutines. We parallelize the execution of the goroutines and keep
			the order of the results.
		**/

		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}
