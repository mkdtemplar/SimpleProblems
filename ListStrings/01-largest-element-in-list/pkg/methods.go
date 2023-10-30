package pkg

type List struct {
	list []int
}

func (l *List) LargestElement() int {
	maxElement := l.list[0]

	for _, e := range l.list {
		if e > maxElement {
			maxElement = e
		}
	}
	return maxElement
}

func NewList(list []int) IList {
	return &List{list: list}
}

type IList interface {
	LargestElement() int
}
