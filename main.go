package main

import (
	"feat/hangman"
	"log"
	"os"
)

func main() {
	HMD := hangman.HangManData{
		Attempts:         10,
		HangmanPositions: hangman.FileJose(),
		Letterstock:      []string{},
		Foundword:        false,
		IsAscii:          false,
	}
	//repository
	var filepath string
	if len(os.Args) == 4 && os.Args[2] == "--letterFile" {
		HMD.IsAscii = true
		HMD.AsciiFile = os.Args[3]
	}
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
	hangman.ChooseLetters(&HMD)
	// reveal letters or word
	// display loose
	hangman.DisplayLosseWin(&HMD)
}
