package cli

import (
	"github.com/valdemarceccon/golang-tdd-study/app_poker/player"
	"time"
)

type Game struct {
	alerter BlindAlerter
	store   player.PlayerStore
}

func (p *Game) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5*numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (p *Game) Finish(winner string) {
	p.store.RecordWin(winner)
}

func NewGame(alerter BlindAlerter, store player.PlayerStore) *Game {
	return &Game{
		alerter: alerter,
		store:   store,
	}
}
