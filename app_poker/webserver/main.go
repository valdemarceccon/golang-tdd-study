package main

import (
	"github.com/valdemarceccon/golang-tdd-study/app_poker"
	"log"
	"net/http"
	"os"
)

const dbFilename = "game.db.json"

func main() {

	db, err := os.OpenFile(dbFilename, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFilename, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem create file system player store %v", err)
	}

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
