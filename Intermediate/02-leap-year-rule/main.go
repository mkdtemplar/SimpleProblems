package main

import (
	"fmt"
	"math"
)

func leapYearRule(yearDuration float32) []int  {
	var currentYear = 2023
	var leapYear []int
	var difference = 0.0

	for difference < 1 && currentYear <= 2043 {
		fraction := yearDuration - float32(math.Trunc(float64(yearDuration)))
		difference += float64(fraction)

		if difference >= 0.5 {
			difference -= 1
			leapYear = append(leapYear, currentYear)
		}
		currentYear++
	}

	return leapYear
	
}

func main() {
	fmt.Println(leapYearRule(365.26))
	
}