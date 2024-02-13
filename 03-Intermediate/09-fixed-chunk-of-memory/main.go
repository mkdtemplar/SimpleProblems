package main

import (
	"github.com/mkdtemplar/SimpleProblems/Intermediate/08-interface-for-List/pkg"
)

type List pkg.List

func (l List) Find(element int) string {
	//TODO implement me
	panic("implement me")
}

func (l List) MaxElement() string {
	//TODO implement me
	panic("implement me")
}

func (l List) MinElement() string {
	//TODO implement me
	panic("implement me")
}

func (l List) isValid() (bool, error) {
	//TODO implement me
	panic("implement me")
}

type ValidList interface {
	pkg.ListInterface
	isValid() (bool, error)
}

// Implement the interface from other package
// Add new is valid method
// newconstructor which returns both list and error

//func NewList(l List) (pkg.ListInterface, error) {
//	if len(l.List) == 100 {
//		return pkg.NewList(l.List), nil
//	}
//	return nil, errors.New("out of memory")
//}

func NewList(list List) ValidList {
	return &list
}
func main() {
	slice := List{List: []int{5, 1, 7, 9, 0, 32, 34, 56, 8}}
	for i := 0; i < 123; i++ {
		slice.List = append(slice.List, i)
	}
	list := NewList(slice)

	list.MinElement()

}
