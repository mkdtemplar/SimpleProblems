package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/02-ListStrings/01-largest-element-in-List/pkg"
)

func main() {
	sliceNew := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sliceNewString := []string{"a", "b", "c", "d"}

	sliceNewSpecial := []string{"!", "{", "_", "*", "&"}

	newPkg := pkg.NewList(sliceNew)

	newPkgString := pkg.NewList(sliceNewString)

	newPkgSpecial := pkg.NewList(sliceNewSpecial)

	fmt.Println(newPkg.LargestElement())

	fmt.Println(newPkgString.LargestElement())

	fmt.Println(newPkgSpecial.LargestElement())
}
