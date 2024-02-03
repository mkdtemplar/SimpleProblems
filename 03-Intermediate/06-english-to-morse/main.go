package main

import (
	"fmt"
	"strings"

	"github.com/mkdtemplar/SimpleProblems/Intermediate/06-english-to-morse/pkg"
)

var morseCode = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	" ": "/",
}

func convertToMorseCode(text string) string {
	builder := strings.Builder{}

	for _, s := range text {
		builder.WriteString(morseCode[string(s)])
	}

	return builder.String()
}

func main() {

	text := pkg.ProcessTextFromFile("text.txt")
	fmt.Println(text)

	morseCodeConvert := convertToMorseCode(text)
	fmt.Println("Blank space separator /")

	fmt.Println(morseCodeConvert)
}
