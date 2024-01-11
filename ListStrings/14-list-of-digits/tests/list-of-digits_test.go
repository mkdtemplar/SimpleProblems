package tests

import (
	"reflect"
	"testing"

	methodsforlistofdigits "github.com/mkdtemplar/SimpleProblems/ListStrings/14-list-of-digits/methods-for-list-of-digits"
)

func TestPrintListOfDigits(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	newPkg := methodsforlistofdigits.NewListOfDigits(1234567890)

	got := newPkg.PrintListOfDigits()

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test failed want %v, got %v", want, got)
	}
}
