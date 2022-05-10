package game

import "fmt"

func startText(wc *WordleCli) {

	fmt.Printf(
		"> Welcome to the wordle-cli game!\n"+
			"> You will have %d attempts to guess the randomWord. Good luck!\n",
		wc.Attempts)

}
