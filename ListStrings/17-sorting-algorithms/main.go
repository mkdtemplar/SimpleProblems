package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/17-sorting-algorithms/methods"
)

func main() {
	var inputSlice = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	var choice int
	fmt.Println("1 - Selection sort, 2 - Insertion sort, 3 - MergeSort, 4 - Quicksort, 5 - StoogeSort")
	fmt.Print("Enter choice for sorting algorithm: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		return
	}

	switch choice {
	case 1:
		fmt.Println(methods.SelectionSort(inputSlice))
	case 2:
		fmt.Println(methods.InsertionSort(inputSlice))
	case 3:
		fmt.Println(methods.MergeSort(inputSlice))
	case 4:
		fmt.Println(methods.Quicksort(inputSlice))
	case 5:
		methods.StoogeSort(inputSlice, 0, len(inputSlice)-1)
	default:
		fmt.Println("Wrong choice")
		return
	}

}
