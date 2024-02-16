package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup

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

// Function to return longest palindrome substring
func getMaxLengthPalindrome(s []string) string {
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

// Function to find all palindromes
func longestPalindrome(s string, strChan chan []string, wg *sync.WaitGroup) chan []string {
	var palindromes []string
	// go routines inside for loop
	for i := 0; i < len(s); i++ {
		wg.Add(1)
		i := i
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for j := i; j < len(s); j++ {
				if isPalindrome(s[i:j+1]) && len(s[i:j+1]) > 1 {
					palindromes = append(palindromes, s[i:j+1])
				}
			}
			wg.Wait()
		}(wg)

	}
	strChan <- palindromes

	return strChan

}

func main() {

	strChan := make(chan []string)
	maxStringChan := make(chan string)

	palindromes := <-longestPalindrome("asdfasdf1234321asd32", strChan, wg)
	fmt.Println(palindromes)

	maxSubString := getMaxLengthPalindrome(palindromes)
	fmt.Println(maxSubString)

	defer close(strChan)
	defer close(maxStringChan)
}
