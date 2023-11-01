package test

import (
	"reflect"
	"testing"

	pkgrotatesliceelements "github.com/mkdtemplar/SimpleProblems/ListStrings/12-rotates-a-list-by-k-elements/pkg_rotate_slice_elements"
)

func TestLeftRotation(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}

	want := []int{2, 3, 4, 5, 6, 1}

	newPkg := pkgrotatesliceelements.NewRotate(input)

	got := newPkg.LeftRotation()

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test failed want %v, got %v", want, got)
	}
}

func TestRightRotation(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}

	want := []int{6, 1, 2, 3, 4, 5}

	newPkg := pkgrotatesliceelements.NewRotate(input)

	got := newPkg.RightRotation()

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test failed want %v, got %v", want, got)
	}
}

func TestRotateElementsLeft(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}

	newPkg := pkgrotatesliceelements.NewRotate(input)

	got := newPkg.RotateElementsLeft(2)

	want := []int{3, 4, 5, 6, 1, 2}

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test failed want %v, got %v", want, got)
	}
}

func TestRotateElementsRight(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}

	newPkg := pkgrotatesliceelements.NewRotate(input)

	got := newPkg.RotateElementsRight(2)

	want := []int{5, 6, 1, 2, 3, 4}

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("Test failed want %v, got %v", want, got)
	}
}
