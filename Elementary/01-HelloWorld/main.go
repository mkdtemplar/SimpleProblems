package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mkdtemplar/SimpleProblems/Elementary/01-HelloWorld/pkg"
)

func main() {
	var greetMessage string

	fmt.Print("Enter greet message: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		greetMessage = scanner.Text()
	}
	greet := pkg.NewGreetMessage(greetMessage)

	greet.PrintGreetMessage()
}
