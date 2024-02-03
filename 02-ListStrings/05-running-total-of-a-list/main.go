package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/05-running-total-of-a-list/pkg"
)

func main() {
	inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	newSlice := pkg.NewSumSlice[int](inputSlice)

	fmt.Println(newSlice.SumOfSliceElements())
}
