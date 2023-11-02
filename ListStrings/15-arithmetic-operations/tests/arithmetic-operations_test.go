package tests

import (
	"testing"

	pkgforarithmetic "github.com/mkdtemplar/SimpleProblems/ListStrings/15-arithmetic-operations/pkg-for-arithmetic"
)

func TestAddition(t *testing.T) {
	want := "19998"

	newPkg := pkgforarithmetic.NewNumbers(9999, 9999)

	got := newPkg.Addition()

	if want != got {
		t.Errorf("Test failed want %s got %s", want, got)
	}
}
