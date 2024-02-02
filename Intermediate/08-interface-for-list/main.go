package main

import (
	"fmt"

	sortAlg "github.com/mkdtemplar/SimpleProblems/ListStrings/17-sorting-algorithms/methods"
)

type Lists struct {
	list []int
}

func (l *Lists) maxElement() string {
	return fmt.Sprintf("Max element in the list is %v", sortAlg.MergeSort(l.list)[len(l.list)-1])
}

func (l *Lists) minElement() string {
	return fmt.Sprintf("Max element in the list is %v", sortAlg.MergeSort(l.list)[0])
}

func (l *Lists) find(element int) string {
	if binarySearch(sortAlg.MergeSort(l.list), element) == -1 {
		return fmt.Sprintf("Element %v not found!", element)
	}
	return fmt.Sprintf("Element %v found in the list", element)
}

type ListInterface[T comparable] interface {
	find(T) T
	maxElement() T
	minElement() T
}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	list := Lists{list: []int{5, 1, 7, 9, 0, 32, 34, 56, 8}}
	fmt.Println(list.find(32))
	fmt.Println(list.maxElement())
	fmt.Println(list.minElement())

}
