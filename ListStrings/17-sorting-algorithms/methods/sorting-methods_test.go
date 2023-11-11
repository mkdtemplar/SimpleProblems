package methods

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestSelectionSort(t *testing.T) {
	type args[T constraints.Ordered] struct {
		s []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "integer slice",
			args: args[int]{s: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectionSort(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	type args[T constraints.Ordered] struct {
		s []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "String slice",
			args: args[string]{s: []string{"h", "g", "f", "e", "d", "c", "b", "a"}},
			want: []string{"a", "b", "c", "d", "e", "f", "g", "h"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertionSort(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args[T constraints.Ordered] struct {
		left  []T
		right []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "integer slice",
			args: args[int]{left: []int{1, 2, 3, 4, 5}, right: []int{6, 7, 8, 9, 10}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args[T constraints.Ordered] struct {
		s []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "integer slice",
			args: args[int]{s: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuicksort(t *testing.T) {
	type args[T constraints.Ordered] struct {
		s []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "integer slice",
			args: args[int]{s: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quicksort(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStoogeSort(t *testing.T) {
	input := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	type args[T constraints.Ordered] struct {
		s    []T
		low  int
		high int
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "integer slice",
			args: args[int]{s: input, low: 0, high: len(input) - 1},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StoogeSort(tt.args.s, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StoogeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
