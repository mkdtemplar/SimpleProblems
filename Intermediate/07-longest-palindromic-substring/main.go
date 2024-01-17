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

// Function to find the longest palindromic substring of a given string
func longestPalindrome(s string) string {
	n := len(s)
	maxLength := 0
	start := 0
	end := 0
	// Consider all substrings of length from 1 to n
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// Check if the substring s[i:j+1] is a palindrome
			if isPalindrome(s[i : j+1]) {
				// Update the result if the substring is longer than the current longest palindrome
				if j-i+1 > maxLength {
					maxLength = j - i + 1
					start = i
					end = j
				}
			}
		}
	}
	// Return the longest palindromic substring
	return s[start : end+1]
}

func main() {
	fmt.Println(longestPalindrome("palindromeanavolimilovanacheck"))

}
