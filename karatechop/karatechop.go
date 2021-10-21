package karatechop

// Different approaches to the same problem: karate chop = binary search
// All different methods receive a number and an array, and return the position the number is in the array, or -1 if not in the array


func Chop1(number int, list []int) (int) {

	return -1
}

func Chop2(number int, list []int) (int) {

	return -1
}

func Chop3(number int, list []int) (int) {

	return -1
}

func Chop4(number int, list []int) (int) {

	return -1
}

func Chop5(number int, list []int) (int) {

	return -1
}

/*
// A recursive binary search function. It returns location of x in
// given array arr[l..r] is present, otherwise -1
int binarySearch(int arr[], int l, int r, int x)
{
   if (r >= l)
   {
        int mid = l + (r - l)/2;
  
        // If the element is present at the middle itself
        if (arr[mid] == x)  return mid;
  
        // If element is smaller than mid, then it can only be present
        // in left subarray
        if (arr[mid] > x) return binarySearch(arr, l, mid-1, x);
  
        // Else the element can only be present in right subarray
        return binarySearch(arr, mid+1, r, x);
   }
  
   // We reach here when element is not present in array
   return -1;
}
*/


/*
func binarySearch(needle int, haystack []int) bool {

  low := 0
  high := len(haystack) - 1

  for low <= high{
    median := (low + high) / 2

    if haystack[median] < needle {
      low = median + 1
    }else{
      high = median - 1
    }
  }

  if low == len(haystack) || haystack[low] != needle {
    return false
  }

  return true
}


func main(){
  items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
  fmt.Println(binarySearch(63, items))
}*/