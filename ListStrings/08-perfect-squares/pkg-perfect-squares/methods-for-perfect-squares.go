package pkg_perfect_squares

import "math"

type SlicePerfectSquares struct {
	slicePerfectSquares []int
}

func (s *SlicePerfectSquares) PerfectSquare(f func(n int) int) []int {
	var perfectSquares []int

	for _, element := range s.slicePerfectSquares {
		perfectSquares = append(perfectSquares, f(element))
	}

	return perfectSquares
}

func NewSlicePerfectSquares(slicePerfectSquares []int) ISlicePerfectSquares {
	return &SlicePerfectSquares{slicePerfectSquares: slicePerfectSquares}
}

type ISlicePerfectSquares interface {
	PerfectSquare(f func(n int) int) []int
}

func CalculatePerfectSquare(n int) int {
	return int(math.Pow(float64(n), 2))
}

func CreateSlice() []int {
	var slice = make([]int, 0)

	for i := 1; i <= 20; i++ {
		slice = append(slice, i)
	}

	return slice
}
