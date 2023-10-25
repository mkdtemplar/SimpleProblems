package pkg

import (
	"fmt"
)

type Greetings struct {
	greetName string
}

type IGreetings interface {
	Greeting()
}

func NewGreetings(greetName string) IGreetings {
	return &Greetings{greetName: greetName}
}

func (g *Greetings) Greeting() {
	fmt.Println(fmt.Sprintf("Hello %s", g.greetName))
}
