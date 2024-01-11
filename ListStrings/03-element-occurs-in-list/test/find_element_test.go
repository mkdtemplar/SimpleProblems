package test

import (
	"testing"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/03-element-occurs-in-list/pkg"
)

func TestFindElement(t *testing.T) {
	sliceNew := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newSlice := pkg.NewSlice(sliceNew)

	got := newSlice.FindElement(10)

	if got != true {
		t.Errorf("Test failed want %v, got %v", true, got)
	}
}
