package hangman

import (
	"bufio"
	"log"
	"os"
)

func ASCIIART(HMD *HangManData) []string {
	// p := []string{}
	f, err := os.Open(os.Args[3])

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	content := []string{}
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	lignes := []string{}
	for nbligne := 0; nbligne < 10; nbligne++ {
		ligne := ""
		for _, v := range HMD.Word {
			ligne += content[int(v-32)*9+nbligne]
		}
		lignes = append(lignes, ligne)
	}
	return lignes
	// debut, fin := 0, 9
	// for i := 0; i < 100; i++ {
	// 	letter := ""
	// 	for _, v := range content[debut:fin] {
	// 		letter += v + "\n"
	// 	}
	// 	p = append(p, letter)
	// 	debut += 9
	// 	fin += 9
	// }
	// return p
}
func DisplayWordAscii(HMD *HangManData) {
	ligne_array := ASCIIART(HMD)
	for _, v := range ligne_array {
		println(v)
	}
}
