package main

import (
	"github.com/valdemarceccon/golang-tdd-study/app"
	"log"
	"net/http"
)

func main() {
	server := &app.PlayerServer{PlayerStore: app.NewInMemoryPlayerStore()}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
