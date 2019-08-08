package cli_test

import (
	"fmt"
	cli2 "github.com/valdemarceccon/golang-tdd-study/app_poker/cmd/cli"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/player"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/pokertesting"
	"strings"
	"testing"
	"time"
)

type SpyBlindAlerter struct {
	alerts []struct {
		scheduledAt time.Duration
		amount      int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		scheduledAt time.Duration
		amount      int
	}{scheduledAt: duration, amount: amount})
}

func TestCLI(t *testing.T) {
	var dummySpyAlerter = &SpyBlindAlerter{}

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &player.StubPlayerStore{}

		cli := cli2.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		pokertesting.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &player.StubPlayerStore{}

		cli := cli2.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		pokertesting.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &player.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := cli2.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []struct {
			expectedScheduleTime time.Duration
			expectedAmount       int
		}{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d scheduled for %v", c.expectedAmount, c.expectedScheduleTime), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d wat not scheduled %v", i, blindAlerter.alerts)
				}

				alert := blindAlerter.alerts[i]

				amountGot := alert.amount

				if amountGot != c.expectedAmount {
					t.Fatalf("got amount %d, want %v", amountGot, c.expectedAmount)
				}

				gotScheduleTime := alert.scheduledAt

				if gotScheduleTime != c.expectedScheduleTime {
					t.Errorf("got schedule time of %v, want %v", gotScheduleTime, c.expectedScheduleTime)
				}
			})
		}
	})
}
