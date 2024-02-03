package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/Intermediate/08-interface-for-List/pkg"
	"github.com/mkdtemplar/SimpleProblems/ListStrings/17-sorting-algorithms/methods"
)

type Lists pkg.Lists

func (l *Lists) MaxElement() string {
	if len(l.List) < 100 {
		return fmt.Sprintf("Max element in the List is %v", methods.MergeSort(l.List)[len(l.List)-1])
	} else {
		return fmt.Sprintf("Maximul length allowed of the slice is %d your slice is %d", 100, len(l.List))
	}
}

func (l *Lists) MinElement() string {
	if len(l.List) < 100 {
		return fmt.Sprintf("Max element in the List is %v", methods.MergeSort(l.List)[0])
	} else {
		return fmt.Sprintf("Maximul length allowed of the slice is %d your slice is %d", 100, len(l.List))
	}
}

func (l *Lists) Find(element int) string {
	if len(l.List) < 100 {
		if pkg.BinarySearch(methods.MergeSort[int](l.List), element) == -1 {
			return fmt.Sprintf("Element %v not found!", element)
		} else {
			return fmt.Sprintf("Element %v found in the List", element)
		}

	} else {
		return fmt.Sprintf("Maximul length allowed of the slice is %d your slice is %d", 100, len(l.List))
	}
}

func main() {
	slice := []int{5, 1, 7, 9, 0, 32, 34, 56, 8}
	for i := 0; i < 123; i++ {
		slice = append(slice, i)
	}
	list := Lists{List: slice}
	fmt.Println(list.Find(7))

}
