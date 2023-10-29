package methods_for_task

type Slice struct {
	NewSlice []int
}

type ISlice interface {
	IterativeSumOfSlice() int
	ConditionalSumOfSlice() int
	RecursiveSumOfSlice(length int) int
}

func NewSlice(newSlice []int) ISlice {
	return &Slice{NewSlice: newSlice}
}

func (s *Slice) IterativeSumOfSlice() int {
	var sum int
	for _, element := range s.NewSlice {
		sum += element
	}
	return sum
}

func (s *Slice) ConditionalSumOfSlice() int {
	var i = 0
	var sum int
	for i < len(s.NewSlice) {
		sum += s.NewSlice[i]
		i++
	}
	return sum
}

func (s *Slice) RecursiveSumOfSlice(length int) int {

	if length == 0 {
		return 0
	}
	return s.RecursiveSumOfSlice(length-1) + s.NewSlice[length-1]
}
