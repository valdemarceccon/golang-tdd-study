package main

import (
	"fmt"
	poker "github.com/valdemarceccon/golang-tdd-study/app_poker"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/cmd/cli"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {name} wins to record a win")

	game := cli.NewTexasHoldem(cli.BlindAlerterFunc(cli.StdOutAlerter), store)
	cli := cli.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
