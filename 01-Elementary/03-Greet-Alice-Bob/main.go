package main

import (
	"fmt"

	"github.com/mkdtemplar/SimpleProblems/Elementary/03-Greet-Alice-Bob/pkg"
)

func main() {
	names := pkg.NewNames([]string{"Alice", "Bob"})

	fmt.Println(names.PrintNames())
}
