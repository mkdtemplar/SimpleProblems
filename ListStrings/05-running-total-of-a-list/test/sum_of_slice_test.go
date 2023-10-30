package test

import (
	"testing"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/05-running-total-of-a-list/pkg"
)

func TestSumOfSlice(t *testing.T) {
	inputSlice := []int{1,2,3,4,5,6,7,8,9,10}

	newSlice := pkg.NewSumSlice[int](inputSlice)

	want := 55

	got := newSlice.SumOfSliceElements()

	if want != got {
		t.Errorf("Test failed want %d, got %d", want, got)
	}
}
