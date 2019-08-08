package cli_test

import (
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

	t.Run("it schedules priting of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		stubPlayerStore := &player.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := cli2.NewCLI(stubPlayerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) != 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})
}
