package main

import (
	"feat/feat"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type HangManData struct {
	Word             []rune     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	letterstock      []string   // stock the letter used
}

func main() {
	var HMD HangManData
	HMD.Attempts = 10
	HMD.HangmanPositions = feat.FileJose()
	HMD.letterstock = []string{}
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// words found  random
	words := strings.Fields(string(content))
	randomIndex := rand.Intn(len(words))
	randomWord := words[randomIndex]
	HMD.ToFind = randomWord
	for i := 0; i < len(HMD.ToFind); i++ {
		HMD.Word = append(HMD.Word, '_')
	}
	// letter display random
	n := len([]rune(HMD.ToFind))/2 - 1
	for letter := 0; letter < n; letter++ {
		randompos := rand.Intn(len(HMD.ToFind))
		if HMD.Word[randompos] == '_' {
			HMD.Word[randompos] = rune(HMD.ToFind[randompos])
			found := Revealletters(HMD, string(HMD.ToFind[randompos]))
			if found {
				found = false
				letter++
			}
		} else {
			letter--
		}
	}
	// display
	var l string
	foundword := false
	fmt.Println("Good luck you have 10 attemps.")
	for HMD.Attempts > 0 && countnbr(HMD.Word) > 0 {
		fmt.Printf("%c\n", HMD.Word)
		fmt.Print("Choose : ")
		fmt.Scan(&l)
		for _, i := range HMD.letterstock {
			if i == l {
				fmt.Printf("You said his letter earlier \n")
			}
		}
		HMD.letterstock = append(HMD.letterstock, l)
		found := Revealletters(HMD, l)
		if l == string(HMD.ToFind) {
			HMD.Word = []rune(HMD.ToFind)
			foundword = true
		}
		if !found && !foundword {
			HMD.Attempts--
			if strings.Count(l, "") > 2 {
				HMD.Attempts--
			}
			fmt.Println("Not present in the word,", HMD.Attempts, "attempts remaining")
		} else {
			found = false
		}
		if HMD.Attempts < 10 {
			fmt.Printf(HMD.HangmanPositions[9-HMD.Attempts])
		}
	}
	// display loose
	if HMD.Attempts == 0 && countnbr(HMD.Word) > 0 {
		fmt.Println("---------------------------------------")
		fmt.Println(" 		VOUS AVEZ PERDU 	")
		fmt.Println("---------------------------------------")
	}
	// display victory
	if HMD.Attempts > 0 && countnbr(HMD.Word) == 0 {
		fmt.Printf("%c\n", HMD.Word)
		fmt.Println("Congrats !")
	}
}

func Revealletters(HMD HangManData, l string) bool {
	found := false
	for i := 0; i < len(HMD.ToFind); i++ {
		if l == string(HMD.ToFind[i]) {
			HMD.Word[i] = rune(HMD.ToFind[i])
			found = true
		}
	}
	return found
}

func countnbr(word []rune) int {
	c := 0
	for i := 0; i < len(word); i++ {
		if word[i] == '_' {
			c++
		}
	}
	return c
}
