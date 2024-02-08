package main

import ( 
	"fmt" 
	"sort" 
)

func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * factorial(n - 1)
}

func binarySearch(array []int, target int) int {
	low, high := 0, len(array)-1

	for low <= high {
		mid := (low + high) / 2
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1 
		}
	}

	return -1
}

func bubbleSort(array []int){
	n := len(array)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if(array[j] < array[j+1]){
				array[j], array[j+1] = array[j+1], array[j] 
			}
		}
	}
}

func quickSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	pivot := array[len(array)/2]

	var left, right []int

	for _, value := range array {
		if value < pivot {
			left = append(left, value)
		}else if value > pivot {
			right = append(right, value)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	array := []int{64, 34, 25, 12, 22, 11, 90}
	array2 := []int{64, 34, 25, 12, 22, 11, 90}
	array3 := []int{64, 34, 25, 12, 22, 11, 90}
	target := 11
	num := 5
	
	// Bubble Sort
	fmt.Println("Array before reordering", array)
	bubbleSort(array)
	fmt.Println("Array after reordering with bubbleSort", array)
	
	// Quick Sort
	fmt.Println("Array before reordering", array2)
	quick := quickSort(array2)
	fmt.Println("Array after reordering with quickSort", quick)
	
	// Sort imported
	fmt.Println("Array before reordering", array3)
	sort.Ints(array3)
	fmt.Println("Array after reordering with sort imported", array3)
	
	// Binary search
	binary := binarySearch(array, target)
	fmt.Println("The array", array)
	fmt.Println("The target", target)
	fmt.Println("The result of binary search on the array", binary)

	// Factorial
	factorialResult := factorial(num)
	fmt.Println("Factorial of", num,  factorialResult)
	

}
