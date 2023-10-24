package main

import (
	"strings"
	"testing"
)

func TestReverseWord(t *testing.T) {
	input := strings.ToLower("ANA")
	output := NewPalindrome(input)

	if input != output.reverseWord() {
		t.Errorf("Test failed want %s got %s", input, output.reverseWord())
	}
}

func TestCheckPalindrome(t *testing.T) {
	input := strings.ToLower("ANA")
	output := NewPalindrome(input)

	if output.checkPalindrome() == false {
		t.Errorf("Test failed want %v, got %v", true, output.checkPalindrome())
	}
}
