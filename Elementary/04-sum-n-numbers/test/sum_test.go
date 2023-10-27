package test

import (
	"testing"

	"github.com/mkdtemplar/SimpleProblems/Elementary/04-sum-n-numbers/pkg"
)

func TestPrintSum(t *testing.T) {
	want := 55

	got := pkg.PrintSum(10)

	if want != got {
		t.Errorf("Test failed want %d, got %d", want, got)
	}
}

func TestPrintSumRecursively(t *testing.T) {
	want := 55

	got := pkg.PrintSum(10)

	if want != got {
		t.Errorf("Test failed want %d, got %d", want, got)
	}
}
