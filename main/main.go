package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	GREY   = "\033[37m"
	RESET  = "\033[0m"
)

var (
	//refactor to get attempts from cli
	attempts = 5
)

// main function that fill parse input
func main() {
	fmt.Printf(
		"> Welcome to the wordle-cli game!\n"+
			"> You will have %d attempts to guess the randomWord. Good luck!\n",
		attempts)

	words := []string{"hello", "big", "me"}

	//read input steam
	reader := bufio.NewReader(os.Stdin)

	// select randomly a randomWord to guess
	rand.Seed(int64(time.Now().UnixNano()))
	randomWord := words[rand.Intn(len(words))]

	for attempts > 0 {
		fmt.Print("$ ")
		line, _ := reader.ReadString('\n')
		text := strings.TrimSpace(strings.TrimSuffix(line, "\n"))

		if text == randomWord {
			fmt.Println("> You guessed the randomWord!")
			fmt.Printf("%s%s\n", GREEN, text)
			fmt.Print(RESET)
			fmt.Print("> Do you want to play again? (y/n)\n")
			yes, _ := reader.ReadString('\n')

			if strings.ToLower(strings.TrimSuffix(yes, "\n")) != "y" {
				os.Exit(0)
			}

			attempts = 5
			randomWord = words[rand.Intn(len(words))]
			fmt.Println("Starting new round. Ready, set and go!")
			continue
		}

		var guess string
		var suffix string

		for index, letter := range text {
			if index > len(randomWord)-1 {
				break
			}

			guessLetter := string(letter)
			wordLetter := string(randomWord[index])

			if guessLetter == wordLetter {
				guess += fmt.Sprintf("%s%s", GREEN, guessLetter)
			} else if strings.Contains(randomWord, guessLetter) {
				guess += fmt.Sprintf("%s%s", YELLOW, guessLetter)
			} else {
				guess += fmt.Sprintf("%s%s", GREY, guessLetter)
			}

		}

		if len(randomWord) > len(text) {
			suffix = strings.Repeat("_", len(randomWord)-len(text))
		}

		fmt.Println(guess + suffix)
		fmt.Print(RESET)
		attempts -= 1
	}

	os.Exit(0)
}
