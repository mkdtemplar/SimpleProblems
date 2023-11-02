package main

import (
	"fmt"

	pkgforarithmetic "github.com/mkdtemplar/SimpleProblems/ListStrings/15-arithmetic-operations/pkg-for-arithmetic"
)

func main() {

	newPkg := pkgforarithmetic.NewNumbers(9999, 9999)

	fmt.Println(newPkg.Addition())
}
