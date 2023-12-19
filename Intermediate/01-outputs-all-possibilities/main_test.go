package main

import (
	"log"
	"strconv"
	"testing"
)

func Test_eval(t *testing.T) {
	n, err := strconv.Atoi("123456789")
	if err != nil {
		log.Fatal(err)
	}
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Possibilities",
			args: args{expr: "123456789"},
			want: n,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfCombinations(tt.args.expr); got != tt.want {
				t.Errorf("sumOfCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generate(t *testing.T) {
	digits := "123456789"
	type args struct {
		digits  string
		fillers []string
		expr    string
		index   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "print combos",
			args: args{
				digits:  digits,
				fillers: []string{"+", "-", ""},
				expr:    string(digits[0]),
				index:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			combinations(tt.args.digits, tt.args.fillers, tt.args.expr, tt.args.index)
		})
	}
}
