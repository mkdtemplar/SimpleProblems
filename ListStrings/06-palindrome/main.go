package main

import (
	"fmt"
	"strings"
)

type Palindrome struct {
	word string
}

func (p *Palindrome) reverseWord() string {
	wordToSlice := []rune(p.word)

	for i, j := 0, len(wordToSlice)-1; i < j; i, j = i+1, j-1 {
		wordToSlice[i], wordToSlice[j] = wordToSlice[j], wordToSlice[i]
	}

	return string(wordToSlice)
}

func (p *Palindrome) checkPalindrome() bool {

	if p.word != p.reverseWord() {
		return false
	}

	return true
}

func NewPalindrome(word string) IPalindrome {
	return &Palindrome{word: strings.ToLower(word)}
}

type IPalindrome interface {
	reverseWord() string
	checkPalindrome() bool
}

func main() {

	var word string

	fmt.Print("Enter Word: ")
	_, err := fmt.Scan(&word)
	if err != nil {
		return
	}

	palindrome := NewPalindrome(word)
	if palindrome.checkPalindrome() {
		fmt.Printf("Word %s is palindrome\n", word)
	} else {
		fmt.Printf("Word %s is not palindrome\n", word)
	}

}
