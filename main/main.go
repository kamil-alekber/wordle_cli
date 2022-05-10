package main

import (
	"fmt"
	"os"
	"strconv"
	"wordle_cli/argumentParser"
	"wordle_cli/game"
)

func main() {
	args := argumentParser.CliArgumentParser()
	attempts, err := strconv.Atoi(args["--attempts"])

	if err != nil {
		fmt.Println("Attempts not set. Defaulting to 5.")
		attempts = 5
	}

	fmt.Printf("Num of attempts: %d\n", attempts)

	game := game.WordleCli{
		Attempts: attempts,
	}

	randomWord := game.GetRandomWord()

	for game.GetAttempts() > 0 {
		game.Start(randomWord)
	}

	os.Exit(0)
}
