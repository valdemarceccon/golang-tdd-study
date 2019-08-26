package poker_test

import (
	"github.com/gorilla/websocket"
	"github.com/valdemarceccon/golang-tdd-study/app_poker"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/player"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/pokertesting"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const jsonContentType = "application/json"

var (
	dummyGame = &GameSpy{}
)

func TestGETPlayers(t *testing.T) {
	store := &player.StubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := mustMakePlayerServer(t, store, dummyGame)
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
	server := mustMakePlayerServer(t, store, dummyGame)

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

	server := mustMakePlayerServer(t, store, dummyGame)
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
		server := mustMakePlayerServer(t, store, dummyGame)

		request := pokertesting.NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := pokertesting.GetLeagueFromResponse(t, response.Body)
		pokertesting.AssertStatus(t, response.Code, http.StatusOK)
		pokertesting.AssertLeague(t, got, wantedLeague)
		pokertesting.AssertContentType(t, response, jsonContentType)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /playGame return 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, &player.StubPlayerStore{}, dummyGame)

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
	})

	t.Run("start a playGame with 3 players and declare Ruth the winner", func(t *testing.T) {
		game := &GameSpy{}
		winner := "Ruth"
		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		ws, cleanUp := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer server.Close()
		defer cleanUp()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, winner)
	})
}

func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}

func writeWSMessage(t *testing.T, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}

func mustMakePlayerServer(t *testing.T, store player.PlayerStore, game poker.Game) *poker.PlayerServer {
	t.Helper()
	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) (*websocket.Conn, func()) {
	t.Helper()
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws, func() {
		err := ws.Close()
		if err != nil {
			t.Fatal("error closing connection")
		}
	}
}
