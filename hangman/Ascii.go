package hangman

import (
	"bufio"
	"log"
	"os"
)

func ASCIIART(HMD *HangManData) []string {
	f, err := os.Open(HMD.AsciiFile)
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
}
func DisplayWordAscii(HMD *HangManData) {
	ligne_array := ASCIIART(HMD)
	for _, v := range ligne_array {
		println(v)
	}
}
