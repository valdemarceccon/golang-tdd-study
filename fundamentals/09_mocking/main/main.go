package main

import (
	countdown "github.com/valdemarceccon/golang-tdd-study/fundamentals/09_mocking"
	"os"
	"time"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{
		Duration: 1 * time.Second,
		SleepH:   time.Sleep,
	}

	countdown.Countdown(os.Stdout, sleeper)
}
