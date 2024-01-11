package pkg

import "fmt"

type GreetMessage struct {
	greetMessage string
}

func NewGreetMessage(greetMessage string) IGreetMessage {
	return &GreetMessage{greetMessage: greetMessage}
}

type IGreetMessage interface {
	PrintGreetMessage()
	PrintGreetMessage1() string
}

func (g *GreetMessage) PrintGreetMessage() {
	fmt.Println(g.greetMessage)
}

func (g *GreetMessage) PrintGreetMessage1() string {
	return g.greetMessage
}
