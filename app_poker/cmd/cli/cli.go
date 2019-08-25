package cli

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/valdemarceccon/golang-tdd-study/app_poker/player"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	playerStore player.PlayerStore
	in          *bufio.Scanner
	out         io.Writer
	game        Game
}

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "bad value received for number of players, please try again with a number"
const BadWinnerInputMsg = "bad value received for name of winner, please try again with a {Name} wins"

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner, err := extractWinner(winnerInput)

	if err != nil {
		fmt.Fprint(cli.out, err.Error())
		return
	}
	cli.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", errors.New(BadWinnerInputMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}
