package main

import (
	"fmt"

	methodsforlistofdigits "github.com/mkdtemplar/SimpleProblems/ListStrings/14-list-of-digits/methods-for-list-of-digits"
)

func main() {
	var n int

	fmt.Print("Enter integer:")

	_, err := fmt.Scan(&n)

	if err != nil {
		return
	}

	newPkg := methodsforlistofdigits.NewListOfDigits(n)

	fmt.Println(newPkg.PrintListOfDigits())

}
