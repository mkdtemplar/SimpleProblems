package pkg

func PrimeNumbers(num int) []int {
	var primesSlice []int
	for i := 2; i <= num; i++ {
		isPrime := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primesSlice = append(primesSlice, i)
		}
	}
	return primesSlice
}
