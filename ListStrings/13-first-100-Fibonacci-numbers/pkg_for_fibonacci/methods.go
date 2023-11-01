package pkgforfibonacci

type Fibonacci struct {
	fibonacciSlice []int64
}

type IFibonacci interface {
	CalculateFibonacciRecursiv(n int) int64
	CalculateFibonacciIterative(n int) []int64
}

func NewFibonacci(fibonacciSlice []int64) IFibonacci {
	return &Fibonacci{fibonacciSlice: fibonacciSlice}
}

// CalculateFibonacciIterative implements IFibonacci.
func (f *Fibonacci) CalculateFibonacciIterative(n int) []int64 {
	f.fibonacciSlice = make([]int64, n+1, n+2)
	if n < 2 {
		f.fibonacciSlice = f.fibonacciSlice[0:2]
	}
	f.fibonacciSlice[0] = 0
	f.fibonacciSlice[1] = 1
	for i := 2; i <= n; i++ {
		f.fibonacciSlice[i] = f.fibonacciSlice[i-1] + f.fibonacciSlice[i-2]
		f.fibonacciSlice = append(f.fibonacciSlice, f.fibonacciSlice[i])
	}
	return f.fibonacciSlice
}

// CalculateFibonacciRecursiv implements IFibonacci.
func (f *Fibonacci) CalculateFibonacciRecursiv(n int) int64 {
	if n <= 1 {
		return int64(n)
	}

	return f.CalculateFibonacciRecursiv(n-1) + f.CalculateFibonacciRecursiv(n-2)
}
