package hangman

import (
	"log"
	"os"
)

func FileJose() [10]string {
	pos := [10]string{}
	content, err := os.ReadFile("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	debut, fin := 0, 71
	for i := 0; i < 10; i++ {
		pos[i] = string(content[debut:fin])
		debut += 71
		fin += 71
	}
	return pos
}
