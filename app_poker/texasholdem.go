package poker

import (
	"github.com/valdemarceccon/golang-tdd-study/app_poker/player"
	"io"
	"os"
	"time"
)

type TexasHoldem struct {
	alerter BlindAlerter
	store   player.PlayerStore
}

func NewTexasHoldem(alerter BlindAlerter, store player.PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}

func (p *TexasHoldem) Start(numberOfPlayers int, to io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind, os.Stdout)
		blindTime = blindTime + blindIncrement
	}
}

func (p *TexasHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}
