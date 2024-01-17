package test

import (
	"testing"

	"github.com/mkdtemplar/SimpleProblems/Elementary/10-print-leap-years/pkg"
)

func TestCheckLeapYear(t *testing.T) {
	got := pkg.CheckLeapYear(2024)

	if got == false {
		t.Errorf("Test failed want %v, got %v", true, got)
	}
}
