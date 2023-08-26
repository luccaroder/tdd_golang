package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	lastWord  = "Go!"
	countDown = 3
)

type Sleeper interface {
	Sleep()
}
type SleeperDefault struct{}

func (d *SleeperDefault) Sleep() {
	time.Sleep(1 * time.Second)
}

func Count(out io.Writer, sleeper Sleeper) {
	for i := countDown; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintf(out, "%d\n", i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, lastWord)
}

func main() {
	sleeper := &SleeperDefault{}
	Count(os.Stdout, sleeper)
}
