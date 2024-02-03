package tests

import (
	"reflect"
	"testing"

	pkgforconcatenatestwolists "github.com/mkdtemplar/SimpleProblems/ListStrings/09-concatenates-two-lists/pkg-for-concatenates-two-lists"
)

func TestConvertToString(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	want := []string{"1", "2", "3", "4", "5"}

	got := pkgforconcatenatestwolists.ConvertToString(input)

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test for ConvertToString failed! want %v, got %v", want, got)
	}
}

func TestConcatenatesTwoLists(t *testing.T) {
	input1 := []int{1, 2, 3, 4, 5}
	input2 := []string{"a", "b", "c"}

	want := []string{"1", "2", "3", "4", "5", "a", "b", "c"}

	slice1Converted := pkgforconcatenatestwolists.ConvertToString(input1)

	newPkg := pkgforconcatenatestwolists.NewTwoLists(slice1Converted, input2)

	got := newPkg.ConcatenatesTwoLists()

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test for ConcatenatesTwoLists failed! want %v, got %v", want, got)
	}

}
