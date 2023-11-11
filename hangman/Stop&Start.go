package hangman

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// var saveFileName string

// // func init() {
// // 	flag.StringVar(&saveFileName, "--startwith", "", ".txt")
// // 	flag.Parse()
// // }

func Save(HMD *HangManData) error {
	encryptedData, err := json.Marshal(HMD)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("save.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	defer f.Close()
	_, err = f.Write(encryptedData)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Game saved in save.txt\n")
	os.Exit(0)
	return nil
}

func Start(HMD *HangManData) error {

	data, err := os.ReadFile("save.txt")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, HMD)
	if err != nil {
		return err
	}
	return nil
}
