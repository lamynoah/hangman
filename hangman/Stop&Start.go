package hangman

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Save(HMD *HangManData) {
	EncryptedData, err := json.Marshal(HMD)
	f, err := os.Create("save.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.Write(EncryptedData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Game saved in save.txt")
	os.Exit(0)
}

func Start(HMD *HangManData, EncryptedData []byte) error {
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
