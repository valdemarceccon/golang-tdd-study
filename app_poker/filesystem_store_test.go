package poker_test

import (
	poker "github.com/valdemarceccon/golang-tdd-study/app_poker"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/player"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/pokertesting"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := pokertesting.CreateTempFile(t, `[
			{"Name": "Cleo","Wins": 10},
			{"Name": "Chris","Wins": 33}]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)

		pokertesting.AssertNotError(t, err)

		want := []player.Player{
			{Name: "Chris", Wins: 33},
			{Name: "Cleo", Wins: 10},
		}

		got := store.GetLeague()
		pokertesting.AssertLeague(t, got, want)

		got = store.GetLeague()
		pokertesting.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := pokertesting.CreateTempFile(t, `[
			{"Name": "Cleo","Wins": 10},
			{"Name": "Chris","Wins": 33}]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)

		pokertesting.AssertNotError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33

		pokertesting.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := pokertesting.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)

		pokertesting.AssertNotError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		pokertesting.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := pokertesting.CreateTempFile(t, `[
		        {"Name": "Cleo", "Wins": 10},
		        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, _ := poker.NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		pokertesting.AssertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := pokertesting.CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := poker.NewFileSystemPlayerStore(database)

		pokertesting.AssertNotError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := pokertesting.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)

		pokertesting.AssertNotError(t, err)

		got := store.GetLeague()

		want := []player.Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		pokertesting.AssertLeague(t, got, want)

		got = store.GetLeague()
		pokertesting.AssertLeague(t, got, want)
	})
}
