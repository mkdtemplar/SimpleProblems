package pkg

import (
	"fmt"
)

const (
	n = 12
)

func CalculateProduct(num int, i int) int {
	if i > 10 {
		return 0
	}
	fmt.Printf("%d X %d = %d\n", num, i, num*i)

	return CalculateProduct(num, i+1)
}

func CalculateTable() {
	for i := 1; i <= n; i++ {
		CalculateProduct(i, 1)
		fmt.Println("____________")
	}
}
