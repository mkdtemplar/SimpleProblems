package pkg_combines_two_lists

import (
	"fmt"
	"math"
)

type TwoLists[T any] struct {
	listOne []T
	listTwo []T
}

func (t TwoLists[T]) CombinesTwoSlices() []string {

	var (
		j int
		k int
	)

	maxLength := int(math.Max(float64(len(t.listOne)), float64(len(t.listTwo))))
	var concatenatedSlice []string
	converted1 := ConvertToString(t.listOne)

	converted2 := ConvertToString(t.listTwo)

	for i := 0; i < maxLength; i++ {
		if j < len(converted1) {
			concatenatedSlice = append(concatenatedSlice, converted1[j])
			j++
		}
		if k < len(converted2) {
			concatenatedSlice = append(concatenatedSlice, converted2[k])
			k++
		}
	}

	return concatenatedSlice
}

func ConvertToString[T any](str []T) []string {
	var converted []string
	for _, e := range str {
		converted = append(converted, fmt.Sprintf("%v", e))
	}
	return converted
}

func NewTwoLists[T any](listOne []T, listTwo []T) ITwoLists {
	return &TwoLists[T]{listOne: listOne, listTwo: listTwo}
}

type ITwoLists interface {
	CombinesTwoSlices() []string
}
