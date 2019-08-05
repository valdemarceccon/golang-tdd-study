package app

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	PlayerStore
}

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}

}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (p *PlayerServer) processWin(writer http.ResponseWriter, player string) {
	p.PlayerStore.RecordWin(player)
	writer.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(writer http.ResponseWriter, player string) {
	score := p.PlayerStore.GetPlayerScore(player)

	if score == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}

	_, _ = fmt.Fprint(writer, score)
}

func (s *StubPlayerStore) GetPlayerScore(name string) (score int) {
	score = s.Scores[name]
	return
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}
