package tests

import (
	"reflect"
	"testing"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/11-merging-two-sorted-lists/pkg_merging_two_sorted_lists"
)

func TestSortSlice(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 7}
	input := []int{4, 1, 7, 2, 5, 6, 3}
	got := pkg_merging_two_sorted_lists.SortSlice(input)

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test failed want %v, got %v", want, got)
	}
}

func TestMergeTwoSortedSlices(t *testing.T) {
	input1 := []int{5, 2, 7, 3, 8}
	input2 := []int{300, 123, 2, 1, 678}

	want := []int{1, 2, 2, 3, 5, 7, 8, 123, 300, 678}

	newPkg := pkg_merging_two_sorted_lists.NewTwoLists(input1, input2)

	got := newPkg.MergeTwoSortedSlices()

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test failed want %v, got %v", want, got)
	}
}
