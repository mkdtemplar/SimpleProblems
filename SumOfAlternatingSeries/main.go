package main

import (
	"fmt"
	"math"
)

func main() {

	var sum int

	for k := 0; k <= 1000000; k++ {
		sum += (int(math.Pow(-1, float64(k+1)))) / (2*k - 1)
	}

	sum *= 4
	fmt.Printf("Sum is %d\n", sum)
}
