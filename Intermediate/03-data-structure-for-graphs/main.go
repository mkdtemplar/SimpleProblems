package main

import (
	"github.com/mkdtemplar/SimpleProblems/Intermediate/03-data-structure-for-graphs/pkg"
)

func main() {
	test := pkg.Graph[int]{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(1, 2)
	test.AddEdge(1, 3)
	test.AddEdge(2, 4)

	test.Print()
}
