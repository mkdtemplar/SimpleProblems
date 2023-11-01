package pkgrotatesliceelements

type RotateSlice[T any] struct {
	newSlice []T
}

type IRotateElements[T any] interface {
	RotateElementsLeft(k int) []T
	RotateElementsRight(k int) []T
	LeftRotation() []T
	RightRotation() []T
}

func NewRotate[T any](newSlice []T) IRotateElements[T] {
	return &RotateSlice[T]{newSlice: newSlice}
}

func (r *RotateSlice[T]) LeftRotation() []T {
	lastElement := r.newSlice[0]
	for j := 0; j < len(r.newSlice)-1; j++ {
		r.newSlice[j] = r.newSlice[j+1]
	}
	r.newSlice[len(r.newSlice)-1] = lastElement
	return r.newSlice
}

func (r *RotateSlice[T]) RightRotation() []T {
	lastElement := r.newSlice[len(r.newSlice)-1]
	for j := len(r.newSlice) - 1; j > 0; j-- {
		r.newSlice[j] = r.newSlice[j-1]
	}
	r.newSlice[0] = lastElement
	return r.newSlice
}

// RotateElementsLeft implements IRotateElements.
func (r *RotateSlice[T]) RotateElementsLeft(k int) []T {
	for i := 1; i <= k; i++ {
		r.LeftRotation()
	}

	return r.newSlice
}

func (r *RotateSlice[T]) RotateElementsRight(k int) []T {
	for i := 1; i <= k; i++ {
		r.RightRotation()
	}

	return r.newSlice
}
