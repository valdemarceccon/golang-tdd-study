package cli_test

import (
	cli2 "github.com/valdemarceccon/golang-tdd-study/app_poker/cmd/cli"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/player"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/pokertesting"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &player.StubPlayerStore{}

		cli := cli2.NewCLI(playerStore, in)
		cli.PlayPoker()

		pokertesting.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &player.StubPlayerStore{}

		cli := cli2.NewCLI(playerStore, in)
		cli.PlayPoker()

		pokertesting.AssertPlayerWin(t, playerStore, "Cleo")
	})
}

//func assertPlayerWin(t *testing.T, store *player.StubPlayerStore, winner string) {
//	t.Helper()
//
//	if len(store.WinCalls) != 1 {
//		t.Fatalf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
//	}
//
//	if store.WinCalls[0] != winner {
//		t.Errorf("did not store correct winner got %q want %q", store.WinCalls[0], winner)
//	}
//}
