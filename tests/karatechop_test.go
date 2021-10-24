package tests 

import (
	"testing"

	"cycloid/karatechop"

	"fmt"
	"strconv"
)

var (
	lists 	[][]int
	values	[]int
	results []int
)

// en el directorio tests/ => go test -v

func startVariables () () {
	if lists == nil || values == nil || results == nil  {
		lists = [][]int{{}, {1}, {1}, {1, 3, 5}, {1, 3, 5}, {1, 3, 5}, {1, 3, 5}, {1, 3, 5}, {1, 3, 5}, {1, 3, 5}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}}
		values = []int{3, 3, 1, 1, 3, 5, 0, 2, 4, 6, 1, 3, 5, 7, 0, 2, 4, 6, 8}
  		results = []int{-1, -1, 0, 0, 1, 2, -1, -1, -1, -1, 0, 1, 2, 3, -1, -1, -1, -1, -1}
  	}
}

func TestIterativeChop (t *testing.T) {

	startVariables()

	for i, _ := range lists {

		list   := lists[i]
	    value  := values[i]
	    result := results[i]

		if !karatechop.CheckParameters(list) {
			
			if result != karatechop.IterativeChop(list, value) {
				t.Errorf("Something went wrong: IterativeChop(" + fmt.Sprint(list) + ", " + strconv.Itoa(value) + ") != " + strconv.Itoa(result))	
      		}
    	} else {
    		t.Errorf("Invalid list: " + fmt.Sprint(list))
    	}
	}
}

func TestRecursiveChop (t *testing.T) {
	startVariables()

	for i, _ := range lists {

		list   := lists[i]
	    value  := values[i]
	    result := results[i]

		if !karatechop.CheckParameters(list) {
			
			if result != karatechop.RecursiveChop(list, value, 0, len(list) - 1 ) {
				t.Errorf("Something went wrong: RecursiveChop(" + fmt.Sprint(list) + ", " + strconv.Itoa(value) + ") != " + strconv.Itoa(result))	
      		}
    	} else {
    		t.Errorf("Invalid list: " + fmt.Sprint(list))
    	}
	}
}

func TestStandardLibraryChop (t *testing.T) {
	startVariables()

	for i, _ := range lists {

		list   := lists[i]
	    value  := values[i]
	    result := results[i]

		if !karatechop.CheckParameters(list) {
			
			if result != karatechop.StandardLibraryChop(list, value) {
				t.Errorf("Something went wrong: StandardLibraryChop(" + fmt.Sprint(list) + ", " + strconv.Itoa(value) + ") != " + strconv.Itoa(result))	
      		}
    	} else {
    		t.Errorf("Invalid list: " + fmt.Sprint(list))
    	}
	}
}

func TestSliceSplittingChop (t *testing.T) {
	startVariables()

	for i, _ := range lists {

		list   := lists[i]
	    value  := values[i]
	    result := results[i]

		if !karatechop.CheckParameters(list) {
			
			if result != karatechop.SliceSplittingChop(list, value) {
				t.Errorf("Something went wrong: SliceSplittingChop(" + fmt.Sprint(list) + ", " + strconv.Itoa(value) + ") != " + strconv.Itoa(result))	
      		}
    	} else {
    		t.Errorf("Invalid list: " + fmt.Sprint(list))
    	}
	}
}
/*
func TestRecursiveAndSliceSplittingChop (t *testing.T) {
	startVariables()

	for i, _ := range lists {

		list   := lists[i]
	    value  := values[i]
	    result := results[i]

		if !karatechop.CheckParameters(list) {
			
			if result != karatechop.RecursiveAndSliceSplittingChop(-1, list, value) {
				t.Errorf("Something went wrong: RecursiveAndSliceSplittingChop(" + fmt.Sprint(list) + ", " + strconv.Itoa(value) + ") != " + strconv.Itoa(result))	
      		}	
    	} else {
    		t.Errorf("Invalid list: " + fmt.Sprint(list))
    	} 
	}
}
*/