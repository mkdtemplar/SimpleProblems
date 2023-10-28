package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/Elementary/06-choose-operation/pkg"
)

func main() {

	var n int
	var choice int

	fmt.Print("Enter choice (1) for sum or (2) for product: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		return
	}

	fmt.Print("Enter upper bond: ")
	_, err = fmt.Scan(&n)

	if err != nil {
		return
	}

	switch choice {
	case 1:
		fmt.Printf("Sum of the numbers up to %d : %d", n, pkg.ComputeSum(n))
		return
	case 2:
		fmt.Printf("Product of the numbers up to %d : %d", n, pkg.ComputeProduct(n))
		return
	default:
		fmt.Println("Wrong choice")
		return
	}

}
