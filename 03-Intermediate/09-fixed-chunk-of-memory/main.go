package main

import (
	"errors"
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/Intermediate/08-interface-for-List/pkg"
)

type List pkg.List

//Implement the interface from other package
//Add new is valid method
//new constructor which returns both list and error

func NewList(l List) (pkg.ListInterface, error) {
	if len(l.List) <= 100 {
		return pkg.NewList(l.List), nil
	}
	return nil, errors.New("out of memory")
}

func main() {
	slice := List{List: []int{5, 1, 7, 9, 0, 32, 34, 56, 8}}

	list, err := NewList(slice)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(list.MinElement())
}
