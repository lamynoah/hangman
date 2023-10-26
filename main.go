package main

import (
	"feat/hangman"
	"log"
	"os"
)

func main() {
	var HMD hangman.HangManData
	HMD.Attempts = 10
	HMD.HangmanPositions = hangman.FileJose()
	HMD.Letterstock = []string{}
	HMD.Foundword = false
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// words found  random
	hangman.RandomWord(content, &HMD)
	// letter display random
	hangman.DisplayRandomLetter(&HMD)
	// display
	hangman.ChoosseLetters(&HMD)
	// reveal letters or word
	// display loose
	hangman.DisplayLosseWin(&HMD)
}
