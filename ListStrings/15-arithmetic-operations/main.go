package main

import (
	"fmt"

	pkgforarithmetic "github.com/mkdtemplar/SimpleProblems/ListStrings/15-arithmetic-operations/pkg-for-arithmetic"
)

func main() {

	fmt.Println(pkgforarithmetic.StringToDigitsSlice("999999999999999999999"))

	newPkg := pkgforarithmetic.NewNumbers("999999999999999999999", "999999999999999999999")

	fmt.Println(newPkg.Addition())
}
