package pkg

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/17-sorting-algorithms/methods"
)

type Lists struct {
	List []int
}

func (l *Lists) MaxElement() string {
	return fmt.Sprintf("Max element in the List is %v", methods.MergeSort(l.List)[len(l.List)-1])
}

func (l *Lists) MinElement() string {
	return fmt.Sprintf("Max element in the List is %v", methods.MergeSort(l.List)[0])
}

func (l *Lists) Find(element int) string {
	if BinarySearch(methods.MergeSort(l.List), element) == -1 {
		return fmt.Sprintf("Element %v not found!", element)
	}
	return fmt.Sprintf("Element %v found in the List", element)
}

type ListInterface[T comparable] interface {
	Find(T) T
	MaxElement() T
	MinElement() T
}

func BinarySearch(arr []int, target int) int {
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
