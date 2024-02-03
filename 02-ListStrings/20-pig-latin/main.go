package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func rotateCharactersLeft(s []string) []string {
	for i := 0; i < 1; i++ {
		lastElement := s[0]
		for j := 0; j < len(s)-1; j++ {
			s[j] = s[j+1]
		}
		s[len(s)-1] = lastElement
	}

	return s
}

func pigLatin(str string) string {

	stringToSlice := strings.Split(str, " ")
	var rotatedString []string
	var resultString = ""

	for _, s := range stringToSlice {
		s1 := strings.Split(s, "")
		rotatedString = rotateCharactersLeft(s1)
		rotatedString = append(rotatedString, "ay")
		for _, s2 := range rotatedString {
			resultString += s2
		}
		resultString += " "

	}
	resultString = strings.ToLower(resultString)
	resultString = cases.Title(language.AmericanEnglish).String(resultString)
	return resultString

}

func main() {

	fmt.Println(pigLatin("The quick brown fox"))
}
