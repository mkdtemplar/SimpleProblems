package tests

import (
	"testing"

	pkgforarithmetic "github.com/mkdtemplar/SimpleProblems/ListStrings/15-arithmetic-operations/pkg-for-arithmetic"
)

func TestAddition(t *testing.T) {
	want := "12666666666666666666654"

	newPkg := pkgforarithmetic.NewNumbers("999999999999999999999", "999999999999999999999")

	got := newPkg.Addition()

	if want != got {
		t.Errorf("Test failed want %s got %s", want, got)
	}
}
