package wordchains

import (
	"os"
	"fmt"
	"log"
	"bufio"
	"errors"
	"strconv"
)


const (

)

var (

	// All real worlds in the .txt with an specific number of letters
	Dictionary map[string]bool
	//All letters present in the 'dictionary' list
	Letters map[rune]bool
)

func Start(word1 string, word2 string) {

	length, err := CheckWords(word1, word2)
	if err != nil {
		log.Fatal(err.Error())
	}

	StartVariables()

	ReadFile("wordchains/wordlist.txt", length)

	wordchain := FindPath(word1, word2)
	fmt.Println("**********")
	fmt.Println("Found shortest path: ")
	fmt.Println(wordchain)
	fmt.Println("**********")
}

func StartVariables () () {
	Dictionary 	= map[string]bool{}
	Letters 	= map[rune]bool{}
}


// First we check the length of the words so they aren't invalid
// Retuns the len of the first word(which should be the same as the second one) and if an error occurred
func CheckWords(word1 string, word2 string) (int, error) {

	len1 := len(word1)
	len2 := len(word2)

	if (len1 != len2) || (len1 == 0) {
		return 0, errors.New("invalid words: " + word1 + ", " + word2)
	}
	return len1, nil
}

// We read all lines from the .txt file and store those with the length we need
func ReadFile(filePath string, length int) {
	fmt.Println("Reading word list...")

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
    	log.Fatal("Error opening word list archive: " + err. Error())
	}
	defer file.Close()


    scanner := bufio.NewScanner(file)

    // We read the file until EOF
	for scanner.Scan() {
		eachWord := scanner.Text()
	//	fmt.Println(eachWord + " " + strconv.Itoa(len(eachWord)) + " against " + strconv.Itoa(length))
		// If a word is with the same length as the ones from the chain
		if (len(eachWord) == length) {

			addWord(Dictionary, eachWord)

			for _, eachLetter := range eachWord {

				if !IsStoredLetter(Letters, eachLetter) {
					addLetter(Letters, eachLetter)
				}
			} 
		}
	}
}


// Returns a list with all real words that differ from a specific word in just a letter
func SimilarWords (word string) ([]string) {

	copyOfWord := []rune(word)

	result := []string{}

	for i, _ := range copyOfWord {
		for eachRune, _ := range Letters {

			copyOfWord[i] = eachRune
			changedWord := string(copyOfWord)

			// If the new word is a real word (is in the dictionary) we store them
			// We also check if the word is the same just in case
			if (IsStoredWord(Dictionary, changedWord) && changedWord != word) {
				//addWord(result, changedWord)
				result = append(result, changedWord)
			}

			// We restore the word so it only changes a letter each time
			copyOfWord = []rune(word)
		}
	}
	//fmt.Print("Similar words to " + word + ":   ")
	//fmt.Println(result)

	return result
}


func addWord (list map[string]bool, word string) () {
	list[word] = true
}

func addLetter (list map[rune]bool, letter rune) () {
	list[letter] = true
}

func IsStoredWord (list map[string]bool, word string) (bool) {
	return list[word]
}

func IsStoredLetter (list map[rune]bool, letter rune) (bool) {
	return list[letter]
}

func FindPath (word1 string, word2 string) ([]string) {

	fmt.Println("Started creating wordchain " + word1 + " => " + word2)

	paths := [][]string{}
	paths = append(paths, []string{word1})

	nextRound := [][]string{}
	
	shortestPath := []string{}
	numRounds := 1

	loop:
	for {
		fmt.Println("-----Starting round " + strconv.Itoa(numRounds) + " with " + strconv.Itoa(len(paths)) + " paths...")
		//printPaths(paths)
		
		for _, eachPath := range paths {
			
			//fmt.Println(eachPath)
			lastWord := lastWordInPath(eachPath)
			nextWords := SimilarWords(lastWord)

			for _, eachNextWord := range nextWords {
				if eachNextWord != word2 {
					if !pathContainsWord(eachPath, eachNextWord) {
						newPath := addWordToPath(eachPath, eachNextWord)
						nextRound = append(nextRound, newPath)
					}
				} else {
				//	fmt.Println("/////////////////////// Found " + eachNextWord + "! ///////////////////////")
					shortestPath = addWordToPath(eachPath, eachNextWord)
					break loop
				}
			}

			
		}
		paths = nextRound
		nextRound = [][]string{}
		numRounds++
		//fmt.Println("Size of paths: " + strconv.Itoa(len(paths)))
		//fmt.Println("Size of nextRound: " + strconv.Itoa(len(nextRound)))
	}
	
	return shortestPath
}

func addPath (allPaths [][]string, newPath []string) ([][]string) {

	allPaths = append(allPaths, newPath)
	return allPaths
}


func printPaths (allPaths [][]string) () {
	fmt.Println("List of paths: ")
	for _, eachPath := range allPaths {
		i := 0
		lenPath := len(eachPath)
		for _, eachWord := range eachPath {
			fmt.Print(eachWord)
			if (i != lenPath-1) {
				fmt.Print(" -> ")
			} else {
				fmt.Println()
			}
			i++
		}
	}
}

func pathContainsWord (path []string, word string) (bool) {
	for i:=0; i<len(path); i++ {
		if path[i] == word{
			return true
		}
	}
	return false
}

func lastWordInPath (path []string) (string) {
	return path[len(path) - 1]
}

func addWordToPath (path []string, newWord string) ([]string) {
	return append(path, newWord)
}