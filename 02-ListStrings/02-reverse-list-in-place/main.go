package main

import "fmt"

type Slice[T any] struct {
	inputSlice []T
}
type ISlice[T any] interface {
	reverseList() []T
}

func (s *Slice[T]) reverseList() []T {
	for i, j := 0, len(s.inputSlice)-1; i < j; i, j = i+1, j-1 {
		s.inputSlice[i], s.inputSlice[j] = s.inputSlice[j], s.inputSlice[i]
	}
	return s.inputSlice
}

func NewSlice[T any](inputSlice []T) ISlice[T] {
	return &Slice[T]{inputSlice: inputSlice}
}

func main() {
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	newSlice := NewSlice(s)

	fmt.Println(newSlice.reverseList())
}
