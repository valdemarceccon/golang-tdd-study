package cmd

import (
	"fmt"
	poker "github.com/valdemarceccon/golang-tdd-study/app_poker"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/cmd/cli"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, closeDB, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer closeDB()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	cli.NewCLI(store, os.Stdin).PlayPoker()
}
