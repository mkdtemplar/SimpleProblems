package test

import (
	"reflect"
	"testing"

	"github.com/mkdtemplar/SimpleProblems/Elementary/05-multiples/pkg"
)

func TestPrintMultiplies(t *testing.T) {
	var multipliesSlice []int

	want := []int{3, 5, 6, 9, 10, 12, 15}

	multiplies := pkg.NewMultiplies(multipliesSlice)

	got := multiplies.PrintMultiplies(17)

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test failed want %v, got %v", want, got)
	}
}
