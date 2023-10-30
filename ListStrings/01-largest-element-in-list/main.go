package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/01-largest-element-in-list/pkg"
)

func main() {
	sliceNew := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	newPkg := pkg.NewList(sliceNew)

	fmt.Println(newPkg.LargestElement())
}
