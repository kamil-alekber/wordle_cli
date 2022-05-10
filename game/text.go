package game

import "fmt"

func startText(wc *wordleCli) {

	fmt.Printf(
		"> Welcome to the wordle-cli game!\n"+
			"> You will have %d attempts to guess the random word. Good luck!\n",
		wc.attempts)

}
