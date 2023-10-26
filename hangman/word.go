package hangman

import (
	"math/rand"
	"strings"
)

func RandomWord(content []byte, HMD *HangManData) {
	words := strings.Fields(string(content))
	randomIndex := rand.Intn(len(words))
	randomWord := words[randomIndex]
	HMD.ToFind = randomWord
	for i := 0; i < len(HMD.ToFind); i++ {
		HMD.Word = append(HMD.Word, '_')
	}
}

func Foundsword(HMD *HangManData, l string) bool {
	if l == HMD.ToFind {
		HMD.Word = []rune(HMD.ToFind)
		return true
	}
	return false
}
