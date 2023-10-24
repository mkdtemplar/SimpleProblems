package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8}
	maxElement := sliceA[0]

	for _, i := range sliceA {
		if i > maxElement {
			maxElement = i
		}
	}

	fmt.Println(maxElement)

}
