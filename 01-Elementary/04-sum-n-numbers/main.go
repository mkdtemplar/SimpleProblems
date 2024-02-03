package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/Elementary/04-sum-n-numbers/pkg"
)

func main() {
	var n int
	fmt.Print("Enter upper bond to calculate sum: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	fmt.Println(pkg.PrintSumRecursively(n))
	fmt.Println(pkg.PrintSum(n))
}
