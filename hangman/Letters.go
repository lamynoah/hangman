package hangman

import (
	"math/rand"
	"strings"
)

//return true if l = Tofind or l in the word 
func Revealletters(HMD *HangManData, l string) bool {
	found := false
	if l == HMD.ToFind {
		HMD.Word = []rune(l)
		return true
	}
	for i := 0; i < len(HMD.ToFind); i++ {
		if l == string(HMD.ToFind[i]) {
			HMD.Word[i] = rune(HMD.ToFind[i])
			found = true
		}
	}
	return found
}

func AppendLetterStock(HMD *HangManData, l string) {
	for HMD.Attempts > 0 && Countnbr(HMD.Word) > 0 {
		found := Revealletters(HMD, l)
		if !found && !UsedLetter(HMD, l) {
			HMD.Letterstock = append(HMD.Letterstock, l)
		}
	}
}

func ChooseLetters(HMD *HangManData, l string) {
	for HMD.Attempts > 0 && Countnbr(HMD.Word) > 0 {
		if strings.Count(l, "") > 2 {
			if l != HMD.ToFind {
				HMD.Attempts -= 2
			} else {
				HMD.Word = []rune(HMD.ToFind)
			}
		} else {
			found := Revealletters(HMD, l)
			if !found && !UsedLetter(HMD, l) {
				HMD.Letterstock = append(HMD.Letterstock, l)
				HMD.Attempts -= 1
			}
		}
	}
}

// return true if l in lettestock
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
