package pkg

import "golang.org/x/exp/constraints"

type List[T constraints.Ordered] struct {
	list []T
}

func (l *List[T]) LargestElement() T {
	maxElement := l.list[0]

	for _, e := range l.list {
		if e > maxElement {
			maxElement = e
		}
	}
	return maxElement
}

func NewList[T constraints.Ordered](list []T) IList[T] {
	return &List[T]{list: list}
}

type IList[T constraints.Ordered] interface {
	LargestElement() T
}
