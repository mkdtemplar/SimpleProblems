package test

import (
	"math"

	pkgperfectsquares "github.com/mkdtemplar/SimpleProblems/ListStrings/08-perfect-squares/pkg-perfect-squares"
)

type Tests []struct {
	input int
	want  int
}

func PopulateTestStruct() Tests {
	var tests Tests
	var inputSlice = pkgperfectsquares.CreateSlice()
	for _, element := range inputSlice {
		tests = append(tests, struct {
			input int
			want  int
		}{input: element, want: int(math.Pow(float64(element), 2))})
	}

	return tests
}
