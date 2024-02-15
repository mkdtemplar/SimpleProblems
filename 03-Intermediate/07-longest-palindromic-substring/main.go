package main

import (
	"fmt"
)

// Function to check if a string is a palindrome
func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func maxLengthPalindrome(s []string) string {
	var maxLength = len(s[0])
	index := 0

	for i, e := range s {
		if len(e) > maxLength {
			maxLength = len(e)
			index = i
		}
	}
	return s[index]
}

// Function to find the longest palindromic substring of a given string
func longestPalindrome(s string) string {
	var palindromes []string

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i:j+1]) && len(s[i:j+1]) > 1 {
				palindromes = append(palindromes, s[i:j+1])
			}
		}
	}

	// Return the longest palindromic substring
	return maxLengthPalindrome(palindromes)
}

func main() {
	fmt.Println(longestPalindrome("asdfasdf1234321asd32"))

}
