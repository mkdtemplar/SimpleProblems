package main

import (
	"fmt"
	"github.com/mkdtemplar/SimpleProblems/Elementary/02-Greetings/pkg"
)

func main() {
	var name string
	fmt.Print("Enter Name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}

	greeting := pkg.NewGreetings(name)

	greeting.Greeting()
}
