package game

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type WordleCli struct {
	//refactor to get attempts from cli
	Attempts int
}

var (
	words = []string{"hello", "big", "me"}
)

func (wc *WordleCli) Start(randomWord string) {

	fmt.Print("$ ")

	reader := bufio.NewReader(os.Stdin)
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

		wc.Attempts = 5
		fmt.Println("Starting new round. Ready, set and go!")
		return
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
	wc.Attempts -= 1

}

func (wc *WordleCli) GetRandomWord() string {
	res, err := http.Get("https://random-word-api.herokuapp.com/word")

	if err != nil {
		log.Printf("Error: %s", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	var wordList []string

	json.NewDecoder(res.Body).Decode(&wordList)

	if len(wordList) < 1 {
		log.Printf("Error. Can not get random word %s", err)
		os.Exit(1)
	}

	return wordList[0]
}

func (wc *WordleCli) GetAttempts() int {
	return wc.Attempts
}
