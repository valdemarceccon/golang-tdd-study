package countdown_test

import (
	"bytes"
	countdown "github.com/valdemarceccon/golang-tdd-study/fundamentals/09_mocking"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	const (
		sleep = "sleep"
		write = "write"
	)

	t.Run("countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		countdown.Countdown(buffer, &countdown.OperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &countdown.OperationsSpy{}
		countdown.Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep, write,
			sleep, write,
			sleep, write,
			sleep, write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &countdown.SpyTime{}

	sleeper := countdown.ConfigurableSleeper{
		Duration: sleepTime,
		SleepH:   spyTime.Sleep,
	}
	sleeper.Sleep()
	if spyTime.DurationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.DurationSlept)
	}
}
