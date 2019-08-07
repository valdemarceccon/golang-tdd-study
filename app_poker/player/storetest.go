package player

type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

func (s *StubPlayerStore) GetPlayerScore(name string) (score int) {
	score = s.Scores[name]
	return
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}
