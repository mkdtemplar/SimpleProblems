package main

import (
	"fmt"
	"math"
)

func main() {
	a := "AGGTAB"
	b := "GXTXAYB"
	m := len(a)
	n := len(b)
	fmt.Println(lcs(a, b, m, n))
}

func lcs(a string, b string, m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if a[m-1] == b[n-1] {
		return 1 + lcs(a, b, m-1, n-1)
	} else {
		return int(math.Max(float64(lcs(a, b, m, n-1)), float64(lcs(a, b, m-1, n))))
	}
}
