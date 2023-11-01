package main

import (
	"fmt"

	pkgforfibonacci "github.com/mkdtemplar/SimpleProblems/ListStrings/13-first-100-Fibonacci-numbers/pkg_for_fibonacci"
)

func main() {
	var input []int64

	newFib := pkgforfibonacci.NewFibonacci(input)

	fmt.Println(newFib.CalculateFibonacciIterative(100))

	fmt.Println(newFib.CalculateFibonacciRecursiv(10))
}
