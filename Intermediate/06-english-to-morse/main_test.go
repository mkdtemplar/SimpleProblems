package main

import (
	"testing"
)

func Test_processTextFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "processTextFromFile",
			args: args{filename: "text.txt"},
			want: "VICTORIOUS WARRIORS WIN FIRST AND THEN GO TO WAR, WHILE DEFEATED WARRIORS GO TO WAR FIRST AND THEN SEEK TO WIN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processTextFromFile(tt.args.filename); got != tt.want {
				t.Errorf("processTextFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertToMorseCode(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "convertToMorseCode",
			args: args{text: "VICTORIOUS WARRIORS WIN FIRST AND THEN GO TO WAR, WHILE DEFEATED WARRIORS GO TO WAR FIRST AND THEN SEEK TO WIN"},
			want: "...-..-.-.----.-...---..-.../.--.-.-..-...---.-..../.--..-./..-....-....-/.--.-../-.-./--.---/----/.--.-.-./.--...-.../-.....-...--.-../.--.-.-..-...---.-..../--.---/----/.--.-.-./..-....-....-/.--.-../-.-./.....-.-/----/.--..-.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToMorseCode(tt.args.text); got != tt.want {
				t.Errorf("convertToMorseCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
