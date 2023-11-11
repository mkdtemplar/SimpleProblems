package main

import (
	"fmt"
)

func longestString(s []string) int {
	var maxlength = len(s[0])
	for _, v := range s {
		if len(v) > maxlength {
			maxlength = len(v)
		}
	}
	return maxlength
}

func RectangularFrame(s []string) string {
	var str string
	maxlength := longestString(s) + 4
	for i := 0; i < maxlength; i++ {
		str += "*"
	}
	str += "\n* "
	for _, v := range s {
		for _, c := range v {
			str += string(c)
		}
		for j := 0; j < maxlength - len(v) - 4; j++ {
			str += " "
		}
		str += " *\n"
		str += "* "
	}

	for i := 0; i < maxlength - 1; i++ {
		str += "*"
	}
	

	return str
}

func main() {
	input := []string{"Hello", "World", "in", "a", "frame"}
	fmt.Println(RectangularFrame(input))

}
