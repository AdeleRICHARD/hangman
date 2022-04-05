package main

import (
	"fmt"
	"os"

	"udemy.go/hangman/dictionary"
	"udemy.go/hangman/hangman"
)

func main() {
	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Couldn't load dictionary: %vn\n", err)
		os.Exit(1)
	}

	g, err := hangman.New(8, dictionary.PickWord())

	hangman.DrawWelcome()

	guess := ""
	for {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal : %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}

}
