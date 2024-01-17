package test

import (
	"reflect"
	"testing"

	"github.com/mkdtemplar/SimpleProblems/Elementary/07-multiplication-table-12/pkg"
)

func TestCalculateProductIteratively(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got := pkg.CalculateProductIteratively(1)

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test failed want %v, got %v", want, got)
	}
}

func TestCalculateTable2DSlice(t *testing.T) {
	want := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
		{3, 6, 9, 12, 15, 18, 21, 24, 27, 30},
		{4, 8, 12, 16, 20, 24, 28, 32, 36, 40},
		{5, 10, 15, 20, 25, 30, 35, 40, 45, 50},
		{6, 12, 18, 24, 30, 36, 42, 48, 54, 60},
		{7, 14, 21, 28, 35, 42, 49, 56, 63, 70},
		{8, 16, 24, 32, 40, 48, 56, 64, 72, 80},
		{9, 18, 27, 36, 45, 54, 63, 72, 81, 90},
		{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
		{11, 22, 33, 44, 55, 66, 77, 88, 99, 110},
		{12, 24, 36, 48, 60, 72, 84, 96, 108, 120},
	}

	got := pkg.CalculateTable2DSlice()

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test failed want %v\n got %v\n", want, got)
	}
}
