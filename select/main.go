package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/** Testing fast URL without concurrence*/

// func Running(a, b string) string {
// 	durantionA := measureResponseTime(a)
// 	durantionB := measureResponseTime(b)

// 	if durantionA < durantionB {
// 		return a
// 	}
// 	return b
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()

// 	_, err := http.Get(url)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return time.Since(start)
// }

var ErrTimeOut = fmt.Errorf("timed out waiting for servers to respond")

func Running(a, b string, limit time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(limit):
		return "", ErrTimeOut
	}
}

func ping(url string) chan bool {
	channel := make(chan bool)

	go func() {
		if _, err := http.Get(url); err != nil {
			log.Printf("Err to get url %v", err)
		}
		channel <- true
	}()

	return channel
}
