package main

import (
	"fmt"
	"github.com/mkdtemplar/SimpleProblems/ListStrings/07-three-functions/methods_for_task"
)

func main() {
	inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	newSlice := methods_for_task.NewSlice(inputSlice)

	fmt.Printf("Iterative sum of elements of slice: %d\n", newSlice.IterativeSumOfSlice())
	fmt.Printf("Conditional sum of elements of slice: %d\n", newSlice.ConditionalSumOfSlice())
	fmt.Printf("Recursive sum of elements of a slice: %d\n", newSlice.RecursiveSumOfSlice(len(inputSlice)))
}
