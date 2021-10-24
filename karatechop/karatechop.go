package karatechop

import (
  "fmt"
  "math"
  "sort"
  "strconv"
)


// Different approaches to the same problem: karate chop = binary search
// All different methods receive a number and an array, and return the position the number is in the array, or -1 if not in the array
func Start () {
  fmt.Println("Checking Chop methods")

  lists := [][]int{{}, {1}, {1}, {1, 3, 5}, {1, 3, 5}, {1, 3, 5}, {1, 3, 5}, {1, 3, 5}, {1, 3, 5}, {1, 3, 5}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}, {1, 3, 5, 7}}
  values := []int{3, 3, 1, 1, 3, 5, 0, 2, 4, 6, 1, 3, 5, 7, 0, 2, 4, 6, 8}
  results := []int{-1, -1, 0, 0, 1, 2, -1, -1, -1, -1, 0, 1, 2, 3, -1, -1, -1, -1, -1}

  for i, _ := range results {
    result := results[i]
    value  := values[i]
    list   := lists[i]

    if !CheckParameters(list) {
      test := SliceSplittingChop(list, value)

      if result != test {
        fmt.Print("Something went wrong: SliceSplittingChop(" + fmt.Sprint(list) + ", " + strconv.Itoa(value) + ") != " + strconv.Itoa(result) + ": " + strconv.Itoa(test) + "???")
      } 
    }
   // fmt.Println("---------------------")
  }
}

// Iterative approach: in a loop this shortens the limits of the list
func IterativeChop (list []int, value int) (int) {

  beginningOfList := 0
  endOfList  := len(list)-1

  var middleOfList int

  for beginningOfList <= endOfList {

    middleOfList = (beginningOfList + endOfList)/2
    valueMiddle := list[middleOfList]

    if (value == valueMiddle) {
      // Valid when the array has an odd number of positions, including only one position
      return middleOfList
    }

    if value < valueMiddle {
      endOfList = middleOfList -1

    } else { //if valueMiddle > value {
      beginningOfList = middleOfList + 1
    } 
  }
	return -1
}

// Recursive method: it calls itself sending each time positions for a smaller list
func RecursiveChop (list []int, value int, beginningOfList int, endOfList int) (int) {

  if beginningOfList > endOfList {
    return -1  
  }

  middle := (beginningOfList + endOfList)/2
  valueMiddle := list[middle]


  if value == valueMiddle {
    return middle
  }

  if value < valueMiddle {
    return RecursiveChop(list, value, beginningOfList, middle - 1)
  } else {
    return RecursiveChop(list, value, middle + 1, endOfList)
  }	
}

// Slice splitting
func SliceSplittingChop (list []int, value int) (int) {

  beginningOfList := 0
  endOfList  := len(list)-1

  position := -1

  for {
    if len(list) == 0 || (len(list) == 1 && list[0] != value) {
      return -1  
    }
    beginningOfList = 0
    endOfList  = len(list)-1

    middle := (beginningOfList + endOfList)/2
    valueMiddle := list[middle]

    //fmt.Println("middle: " + strconv.Itoa(middle))
    //fmt.Println("valueMiddle: " + strconv.Itoa(valueMiddle))

    if value == valueMiddle {
      if position == -1 {
        return middle
      } else {
        return position
      }
    }

    // As we need to use the variable initialPosition, we update its base value so we can edit it in the list splice methods
    if (position == -1) {
      position = 0
    }

    if value < valueMiddle {
      position, list = subSliceLeft(position, list)
    } else {
      position, list = subSliceRight(position, list)
    } 
  }
  return -1
}

// Recursive with custom splicing
func RecursiveAndSliceSplittingChop (initialPosition int, list []int, value int) (int) {

  if len(list) == 0 || (len(list) == 1 && list[0] != value) {
    return -1  
  }
  beginningOfList := 0
  endOfList  := len(list)-1

  middle := (beginningOfList + endOfList)/2
  valueMiddle := list[middle]

  if value == valueMiddle {
    if initialPosition == -1 {
        return middle
      } else {
        return initialPosition  
      }
  }

  // As we need to use the variable initialPosition, we update its base value so we can edit it in the list splice methods
  if (initialPosition == -1) {
    initialPosition = 0
  }

  if value < valueMiddle {
    newInitialPosition, list := subSliceLeft(initialPosition, list)
    return RecursiveAndSliceSplittingChop(newInitialPosition, list, value)//, value, beginningOfList, middle - 1)
  } else {
    newInitialPosition, list := subSliceRight(initialPosition, list)
    return RecursiveAndSliceSplittingChop(newInitialPosition, list, value)//, value, middle + 1, endOfList)
  } 
}

// Binary search implemented in standard library: https://pkg.go.dev/sort#SearchInts
// We have to fix the result because when the value isn't located in the list, it returns "the index at which the item could be placed to maintain sort order"
func StandardLibraryChop (list []int, value int) (int) {

  result := sort.SearchInts(list, value)

  if result == len(list) || list[result] != value {

    return -1
  }

  return result
}

// Returns whether both the list and value are valid before starting any chop method
// Checks if the list isn't empty. We assume a sorted array
func CheckParameters (list []int) (bool) {

  return (list == nil)
}

// Returns the right half of the list, rounding up (3 positions -> 2 positions, 4 positions -> 2 positions...)
func subSliceRight (initialPosition int, list []int) (int, []int) {
  sliceHalf := len(list)/2
  roundSliceHalf := int(math.Round(float64(len(list))/2))
  result := make([]int, roundSliceHalf)

  copy(result, list[sliceHalf : len(list)])
  //fmt.Println(fmt.Sprint(list) + " -> " + fmt.Sprint(result))

  //fmt.Println("returns " + strconv.Itoa(initialPosition + (len(list) - len(result))))
  return (initialPosition + (len(list) - len(result))), result
}


// Returns the left half of the list, rounding up (3 positions -> 2 positions, 4 positions -> 2 positions...)
func subSliceLeft (initialPosition int, list []int) (int, []int){
  sliceHalf := len(list)/2
  roundSliceHalf := int(math.Round(float64(len(list))/2))
  result := make([]int, roundSliceHalf)

  copy(result, list[0 : sliceHalf])
  //fmt.Println(fmt.Sprint(list) + " -> " + fmt.Sprint(result))
  //fmt.Println("returns " + strconv.Itoa(initialPosition))

  return initialPosition, result
}