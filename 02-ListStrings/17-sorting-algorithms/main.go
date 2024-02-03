package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/17-sorting-algorithms/methods"
)

func algorithmChoice(inputSlice []int) {

	var choiceSwitch int
	var reRun = true
	var reRunChoice string

	for reRun {
		fmt.Println("1 - Selection sort, 2 - Insertion sort, 3 - MergeSort, 4 - Quicksort, 5 - StoogeSort")
		fmt.Print("Enter choice for sorting algorithm: ")
		_, err := fmt.Scan(&choiceSwitch)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if !slices.Contains([]int{1, 2, 3, 4, 5}, choiceSwitch) {
			algorithmChoice(inputSlice)
		}
		switch choiceSwitch {
		case 1:
			fmt.Println(methods.SelectionSort(inputSlice))
		case 2:
			fmt.Println(methods.InsertionSort(inputSlice))
		case 3:
			fmt.Printf("Sorting slice %v using Merge sort\n", inputSlice)
			fmt.Println(methods.MergeSort(inputSlice))
		case 4:
			fmt.Printf("Sorting slice %v using Quick sort\n", inputSlice)
			fmt.Println(methods.Quicksort(inputSlice))
		case 5:
			fmt.Printf("Sorting slice %v using Stooge sort\n", inputSlice)
			fmt.Println(methods.StoogeSort(inputSlice, 0, len(inputSlice)-1))
		}

		fmt.Print("Do you want to continue Y/N: ")
		_, err = fmt.Scan(&reRunChoice)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		reRunChoice = strings.ToLower(reRunChoice)
		if reRunChoice != "y" {
			reRun = false
		}
	}

}

func main() {
	var inputSlice = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	algorithmChoice(inputSlice)

}
