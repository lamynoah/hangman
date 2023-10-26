package hangman

import (
	"fmt"
	"math/rand"
	"strings"
)

func Revealletters(HMD *HangManData, l string) bool {
	found := false
	for i := 0; i < len(HMD.ToFind); i++ {
		if l == string(HMD.ToFind[i]) {
			HMD.Word[i] = rune(HMD.ToFind[i])
			found = true
		}
	}
	return found
}

func ChoosseLetters(HMD *HangManData) {
	var l string
	fmt.Println("Good luck you have 10 attemps.")
	for HMD.Attempts > 0 && countnbr(HMD.Word) > 0 {
		fmt.Printf("%c\n", HMD.Word)
		fmt.Print("Choose : ")
		fmt.Scan(&l)
		found := Revealletters(HMD, l)
		Foundsword(HMD, l)
		for _, i := range HMD.Letterstock {
			if i == l {
				fmt.Printf("You said this letter earlier")
				found = true
			}
		}
		if !found {
			HMD.Attempts--
			HMD.Letterstock = append(HMD.Letterstock, l)
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
}

func DisplayRandomLetter(HMD *HangManData) {
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
}
