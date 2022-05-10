package argumentParser

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestCliArgumentParser(t *testing.T) {
	os.Args = []string{"", "--attempts", "10"}

	args := CliArgumentParser()

	attempts, err := strconv.Atoi(args["--attempts"])
	fmt.Printf("attempts %d", attempts)

	if err != nil {
		t.Errorf("Error casting attempts to int: %s", err)
		return
	}
	fmt.Println("arguments", attempts)
	fmt.Println(args)

	if attempts != 10 {
		t.Errorf("Error. Should set number of attempts to 10.")
	}

}
