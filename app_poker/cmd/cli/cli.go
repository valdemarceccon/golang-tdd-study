package cli

import (
	"bufio"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/player"
	"io"
	"strings"
)

type CLI struct {
	playerStore player.PlayerStore
	in          *bufio.Scanner
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func NewCLI(store player.PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}
