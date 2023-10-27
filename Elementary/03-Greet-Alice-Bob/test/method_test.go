package test

import (
	"testing"

	"github.com/mkdtemplar/SimpleProblems/Elementary/03-Greet-Alice-Bob/pkg"
)

func TestPrintNames(t *testing.T) {

	want := "Hello Alice and Bob"

	names := pkg.NewNames([]string{"Alice", "Bob"})

	got := names.PrintNames()

	if want != got {
		t.Errorf("Test failed want %s, got %s", want, got)
	}
}
