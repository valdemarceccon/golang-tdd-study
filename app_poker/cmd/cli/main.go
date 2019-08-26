package cli

import (
	"fmt"
	poker "github.com/valdemarceccon/golang-tdd-study/app_poker"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, cleanUp, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer cleanUp()

	fmt.Println("Let's play poker")
	fmt.Println("Type {name} wins to record a win")

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
