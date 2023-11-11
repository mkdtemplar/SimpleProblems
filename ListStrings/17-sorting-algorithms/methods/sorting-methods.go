package methods

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

func SelectionSort[T constraints.Ordered](s []T) []T {
	var minIndex int
	var temp T
	for i := 0; i < len(s)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[minIndex] {
				minIndex = j
			}
		}
		temp = s[i]
		s[i] = s[minIndex]
		s[minIndex] = temp
	}
	return s
}

func InsertionSort[T constraints.Ordered](s []T) []T {
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[j] > s[i] {
				s[j], s[i] = s[i], s[j]
			}
		}
	}

	return s
}

func merge[T constraints.Ordered](left, right []T) []T {
	sortedArray := make([]T, len(left)+len(right))
	i, j := 0, 0
	for k := 0; k < len(sortedArray); k++ {
		if i >= len(left) {
			sortedArray[k] = right[j]
			j++
		} else if j >= len(right) {
			sortedArray[k] = left[i]
			i++
		} else if left[i] < right[j] {
			sortedArray[k] = left[i]
			i++
		} else {
			sortedArray[k] = right[j]
			j++
		}
	}
	return sortedArray
}
func MergeSort[T constraints.Ordered](s []T) []T {
	if len(s) <= 1 {
		return s
	}
	middle := len(s) / 2
	left := MergeSort(s[:middle])
	right := MergeSort(s[middle:])
	return merge(left, right)
}

func Quicksort[T constraints.Ordered](s []T) []T {
	if len(s) < 2 {
		return s
	}

	left, right := 0, len(s)-1

	pivot := rand.Int() % len(s)

	s[pivot], s[right] = s[right], s[pivot]

	for i := range s {
		if s[i] < s[right] {
			s[left], s[i] = s[i], s[left]
			left++
		}
	}

	s[left], s[right] = s[right], s[left]

	Quicksort(s[:left])
	Quicksort(s[left+1:])

	return s
}

func StoogeSort[T constraints.Ordered](s []T, low int, high int) []T {
	if low >= high {
		return s
	}

	if s[low] > s[high] {
		temp := s[low]
		s[low] = s[high]
		s[high] = temp
	}

	if high-low+1 > 2 {
		temp := (high - low + 1) / 3
		StoogeSort(s, low, high-temp)
		StoogeSort(s, low+temp, high)
		StoogeSort(s, low, high-temp)
	}

	return s
}
