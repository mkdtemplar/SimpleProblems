package main

import (
	"bufio"
	"os"

	"github.com/mkdtemplar/SimpleProblems/03-Intermediate/14-B-Tree/b-tree"
	"github.com/mkdtemplar/SimpleProblems/03-Intermediate/14-B-Tree/cli"
)

func main() {

	tree := b_tree.BTree{}
	scanner := bufio.NewScanner(os.Stdin)
	demo := cli.NewCLI(scanner, &tree)
	demo.Start()
}
