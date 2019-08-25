package cli_test

import (
	"fmt"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/cmd/cli"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/player"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/pokertesting"
	"testing"
	"time"
)

func TestGame_Start(t *testing.T) {
	var dummyPlayerStore = &player.StubPlayerStore{}

	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := cli.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(5)

		cases := []scheduledAlert{
			{scheduledAt: 0 * time.Second, amount: 100},
			{scheduledAt: 10 * time.Minute, amount: 200},
			{scheduledAt: 20 * time.Minute, amount: 300},
			{scheduledAt: 30 * time.Minute, amount: 400},
			{scheduledAt: 40 * time.Minute, amount: 500},
			{scheduledAt: 50 * time.Minute, amount: 600},
			{scheduledAt: 60 * time.Minute, amount: 800},
			{scheduledAt: 70 * time.Minute, amount: 1000},
			{scheduledAt: 80 * time.Minute, amount: 2000},
			{scheduledAt: 90 * time.Minute, amount: 4000},
			{scheduledAt: 100 * time.Minute, amount: 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}
		game := cli.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7)

		cases := []scheduledAlert{
			{scheduledAt: 0 * time.Second, amount: 100},
			{scheduledAt: 12 * time.Minute, amount: 200},
			{scheduledAt: 24 * time.Minute, amount: 300},
			{scheduledAt: 36 * time.Minute, amount: 400},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &player.StubPlayerStore{}
	blindAlerter := &SpyBlindAlerter{}
	game := cli.NewTexasHoldem(blindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	pokertesting.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(cases []scheduledAlert, t *testing.T, blindAlerter *SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {

			if len(blindAlerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
			}

			got := blindAlerter.alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}
