package tests

import (
	"reflect"
	"testing"

	pkg "github.com/mkdtemplar/SimpleProblems/ListStrings/10-combines-two-lists/pkg_combines_two_lists"
)

func TestConvertToString(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	want := []string{"1", "2", "3", "4", "5"}

	got := pkg.ConvertToString(input)

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test for ConvertToString failed! want %v, got %v", want, got)
	}
}

func TestConcatenatesTwoLists(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5}
	input2 := []string{"a", "b", "c"}

	want := []string{"1", "a", "2", "b", "3", "c", "4", "5"}

	slice1Converted := pkg.ConvertToString(input1)

	newPkg := pkg.NewTwoLists(slice1Converted, input2)

	got := newPkg.CombinesTwoSlices()

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test for ConcatenatesTwoLists failed! want %v, got %v", want, got)
	}

}
