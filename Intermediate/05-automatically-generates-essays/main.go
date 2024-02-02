package main

import (
	"strings"

	"github.com/mkdtemplar/SimpleProblems/Intermediate/03-data-structure-for-graphs/pkg"
	processText "github.com/mkdtemplar/SimpleProblems/Intermediate/06-english-to-morse/pkg"
)

func main() {
	text := processText.ProcessTextFromFile("essay.txt")

	str := strings.Split(text, " ")

	g := pkg.Graph[string]{}

	for i := 0; i < len(str); i++ {
		g.AddVertex(str[i])
	}

	for i := 0; i < len(str)-1; i++ {
		g.AddEdgeMultiGraph(str[i], str[i+1])
	}

	g.Print()

}
