package test

import (
	"testing"

	"github.com/mkdtemplar/SimpleProblems/02-ListStrings/01-largest-element-in-List/pkg"
)

func TestLargestElement(t *testing.T) {
	want := 10

	sliceNew := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	newPkg := pkg.NewList(sliceNew)

	got := newPkg.LargestElement()

	if want != got {
		t.Errorf("Test failed want %d, got %d: ", want, got)
	}
}
