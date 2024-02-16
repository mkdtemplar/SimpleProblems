package main

import "fmt"

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

func maxLengthPalindrome(s []string, strChan chan string) chan string {

	var maxLength = len(s[0])
	index := 0
	go func() {
		for i, e := range s {
			if len(e) > maxLength {
				maxLength = len(e)
				index = i
			}
		}
		strChan <- s[index]
	}()

	return strChan

}

// Function to find the longest palindromic substring of a given string
func longestPalindrome(s string, strChan chan []string) chan []string {
	var palindromes []string

	go func() {
		for i := 0; i < len(s); i++ {
			for j := i; j < len(s); j++ {
				if isPalindrome(s[i:j+1]) && len(s[i:j+1]) > 1 {
					palindromes = append(palindromes, s[i:j+1])
				}
			}
		}
		strChan <- palindromes
	}()

	// Return the longest palindromic substring
	return strChan

}

func main() {

	strChan := make(chan []string)
	maxStringChan := make(chan string)

	palindromes := <-longestPalindrome("asdfasdf1234321asd32", strChan)

	maxSubString := <-maxLengthPalindrome(palindromes, maxStringChan)
	fmt.Println(maxSubString)

}
