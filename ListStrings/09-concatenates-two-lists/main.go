package main

import (
	"fmt"

	pkgforconcatenatestwolists "github.com/mkdtemplar/SimpleProblems/ListStrings/09-concatenates-two-lists/pkg-for-concatenates-two-lists"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []string{"a", "b", "c"}

	slice1Converted := pkgforconcatenatestwolists.ConvertToString(slice1)

	newPkg := pkgforconcatenatestwolists.NewTwoLists(slice1Converted, slice2)

	fmt.Println(newPkg.ConcatenatesTwoLists())
}
