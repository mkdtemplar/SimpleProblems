package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](a []T, x T) int {
	// Set low and high boundaries
	low, high := 0, len(a)-1

	for low <= high {
		mid := low + (high-low)/2
		if a[mid] == x {
			return mid
		}

		if a[mid] < x {
			// If target is greater, focus on the right half
			low = mid + 1
		} else {
			// If target is less, focus on the left half
			high = mid - 1
		}
	}
	return -1
}

func main() {

	var input = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var findElement int

	fmt.Print("Enter value of a element to search in the slice: ")
	_, err := fmt.Scan(&findElement)
	if err != nil {
		return
	}

	result := BinarySearch(input, findElement)

	if result != -1 {
		fmt.Printf("Eleent %d found at index %d\n", findElement, result)
	} else {
		fmt.Printf("Element %d not found\n", findElement)
	}

}
