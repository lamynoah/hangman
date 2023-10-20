package main

import (
	"fmt"
	"io/ioutil"
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
}

func main() {
	var HMD HangManData
	HMD.Attempts = 10
	content, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 2 {
		log.Fatal(err)
	}

	words := strings.Fields(string(content))
	randomIndex := rand.Intn(len(words))
	randomWord := words[randomIndex]
	HMD.ToFind = randomWord
	for i := 0; i < len(HMD.ToFind); i++ {
		HMD.Word = append(HMD.Word, '_')
	}
	n := len([]rune(HMD.ToFind))/2 - 1
	for letter := 0; letter < n; letter++ {
		randompos := rand.Intn(len(HMD.ToFind))
		if HMD.Word[randompos] == '_' {
			HMD.Word[randompos] = rune(HMD.ToFind[randompos])
		} else {
			letter--
		}
	}
	var l string
	found := false
	fmt.Println("Good luck you have 10 attemps.")
	for HMD.Attempts > 0 && countnbr(HMD.Word) > 0 {
		fmt.Printf("%c\n", HMD.Word)
		fmt.Print("Choose : ")
		fmt.Scan(&l)
		for i := 0; i < len(HMD.ToFind); i++ {
			if l == string(HMD.ToFind[i]) {
				HMD.Word[i] = rune(HMD.ToFind[i])
				found = true
			}
		}
		if !found {
			HMD.Attempts--
			fmt.Println("Not present in the word,", HMD.Attempts, "attempts remaining")
		} else {
			found = false
		}
	}
	if HMD.Attempts == 0 && countnbr(HMD.Word) > 0 {
		fmt.Println("---------------------------------------")
		fmt.Println(" 		VOUS AVEZ PERDU 	")
		fmt.Println("---------------------------------------")
	}
	if HMD.Attempts > 0 && countnbr(HMD.Word) == 0 {
		fmt.Println("Congrats !")
	}
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
