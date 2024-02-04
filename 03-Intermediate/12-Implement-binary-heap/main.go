package main

import "fmt"

type MaxHeap struct {
	slice    []int
	heapSize int
}

func (h *MaxHeap) heapify(length, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < length && h.slice[left] > h.slice[largest] {
		largest = left
	}

	if right < length && h.slice[right] > h.slice[largest] {
		largest = right
	}

	if largest != i {
		h.slice[i], h.slice[largest] = h.slice[largest], h.slice[i]
		h.heapify(length, largest)
	}
}

func (h *MaxHeap) sort() {
	length := len(h.slice)

	for i := length/2 - 1; i >= 0; i-- {
		h.heapify(length, i)
	}

	for i := length - 1; i >= 0; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapify(i, 0)
	}
}

func main() {
	inputArray := []int{5, 1, 7, 9, 0, 32, 34, 56, 8}
	maxHeap := MaxHeap{
		slice:    inputArray,
		heapSize: len(inputArray),
	}

	maxHeap.sort()

	for _, e := range inputArray {
		fmt.Println(e)
	}
}
