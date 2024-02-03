package pkg

import (
	"fmt"
)

type Greetings struct {
	greetName string
}

// Greeting1 implements IGreetings.

type IGreetings interface {
	Greeting()
	Greeting1() string
}

func NewGreetings(greetName string) IGreetings {
	return &Greetings{greetName: greetName}
}

func (g *Greetings) Greeting() {
	fmt.Println(fmt.Sprintf("Hello %s", g.greetName))
}

func (g *Greetings) Greeting1() string {
	return fmt.Sprintf("Hello %s\n", g.greetName)
}
