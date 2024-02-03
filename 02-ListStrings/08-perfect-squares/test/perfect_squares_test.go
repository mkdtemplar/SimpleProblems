package test

import (
	"strconv"
	"testing"

	pkgperfectsquares "github.com/mkdtemplar/SimpleProblems/ListStrings/08-perfect-squares/pkg-perfect-squares"
)

func TestCalculateSquare(t *testing.T) {

	tests := PopulateTestStruct()

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			got := pkgperfectsquares.CalculatePerfectSquare(tt.input)
			if got != tt.want {
				t.Errorf("CalculateSquare(%d) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}

}

func TestPerfectSquare(t *testing.T) {
	tests := PopulateTestStruct()
	inputSlice := pkgperfectsquares.CreateSlice()
	newPkg := pkgperfectsquares.NewSlicePerfectSquares(inputSlice)
	for i, tt := range tests {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			got := newPkg.PerfectSquare(pkgperfectsquares.CalculatePerfectSquare)
			if got[i] != tt.want {
				t.Errorf("CalculateSquare(%d) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
