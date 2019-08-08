package poker_test

import (
	"github.com/valdemarceccon/golang-tdd-study/app_poker"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/player"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/pokertesting"
	"net/http"
	"net/http/httptest"
	"testing"
)

const jsonContentType = "application/json"

func TestGETPlayers(t *testing.T) {
	store := &player.StubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := poker.NewPlayerServer(store)
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		pokertesting.AssertStatus(t, response.Code, http.StatusOK)
		pokertesting.AssertResponseBody(t, response.Body.String(), "20")
	})
	t.Run("returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		pokertesting.AssertStatus(t, response.Code, http.StatusOK)
		pokertesting.AssertResponseBody(t, response.Body.String(), "10")
	})
	t.Run("return 404 on missing players", func(t *testing.T) {
		request := pokertesting.NewGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d want %d", got, want)
		}
	})
}

func TestStoreWins(t *testing.T) {
	store := &player.StubPlayerStore{Scores: map[string]int{}, WinCalls: nil}
	server := poker.NewPlayerServer(store)

	t.Run("it records wins when POST", func(t *testing.T) {
		somePlayer := "Pepper"
		request := pokertesting.NewPostWinRequest(somePlayer)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		pokertesting.AssertStatus(t, response.Code, http.StatusAccepted)

		pokertesting.AssertPlayerWin(t, store, somePlayer)
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := pokertesting.CreateTempFile(t, "[]")
	defer cleanDatabase()
	store, err := poker.NewFileSystemPlayerStore(database)

	pokertesting.AssertNotError(t, err)

	server := poker.NewPlayerServer(store)
	somePlayer := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), pokertesting.NewPostWinRequest(somePlayer))
	server.ServeHTTP(httptest.NewRecorder(), pokertesting.NewPostWinRequest(somePlayer))
	server.ServeHTTP(httptest.NewRecorder(), pokertesting.NewPostWinRequest(somePlayer))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, pokertesting.NewGetScoreRequest(somePlayer))
		pokertesting.AssertStatus(t, response.Code, http.StatusOK)

		pokertesting.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, pokertesting.NewLeagueRequest())
		pokertesting.AssertStatus(t, response.Code, http.StatusOK)

		got := pokertesting.GetLeagueFromResponse(t, response.Body)

		want := []player.Player{
			{Name: "Pepper", Wins: 3},
		}

		pokertesting.AssertLeague(t, got, want)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []player.Player{
			{Name: "Cleo", Wins: 32},
			{Name: "Chris", Wins: 20},
			{Name: "Tiest", Wins: 14},
		}

		store := &player.StubPlayerStore{League: wantedLeague}
		server := poker.NewPlayerServer(store)

		request := pokertesting.NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := pokertesting.GetLeagueFromResponse(t, response.Body)
		pokertesting.AssertStatus(t, response.Code, http.StatusOK)
		pokertesting.AssertLeague(t, got, wantedLeague)
		pokertesting.AssertContentType(t, response, jsonContentType)
	})
}
