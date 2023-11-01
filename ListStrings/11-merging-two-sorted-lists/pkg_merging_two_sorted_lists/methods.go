package pkg_merging_two_sorted_lists

import (
	"sort"
)

type TwoLists struct {
	listOne []int
	listTwo []int
}

func (t *TwoLists) MergeTwoSortedSlices() []int {

	var concatenatedSlice []int

	concatenatedSlice = append(t.listOne, t.listTwo...)

	concatenatedSlice = SortSlice(concatenatedSlice)

	return concatenatedSlice
}

func SortSlice(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	return slice
}

func NewTwoLists(listOne []int, listTwo []int) ITwoLists {
	return &TwoLists{listOne: listOne, listTwo: listTwo}
}

type ITwoLists interface {
	MergeTwoSortedSlices() []int
}
