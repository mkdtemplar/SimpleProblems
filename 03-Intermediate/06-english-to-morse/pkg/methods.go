package pkg

import (
	"log"
	"os"
	"strings"
)

func ProcessTextFromFile(filename string) string {

	text, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return strings.ToUpper(string(text))
}
