package pkg

import "math"

func SumOfAlternatingSeries() int {
	var sum int

	for k := 0; k <= 1000000; k++ {
		sum += (int(math.Pow(-1, float64(k+1)))) / (2*k - 1)
	}

	sum *= 4
	return sum
}
