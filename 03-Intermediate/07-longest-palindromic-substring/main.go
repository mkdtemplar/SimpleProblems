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

// Function to find all palindromes
func longestPalindrome(s string, strChan chan string) chan string {
	var maxPalindrome string
	for i := 0; i < len(s); i++ {
		i := i
		go func() {
			for j := i; j < len(s); j++ {
				if isPalindrome(s[i:j+1]) && len(s[i:j+1]) > 1 {
					if len(maxPalindrome) < len(s[i:j+1]) {
						maxPalindrome = s[i : j+1]
					}
				}

			}
			strChan <- maxPalindrome
		}()
	}
	return strChan
}

func main() {
	strChan := make(chan string)

	palindromes := <-longestPalindrome("asdfasdf1234321asd32", strChan)
	fmt.Println(palindromes)

}
