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

// sleep field is a function that can be replaced by
type SleeperCustom struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (d *SleeperDefault) Sleep() {
	time.Sleep(1 * time.Second)
}

func (d *SleeperCustom) Sleep() {
	d.sleep(d.duration)
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
	sleeper := &SleeperCustom{1 * time.Second, time.Sleep}
	Count(os.Stdout, sleeper)
}
