package main

import (
	"fmt"

	pkgrotatesliceelements "github.com/mkdtemplar/SimpleProblems/02-ListStrings/12-rotates-a-list-by-k-elements/pkg_rotate_slice_elements"
)

func main() {
	exampleSlice := []int{1, 2, 3, 4, 5, 6}
	input1 := []int{1, 2, 3, 4, 5, 6}
	input2 := []int{1, 2, 3, 4, 5, 6}
	var k int
	newPkg1 := pkgrotatesliceelements.NewRotate(input1)
	newPkg2 := pkgrotatesliceelements.NewRotate(input2)

	fmt.Print("Enter number of rotations: ")
	_, err := fmt.Scan(&k)
	if err != nil {
		return
	}

	fmt.Printf("Rotation of elements of slice %v for %d  positions left %v\n", exampleSlice, k, newPkg1.RotateElementsLeft(k))

	fmt.Printf("Rotation of elements of slice %v for %d  positions right %v\n", exampleSlice, k, newPkg2.RotateElementsRight(k))

}
