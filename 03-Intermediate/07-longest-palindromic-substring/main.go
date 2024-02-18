package main

import (
	"fmt"
	"runtime"
	"sync"
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
func longestPalindrome(s string, strChan chan string) string {
	var maxPalindrome string
	var wg sync.WaitGroup
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			j := j
			wg.Add(1)
			go func(str string) {
				defer wg.Done()
				if isPalindrome(s[i:j+1]) && len(s[i:j+1]) > 1 {
					if len(maxPalindrome) < len(s[i:j+1]) {
						maxPalindrome = s[i : j+1]
					}
				}
				strChan <- maxPalindrome
			}(s)
		}
	}
	fmt.Println("Total number of go routines created: ", runtime.NumGoroutine())
	go func(str chan string, wg *sync.WaitGroup) {
		wg.Wait()
		<-strChan
		defer close(strChan)
	}(strChan, &wg)
	result := <-strChan
	return result
}

func main() {
	strChan := make(chan string)
	palindromes := longestPalindrome("asdfasdf1234321asd32", strChan)

	fmt.Println(palindromes)

}
