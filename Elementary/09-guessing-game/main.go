package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mkdtemplar/SimpleProblems/Elementary/09-guessing-game/pkg"
)

func main() {
	low, high := 1, 100
	rand.New(rand.NewSource(time.Now().UnixNano()))
	guessingNumber := rand.Intn(high-low) + low
	fmt.Println(guessingNumber)

	game := pkg.NewGame(guessingNumber)

	game.Play()

}
