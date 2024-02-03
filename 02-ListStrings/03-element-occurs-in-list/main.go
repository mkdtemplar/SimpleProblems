package main

import "github.com/mkdtemplar/SimpleProblems/ListStrings/03-element-occurs-in-List/pkg"

func main() {

	sliceNew := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newSlice := pkg.NewSlice(sliceNew)

	newSlice.Run()
}
