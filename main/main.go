package main

import (
	"os"
	"wordle_cli/game"
)

func main() {
	game.CreateGame().Start()

	os.Exit(0)
}
