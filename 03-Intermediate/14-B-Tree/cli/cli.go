package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mkdtemplar/SimpleProblems/03-Intermediate/14-B-Tree/b-tree"
)

type CLI struct {
	scanner *bufio.Scanner
	tree    *b_tree.BTree
}

func NewCLI(s *bufio.Scanner, b *b_tree.BTree) *CLI {
	return &CLI{s, b}
}

func (c *CLI) Start() {
	c.printHelp()
	c.printPrompt()
	for {
		if c.scanner.Scan() {
			c.processInput(c.scanner.Text())
		}
	}
}

func (c *CLI) printHelp() {
	fmt.Println(`
B-Tree CLI

Available Commands:
  SET <key> <val> Insert a key-value pair into the B-Tree
  DEL <key>       Remove a key-value pair from the B-Tree
  GET <key>       Retrieve the value for key from the B-Tree
  EXIT            Terminate this session
`)
}

func (c *CLI) printPrompt() {
	fmt.Print("> ")
}

func (c *CLI) processInput(line string) {
	fields := strings.Fields(line)

	if len(fields) < 1 {
		return
	}
	command := strings.ToLower(fields[0])

	switch command {
	default:
		fmt.Printf("Unknown command \"%s\"\n", command)
	case "set":
		c.processSetCommand(fields[1:])
	case "del":
		c.processDeleteCommand(fields[1:])
	case "get":
		c.processGetCommand(fields[1:])
	case "exit":
		os.Exit(0)
	}
	c.printPrompt()
}

func (c *CLI) processSetCommand(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: SET <key> <value>")
		return
	}
	c.tree.Insert([]byte(args[0]), []byte(args[1]))
	fmt.Println(c.tree)
}

func (c *CLI) processDeleteCommand(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: DEL <key>")
		return
	}
	res := c.tree.Delete([]byte(args[0]))

	if !res {
		fmt.Println("Key not found.")
		return
	}
	fmt.Println(c.tree)
}

func (c *CLI) processGetCommand(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: GET <key>")
		return
	}
	val, err := c.tree.Find([]byte(args[0]))

	if err != nil {
		fmt.Println("Key not found.")
		return
	}
	fmt.Println(string(val))
}
