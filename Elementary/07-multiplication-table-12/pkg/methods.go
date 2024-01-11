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
		//CalculateProduct(i, 1)
		fmt.Println(CalculateProductIteratively(i))
		fmt.Println("____________")
	}
}

func CalculateProductIteratively(num int) []int {
	var productSlice []int
	for i := 1; i <= 10; i++ {
		productSlice = append(productSlice, num*i)
	}
	return productSlice
}

func CalculateTable2DSlice() [][]int {
	var tableSlice [][]int

	for i := 1; i <= n; i++ {
		row := CalculateProductIteratively(i)
		tableSlice = append(tableSlice, row)
	}

	return tableSlice
}

func PrintTable2D(table [][]int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			fmt.Print(fmt.Sprintf("%v X %v = %v\n ", table[i][0], j+1, table[i][j]))

		}
		fmt.Println("------------------")
	}
}
