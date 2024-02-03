package test

import (
	"reflect"
	"testing"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/04-odd-position-in-List/pkg"
)

func TestFindOddPosition(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	newSlice := pkg.NewSlice(inputSlice)

	want := []int{2, 4, 6, 8, 10}
	got := newSlice.FindOddPosition()

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test failed want %v, got %v", want, got)
	}
}
