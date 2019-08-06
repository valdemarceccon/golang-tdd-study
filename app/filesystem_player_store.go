package app

import (
	"encoding/json"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	Database io.Writer
	league   League
}

func NewFileSystemPlayerStore(database *os.File) *FileSystemPlayerStore {
	_, _ = database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		Database: &Tape{File: database},
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{Name: name, Wins: 1})
	}

	_ = json.NewEncoder(f.Database).Encode(f.league)
}
