package main

import (
	"reflect"
	"testing"
)

func equalSlices(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	} else {
		for i, element := range a {
			if element != b[i] {
				return false
			}
		}
	}
	return true
}

func TestReverseList(t *testing.T) {
	input := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	reverseListResult := NewSlice(input)

	if reflect.DeepEqual(result, reverseListResult.reverseList()) == false {
		t.Errorf("Result wrong want %v, but got %v", result, reverseListResult)
	}
}

func TestReverseList2(t *testing.T) {
	input := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	reverseListResult := NewSlice(input)

	if equalSlices(reverseListResult.reverseList(), result) == false {
		t.Errorf("Result wrong want %v, but got %v", result, reverseListResult.reverseList())
	}
}
