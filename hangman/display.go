package hangman

import "fmt"

func DisplayLosseWin(HMD *HangManData) {
	if HMD.Attempts == 0 && countnbr(HMD.Word) > 0 {
		fmt.Println("---------------------------------------")
		fmt.Println(" 		VOUS AVEZ PERDU 	")
		fmt.Println("---------------------------------------")
	}
	// display victory
	if HMD.Attempts > 0 && countnbr(HMD.Word) == 0 {
		if HMD.IsAscii {
			DisplayWordAscii(HMD)
		} else {
			fmt.Printf("%c\n", HMD.Word)
		}
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
