package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/Intermediate/08-interface-for-List/pkg"
)

func main() {
	list := pkg.NewLists([]int{5, 1, 7, 9, 0, 32, 34, 56, 8})
	fmt.Println(list.Find(32))
	fmt.Println(list.MaxElement())
	fmt.Println(list.MinElement())

}
