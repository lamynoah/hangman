package main

import (
	"log"
	"os"

	"github.com/lamynoah/hangman/hangman"
)

func main() {
	HMD := hangman.HangManData{
		Attempts:         10,
		Letterstock:      []string{},
		Foundword:        false,
	}
	//repository
	var filepath string
	if len(os.Args) == 3 && os.Args[1] == "--startwith" {
		hangman.Start(&HMD)
		filepath = os.Args[2]
	} else {
		filepath = os.Args[1]
	}
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	if os.Args[1] != "--startwith" {
		// words found  random
		hangman.RandomWord(content, &HMD)
		// letter display random
		hangman.DisplayRandomLetter(&HMD)
	}
	// display

	// reveal letters or word
	// display loose
	hangman.DisplayLosseWin(&HMD)
}
