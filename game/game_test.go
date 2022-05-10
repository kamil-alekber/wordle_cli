package game

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestRandomWordHttpRequest(t *testing.T) {
	res, err := http.Get("https://random-word-api.herokuapp.com/word")

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	defer res.Body.Close()

	var wordList []string

	json.NewDecoder(res.Body).Decode(&wordList)

	if len(wordList) < 1 {
		t.Errorf("Error. Can not get random word %s", err)
	}

}
