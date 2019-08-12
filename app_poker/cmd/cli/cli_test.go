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
	var dummyPlayerStore = &player.StubPlayerStore{}
	var dummyStdOut = &bytes.Buffer{}

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("3\nChris wins\n")
		playerStore := &player.StubPlayerStore{}

		game := cli2.NewGame(dummySpyAlerter, playerStore)

		cli := cli2.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		pokertesting.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("3\nCleo wins\n")
		playerStore := &player.StubPlayerStore{}

		game := cli2.NewGame(dummySpyAlerter, playerStore)

		cli := cli2.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		pokertesting.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		numOfPlayers := 4
		in := strings.NewReader(fmt.Sprintf("%d\nChris wins\n", numOfPlayers))
		playerStore := &player.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}
		game := cli2.NewGame(blindAlerter, playerStore)

		cli := cli2.NewCLI(in, dummyStdOut, game)

		cli.PlayPoker()

		cases := []scheduledAlert{
			{time.Duration(5*numOfPlayers*0) * time.Second, 100},
			{time.Duration(5*numOfPlayers*1) * time.Minute, 200},
			{time.Duration(5*numOfPlayers*2) * time.Minute, 300},
			{time.Duration(5*numOfPlayers*3) * time.Minute, 400},
			{time.Duration(5*numOfPlayers*4) * time.Minute, 500},
			{time.Duration(5*numOfPlayers*5) * time.Minute, 600},
			{time.Duration(5*numOfPlayers*6) * time.Minute, 800},
			{time.Duration(5*numOfPlayers*7) * time.Minute, 1000},
			{time.Duration(5*numOfPlayers*8) * time.Minute, 2000},
			{time.Duration(5*numOfPlayers*9) * time.Minute, 4000},
			{time.Duration(5*numOfPlayers*10) * time.Minute, 8000},
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
		numberOfPlayers := 3
		stdout := &bytes.Buffer{}
		in := strings.NewReader(fmt.Sprintf("%d\nCarlos wins\n", numberOfPlayers))
		blindAlerter := &SpyBlindAlerter{}
		game := cli2.NewGame(blindAlerter, dummyPlayerStore)

		cli := cli2.NewCLI(in, stdout, game)

		cli.PlayPoker()

		got := stdout.String()
		want := cli2.PlayerPrompt

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		cases := []scheduledAlert{
			{time.Duration(0*5*numberOfPlayers) * time.Second, 100},
			{time.Duration(1*5*numberOfPlayers) * time.Minute, 200},
			{time.Duration(2*5*numberOfPlayers) * time.Minute, 300},
			{time.Duration(3*5*numberOfPlayers) * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alerts %d was not scheduled for %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)
			})
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
