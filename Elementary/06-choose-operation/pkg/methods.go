package pkg

func ComputeSum(n int) int {
	var sum int

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func ComputeProduct(n int) int {
	var product = 1

	for i := 1; i <= n; i++ {
		product *= i
	}

	return product
}
