package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/Elementary/05-multiples/pkg"
)

func main() {

	var multipliesSlice []int

	multiplies := pkg.NewMultiplies(multipliesSlice)

	fmt.Println(multiplies.PrintMultiplies(100))
}
