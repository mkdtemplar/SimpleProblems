package main

import (
	"fmt"

	pkgperfectsquares "github.com/mkdtemplar/SimpleProblems/ListStrings/08-perfect-squares/pkg-perfect-squares"
)

func main() {
	inputSlice := pkgperfectsquares.CreateSlice()

	newPkg := pkgperfectsquares.NewSlicePerfectSquares(inputSlice)

	fmt.Println(newPkg.PerfectSquare(pkgperfectsquares.CalculatePerfectSquare))
}
