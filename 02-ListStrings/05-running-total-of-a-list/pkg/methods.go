package pkg

import "golang.org/x/exp/constraints"

type Slice[T constraints.Integer | constraints.Float | constraints.Signed | constraints.Unsigned] struct {
	newSlice []T
}

// SumOfSliceElements implements ISlice.
func (s *Slice[T]) SumOfSliceElements() T {
	var sum T
	for _, v := range s.newSlice {
		sum += v
	}

	return sum
}

type ISlice[T constraints.Integer | constraints.Float | constraints.Signed | constraints.Unsigned] interface {
	SumOfSliceElements() T
}

func NewSumSlice[T constraints.Integer | constraints.Float | constraints.Signed | constraints.Unsigned](newSlice []T) ISlice[T] {
	return &Slice[T]{newSlice: newSlice}
}
