package hangman

type HangManData struct {
	Word             []rune     // Word composed of '_', ex: H_ll_
	ToFind           string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts         int        // Number of attempts left
	HangmanPositions [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	Letterstock      []string   // stock the letter used
	Foundword        bool
	AsciiFile 		 string 
}
