package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/17-sorting-algorithms/methods"
)

func main() {
	var inputSlice = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//var choiceSlice = [4]int{1, 2, 3, 4}
	var choiceSwitch int
	var reRunChoice = true
	var choiceString string

	fmt.Println("1 - Selection sort, 2 - Insertion sort, 3 - MergeSort, 4 - Quicksort, 5 - StoogeSort")
	fmt.Print("Enter choiceSwitch for sorting algorithm: ")
	_, err := fmt.Scan(&choiceSwitch)
	if err != nil {
		return
	}

	for reRunChoice {
		switch choiceSwitch {
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
			fmt.Println("Wrong choiceSwitch")
			return
		}

		fmt.Print("Do you want to continue Y/N: ")
		_, err := fmt.Scan(&choiceString)
		if err != nil {
			return
		}

		if choiceString == "Y" {
			fmt.Println("1 - Selection sort, 2 - Insertion sort, 3 - MergeSort, 4 - Quicksort, 5 - StoogeSort")
			fmt.Print("Enter choiceSwitch for sorting algorithm: ")
			_, err := fmt.Scan(&choiceSwitch)
			if err != nil {
				return
			}
		} else if choiceString == "N" {
			reRunChoice = false
		}
	}

}
