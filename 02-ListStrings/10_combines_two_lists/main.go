package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/ListStrings/10-combines-two-lists/pkg_combines_two_lists"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []string{"a", "b", "c"}

	slice1Converted := pkg_combines_two_lists.ConvertToString(slice1)

	newPkg := pkg_combines_two_lists.NewTwoLists(slice1Converted, slice2)

	fmt.Println(newPkg.CombinesTwoSlices())
}
