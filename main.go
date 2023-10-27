package main

import (
	"feat/hangman"
	"log"
	"os"
)

func main() {
	var HMD hangman.HangManData
	HMD.Attempts = 10
	// HMD.Word = []rune(os.Args[1])
	// hangman.DisplayWordAscii(&HMD)
	// return
	HMD.HangmanPositions = hangman.FileJose()
	HMD.Letterstock = []string{}
	HMD.Foundword = false
	isAscii := false
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 2  && os.Args[2] == "--letterFile" {
		isAscii = true
	} 
	// words found  random
	hangman.RandomWord(content, &HMD)
	// letter display random
	hangman.DisplayRandomLetter(&HMD)
	// display
	hangman.ChoosseLetters(&HMD, isAscii)
	// reveal letters or word
	// display loose
	hangman.DisplayLosseWin(&HMD, isAscii)

}
