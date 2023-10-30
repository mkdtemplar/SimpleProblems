package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/04-odd-position-in-list/pkg"
)

func main() {
	inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	newSlice := pkg.NewSlice(inputSlice)

	fmt.Println(newSlice.FindOddPosition())
}
