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

func ChooseLetters(HMD *HangManData) {
	var l string
	fmt.Println("Good luck you have", HMD.Attempts, "attemps")
	for HMD.Attempts > 0 && countnbr(HMD.Word) > 0 {
		if HMD.IsAscii {
			DisplayWordAscii(HMD)
		} else {
			fmt.Printf("%c\n", HMD.Word)
		}
		fmt.Print("Choose : ")
		fmt.Scan(&l)
		if l == "STOP" {
			Save(HMD)
		}

		
		if strings.Count(l, "") > 2 {
			if l != HMD.ToFind {
				HMD.Attempts -= 2
				fmt.Println("Wrong word,", HMD.Attempts, "attempts remaining")
			} else {
				HMD.Word = []rune(HMD.ToFind)
			}
		} else {
			found := Revealletters(HMD, l)
			if !found && !UsedLetter(HMD,l) {
				HMD.Letterstock = append(HMD.Letterstock, l)
				HMD.Attempts -= 1
				fmt.Println("Not present in the word,", HMD.Attempts, "attempts remaining")
			} else if UsedLetter(HMD,l) {
				fmt.Printf("You said this letter earlier")
			}
		}

		
		if HMD.Attempts < 10 {
			fmt.Printf(HMD.HangmanPositions[9-HMD.Attempts])
		}
	}
}

func UsedLetter(HMD *HangManData, l string) bool {
	for _, i := range HMD.Letterstock {
		if i == l {
			return true
		}
	}
	return false
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
