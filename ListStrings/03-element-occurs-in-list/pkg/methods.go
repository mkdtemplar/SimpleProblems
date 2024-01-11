package pkg

import "fmt"

type Slice[T comparable] struct {
	newSlice []T
}

func (s *Slice[T]) Run() {
	var element T

	fmt.Print("Enter number to check if it is in the slice: ")
	_, err := fmt.Scan(&element)
	if err != nil {
		return
	}

	if s.FindElement(element) {
		fmt.Printf("Entered number %d, is in the slice\n", element)
	} else {
		fmt.Printf("Entered number %d, is not present in the slice\n", element)
	}
}

func (s *Slice[T]) FindElement(input T) bool {

	for _, element := range s.newSlice {
		if input == element {
			return true

		}
	}
	return false
}

func NewSlice[T comparable](newSlice []T) ISlice[T] {
	return &Slice[T]{newSlice: newSlice}
}

type ISlice[T comparable] interface {
	FindElement(input T) bool
	Run()
}
