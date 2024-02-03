package main

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func TestBinarySearch(t *testing.T) {
	type args[T constraints.Ordered] struct {
		a []T
		x T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "Find element",
			args: args[int]{a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, x: 5},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
