package main

import (
	"github.com/valdemarceccon/golang-tdd-study/app_poker"
	"log"
	"net/http"
)

const dbFilename = "game.db.json"

func main() {

	store, closeDB, err := poker.FileSystemPlayerStoreFromFile(dbFilename)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFilename, err)
	}

	defer closeDB()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
