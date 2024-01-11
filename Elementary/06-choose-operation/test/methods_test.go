package test

import (
	"testing"

	"github.com/mkdtemplar/SimpleProblems/Elementary/06-choose-operation/pkg"
)

func TestComputeSum(t *testing.T) {
	want := 55

	got := pkg.ComputeSum(10)

	if want != got {
		t.Errorf("Test failed want %d, got %d", want, got)
	}
}

func TestComputeProduct(t *testing.T) {
	want := 3628800

	got := pkg.ComputeProduct(10)

	if want != got {
		t.Errorf("Test failed want %d, got %d", want, got)
	}
}
