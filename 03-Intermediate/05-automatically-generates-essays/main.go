package main

import (
	"strings"

	processText "github.com/mkdtemplar/SimpleProblems/03-Intermediate/06-english-to-morse/pkg"
	"github.com/mkdtemplar/SimpleProblems/Intermediate/03-data-structure-for-graphs/pkg"
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
