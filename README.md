# cycloid
Cycloid technical tests for backend job offer

(in project root folder)> go run main.go

-----------------------------------------------

* Karate Chop (http://codekata.com/kata/kata02-karate-chop/)

Given a sorted int array and a value, return the position of the value in the array, or -1 if the given value isn't in the array. For this problem we assume a sorted list and a value > 0.

The problem is doing 5 different ways
- Iterative method
- Recursive method
- Slice splitting
- Method from the standard Go library, with a small fix in the return to fulfill the exercise conditions
- A mix of recursive and slice splitting

Difficulties:
	- Finding different approaches for the same problem was more complicated than anticipated
	- During the creation of the RecursiveAndSliceSplittingChop I encountered the issue of having to split the list , saving the correct positions and values. This led to create custom methods to receive left and right parts of the list and saving the position of the first element RELATIVE TO THE ORIGINAL LIST

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