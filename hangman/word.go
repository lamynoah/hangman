package hangman

import (
	"math/rand"
	"strings"
)



// choisis le modts aléatoire 
func RandomWord(content []byte, HMD *HangManData) {
	words := strings.Fields(string(content))
	randomIndex := rand.Intn(len(words))
	randomWord := words[randomIndex]
	HMD.ToFind = randomWord
	for i := 0; i < len(HMD.ToFind); i++ {
		HMD.Word = append(HMD.Word, '_')
	}
}
