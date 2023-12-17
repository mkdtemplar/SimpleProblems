package main

import (
	"fmt"
	"log"
	"strconv"
)

// StringEval A function that evaluates a string expression and returns the result
func StringEval(expr string) int {
	sum := 0
	sign := '+'
	// Loop through each character in the expression
	for i := 0; i < len(expr); i++ {
		c := expr[i]
		// If the character is a digit, append it to a number string
		if c >= '0' && c <= '9' {
			num := string(c)
			// Keep appending digits until reaching a non-digit or the end of the expression
			for i+1 < len(expr) && expr[i+1] >= '0' && expr[i+1] <= '9' {
				num += string(expr[i+1])
				i++
			}
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			// Add or subtract the number to the sum based on the sign
			if sign == '+' {
				sum += n
			} else {
				sum -= n
			}
		} else {
			// If the character is not a digit, update the sign
			sign = int32(c)
		}
	}
	// Return the final sum
	return sum
}

// A function that generates all possible expressions and prints those that equal 100
func generate(digits string, fillers []string, expr string, index int) {
	// If the index reaches the end of the digits, evaluate the expression and print it if it equals 100
	if index == len(digits) {
		if StringEval(expr) == 100 {
			fmt.Printf("%s = 100\n", expr)
		}
		return
	}
	// Loop through each filler and append it to the expression along with the next digit
	for _, filler := range fillers {
		generate(digits, fillers, fmt.Sprintf("%s%s%s", expr, filler, string(digits[index])), index+1)
	}
}

func main() {
	// Define the digits and the fillers
	digits := "123456789"
	fillers := []string{"+", "-", ""}
	// Call the generate function with the first digit and index 1
	generate(digits, fillers, string(digits[0]), 1)

}
