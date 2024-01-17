package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](s []T, x T) int {
	low, high := 0, len(s)-1

	for low <= high {
		mid := low + (high-low)/2
		if s[mid] == x {
			return mid
		}

		if s[mid] < x {
			low = mid + 1
		} else {
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
