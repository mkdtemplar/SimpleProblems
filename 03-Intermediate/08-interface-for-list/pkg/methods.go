package pkg

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/17-sorting-algorithms/methods"
)

type ListInterface interface {
	Find(element int) string
	MaxElement() string
	MinElement() string
}

type List struct {
	List []int
}

func NewList(list []int) ListInterface {
	return &List{List: list}
}

func (l *List) MaxElement() string {
	return fmt.Sprintf("Max element in the List is %v", methods.MergeSort(l.List)[len(l.List)-1])
}

func (l *List) MinElement() string {
	return fmt.Sprintf("Max element in the List is %v", methods.MergeSort(l.List)[0])
}

func (l *List) Find(element int) string {
	if BinarySearch(methods.MergeSort(l.List), element) == -1 { // if BinarySearch
		return fmt.Sprintf("Element %v not found!", element)
	}
	return fmt.Sprintf("Element %v found in the List", element)
}

func BinarySearch(arr []int, target int) int { // IsInArray
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

	return -1 // bool
}
