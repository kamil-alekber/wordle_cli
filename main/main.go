package main

import (
	"os"
	"wordle_cli/game"
)

func main() {
	game := game.CreateGame()
	game.Start()

	os.Exit(0)
}
