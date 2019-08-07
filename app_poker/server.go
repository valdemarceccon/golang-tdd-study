package poker

import (
	"encoding/json"
	"fmt"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/player"
	"net/http"
)

type PlayerServer struct {
	store player.PlayerStore
	http.Handler
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func NewPlayerServer(store player.PlayerStore) (p *PlayerServer) {
	p = new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return
}

func (p *PlayerServer) processWin(writer http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	writer.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(writer http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}

	_, _ = fmt.Fprint(writer, score)
}
