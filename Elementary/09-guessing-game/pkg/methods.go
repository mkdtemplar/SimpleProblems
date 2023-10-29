package pkg

import (
	"fmt"
)

type GuessingGame struct {
	Number int
}

type IGuessingGame interface {
	GuessNumber(number int) bool
	Play()
}

func NewGame(number int) IGuessingGame {
	return &GuessingGame{Number: number}
}

func (g *GuessingGame) GuessNumber(number int) bool {
	if number > g.Number {
		fmt.Printf("Entered number %d to highn", number)
	}
	if number < g.Number {
		fmt.Printf("Entered number %d to low\n", number)
	}
	if number == g.Number {
		fmt.Println("You win")
		return true
	}

	return false
}

func (g *GuessingGame) Play() {
	var attempts = 10
	var number int
	fmt.Print("Enter number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}
	for attempts > 0 && !g.GuessNumber(number) {
		attempts--
		fmt.Printf("You have %d attempts left\n", attempts)
		fmt.Print("Enter number: ")
		_, err := fmt.Scan(&number)
		if err != nil {
			return
		}
	}
}
