package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func processTextFromFile(filename string) string {

	text, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return strings.ToUpper(string(text))
}

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

	text := processTextFromFile("text.txt")
	fmt.Println(text)

	morseCodeConvert := convertToMorseCode(text)
	fmt.Println("Blank space separator /")

	fmt.Println(morseCodeConvert)
}
