package cli_test

import (
	"bytes"
	"fmt"
	cli2 "github.com/valdemarceccon/golang-tdd-study/app_poker/cmd/cli"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/player"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/pokertesting"
	"strings"
	"testing"
	"time"
)

type scheduledAlert struct {
	scheduledAt time.Duration
	amount      int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.scheduledAt)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{scheduledAt: duration, amount: amount})
}

func TestCLI(t *testing.T) {
	var dummySpyAlerter = &SpyBlindAlerter{}
	var dummyBlindAlerter = &SpyBlindAlerter{}
	var dummyPlayerStore = &player.StubPlayerStore{}
	var dummyStdIn = &bytes.Buffer{}
	var dummyStdOut = &bytes.Buffer{}

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &player.StubPlayerStore{}

		cli := cli2.NewCLI(playerStore, in, dummyStdOut, dummySpyAlerter)
		cli.PlayPoker()

		pokertesting.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &player.StubPlayerStore{}

		cli := cli2.NewCLI(playerStore, in, dummyStdOut, dummySpyAlerter)
		cli.PlayPoker()

		pokertesting.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &player.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := cli2.NewCLI(playerStore, in, dummyStdOut, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlert{
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

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d wat not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]

				assertScheduledAlert(t, got, want)

			})
		}
	})

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		cli := cli2.NewCLI(dummyPlayerStore, dummyStdIn, stdout, dummyBlindAlerter)
		cli.PlayPoker()

		got := stdout.String()
		want := cli2.PlayerPrompt

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func assertScheduledAlert(t *testing.T, got scheduledAlert, want scheduledAlert) {
	if got.amount != want.amount {
		t.Fatalf("got amount %d, want %v", got.amount, want.amount)
	}

	if got.scheduledAt != want.scheduledAt {
		t.Errorf("got schedule time of %v, want %v", got.scheduledAt, want.scheduledAt)
	}
}
