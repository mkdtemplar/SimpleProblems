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
}

func (g *GreetMessage) PrintGreetMessage() {
	fmt.Println(g.greetMessage)
}
