package wordchains

import (
	"os"
	"fmt"
	"log"
	"errors"
)


const (
	filePath = "wordlist.txt"
)

var (
	length int
	dictionary []string
)

func Start(word1 string, word2 string) {


	length, err := CheckWords(word1, word2)
	if err != nil {
		log.Fatal(err.Error())
	}
	//readFile()
	fmt.Println("everything ok so far ")
	fmt.Println(length)
}


// First we check the length of the words so they aren't invalid
// Retuns the len of the first word(which should be the same as the second one) and if has been an error
func CheckWords(word1 string, word2 string) (int, error) {

	len1 := len(word1)
	len2 := len(word2)

	if (len1 != len2) || (len1 == 0) {
		return 0, errors.New("invalid words: " + word1 + ", " + word2)
	}
	return len1, nil
}

// We load all data from the .txt file
func readFile() {
	fmt.Println("Reading word list...")

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
    	log.Fatal("Error opening word list archive: " + err. Error())
	}



	defer file.Close()
}