package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("What is your letter? ")
		guess, err = reader.ReadString('\n') // The end of input
		if err != nil {
			return guess, err
		}

		guess = strings.TrimSpace(guess)
		if len(guess) != 1 {
			fmt.Printf("Invalid letter input. letter=%v, len=%v\n", guess, len(guess))
			continue
		}
		valid = true
	}
	return strings.ToUpper(guess), nil
}
