package pkg_for_concatenates_two_lists

import (
	"fmt"
)

type TwoLists[T any] struct {
	listOne []T
	listTwo []T
}

func (t TwoLists[T]) ConcatenatesTwoLists() []string {

	var concatenatedSlice []string

	converted1 := ConvertToString(t.listOne)

	converted2 := ConvertToString(t.listTwo)

	concatenatedSlice = append(converted1, converted2...)

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
	ConcatenatesTwoLists() []string
}
