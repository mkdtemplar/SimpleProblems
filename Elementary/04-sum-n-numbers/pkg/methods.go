package pkg

func PrintSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func PrintSumRecursively(n int) int {

	if n == 0 {
		return 0
	}

	return PrintSumRecursively(n-1) + n
}
