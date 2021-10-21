# cycloid
Cycloid technical tests for backend job offer

(in project root folder)> go run main.go

-----------------------------------------------

* Karate Chop (http://codekata.com/kata/kata02-karate-chop/)



-----------------------------------------------


* Word Chains (http://codekata.com/kata/kata19-word-chains/)

Given a dictionary/list of words (from the same page) , map from a word to another, only changing one letter each time, and only if the changed intermediate word is also in the provided dictionary. 

This implementation solves the problem doing a Breadth-First Search.

Currently using 'lead' and 'gold' as origin and destination; it works with any pair de words in the dictionary as log as they have the same number of letters

Difficulties: 
	- The search for words or letters already stored in slices can be made faster, so I change the structures to maps of booleans. This is for reading the file and creating the dictionary of words and letters.
	- Possibly a large number of words: this needs space and speed. Once the structs are created, I decide to use []string to save the paths and those are stored in a   [][]string to access each path easily

Possible future changes:
	- Change from the current implementation to another more OOP, with structs that map each path or node in the searchs