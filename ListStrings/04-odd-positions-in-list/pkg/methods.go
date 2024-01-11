package pkg

type Slice[T any] struct {
	newSlice []T
}

// FindOddPosition implements ISlice.
func (s *Slice[T]) FindOddPosition() []T {
	var oddElements []T
	for i := 0; i < len(s.newSlice); i++ {
		if i%2 != 0 {
			oddElements = append(oddElements, s.newSlice[i])
		}
	}
	return oddElements
}

type ISlice[T any] interface {
	FindOddPosition() []T
}

func NewSlice[T any](newSlice []T) ISlice[T] {
	return &Slice[T]{newSlice: newSlice}
}
