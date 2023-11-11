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
func MergeSort[T constraints.Ordered](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])
	return merge(left, right)
}

func Quicksort[T constraints.Ordered](a []T) []T {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	Quicksort(a[:left])
	Quicksort(a[left+1:])

	return a
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
