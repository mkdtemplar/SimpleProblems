package main

import (
	"reflect"
	"testing"

	"github.com/mkdtemplar/SimpleProblems/Elementary/08-prime-numbers/pkg"
)

func TestPrimeNumbers(t *testing.T) {
	want := []int{2, 3, 5, 7, 11, 13, 17, 19}
	got := pkg.PrimeNumbers(20)

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test failed want %v got %v\n", want, got)
	}
}
