package pkg

import "fmt"

func CheckLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}

func PrintLeapYears() {
	for y := 2023; y <= 2043; y++ {
		if CheckLeapYear(y) {
			fmt.Println(y)
		}
	}
}
