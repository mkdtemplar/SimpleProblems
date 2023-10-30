package pkg

type Slice struct {
	newSlice []int
}

// FindOddPosition implements ISlice.
func (s *Slice) FindOddPosition() []int {
	var oddElements []int
	for i := 0; i < len(s.newSlice); i++ {
		if i%2 != 0 {
			oddElements = append(oddElements, s.newSlice[i])
		}
	}
	return oddElements
}

type ISlice interface {
	FindOddPosition() []int
}

func NewSlice(newSlice []int) ISlice {
	return &Slice{newSlice: newSlice}
}
