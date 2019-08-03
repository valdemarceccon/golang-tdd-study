package countdown

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
	sleep          = "sleep"
	write          = "write"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

type OperationsSpy struct {
	Calls []string
}

type ConfigurableSleeper struct {
	Duration time.Duration
	SleepH   func(time.Duration)
}

type SpyTime struct {
	DurationSlept time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepH(c.Duration)
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.DurationSlept = duration
}

func (s *OperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *OperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		_, _ = fmt.Fprintln(out, i)
	}

	sleeper.Sleep()

	_, _ = fmt.Fprint(out, finalWord)
}
