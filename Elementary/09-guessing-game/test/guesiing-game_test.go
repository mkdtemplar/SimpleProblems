package test

import (
	"testing"

	"github.com/mkdtemplar/SimpleProblems/Elementary/09-guessing-game/pkg"
)

func TestGuessNumber(t *testing.T) {
	game := pkg.NewGame(77)

	got := game.GuessNumber(77)

	if true != got {
		t.Errorf("Test failed want %v, got %v", true, got)
	}
}
