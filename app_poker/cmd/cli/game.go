package cli

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}
