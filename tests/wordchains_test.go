package tests 

import (
	"testing"

	"cycloid/wordchains"

)

func TestCheckWords (t *testing.T) {

	word1 := "hello"
	word2 := "okbye"

	result, err := wordchains.CheckWords(word1, word2)
	if result != 5 {
		t.Errorf("Length of word should be 5")
	}
	if err != nil {
		t.Errorf("Words should be with same length: " + err.Error())
	}
}

func TestIsStoredWord (t *testing.T) {
	
	list := map[string]bool {"hello": true, "okbye": true}

	if !wordchains.IsStoredWord(list, "hello") {
		t.Errorf("Word not found in the map")
	}
}

func TestIsStoredLetter (t *testing.T) {

	list := map[rune]bool {'a': true, 'b': true, 'c': true}

	if !wordchains.IsStoredLetter(list, 'b') {
		t.Errorf("Word not found in the map")
	}
}

func TestStartVariables (t *testing.T) {

	wordchains.StartVariables()

	wordchains.ReadFile("../wordchains/wordlist.txt", 1)

	dictionary := wordchains.Dictionary
	if len(dictionary) == 0 {
		t.Errorf("Dictionary not prepared")	
	}

	letters    := wordchains.Letters
	if len(letters) == 0 {
		t.Errorf("List of letters not prepared")	
	}
}

func TestFindPath (t *testing.T) {
	
	wordchains.StartVariables()

	wordchains.ReadFile("../wordchains/wordlist.txt", 4)


	slicePath := wordchains.FindPath("lead", "gold")
	testPath := [4]string{}
	copy(testPath[:], slicePath)

	truePath := [4]string{"lead", "load", "goad", "gold"}

	if testPath != truePath {
		t.Errorf("Error in the creation of similar words")
	}
}
