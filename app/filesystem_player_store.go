package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	Database io.Writer
	league   League
}

func NewFileSystemPlayerStore(database *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(database)

	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewLeague(database)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("problem loading player store from file %s, %v", database.Name(), err))
	}

	return &FileSystemPlayerStore{
		Database: &Tape{File: database},
		league:   league,
	}, nil
}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file into from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
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
