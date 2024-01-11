package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/11-merging-two-sorted-lists/pkg_merging_two_sorted_lists"
)

func main() {
	input1 := []int{5, 2, 7, 3, 8}
	input2 := []int{300, 123, 2, 1, 678}

	newPkg := pkg_merging_two_sorted_lists.NewTwoLists(input1, input2)

	fmt.Println(newPkg.MergeTwoSortedSlices())

}
