package main

import (
	"reflect"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isPalindrome",
			args: args{s: "anavolimilovana"},
			want: true,
		},
		{
			name: "isPalindrome false",
			args: args{s: "Ivan Markovski"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s       string
		strChan chan []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "longestPalindrome",
			args: args{
				s:       "asdfasdf1234321asd32",
				strChan: make(chan []string),
			},
			want: []string{"1234321", "23432", "343"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := <-longestPalindrome(tt.args.s, tt.args.strChan); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxLengthPalindrome(t *testing.T) {
	type args struct {
		s       []string
		strChan chan string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "maxLengthPalindrome",
			args: args{
				s:       []string{"1234321", "23432", "343"},
				strChan: make(chan string),
			},
			want: "1234321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := <-maxLengthPalindrome(tt.args.s, tt.args.strChan); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxLengthPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
