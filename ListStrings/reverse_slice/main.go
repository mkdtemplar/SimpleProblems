package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var reverseSlice []int

	for i := len(slice) - 1; i >= 0; i-- {
		reverseSlice = append(reverseSlice, slice[i])
	}

	fmt.Println(reverseSlice)
}
