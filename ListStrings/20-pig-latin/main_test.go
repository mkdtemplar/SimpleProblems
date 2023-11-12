package main

import "testing"

func Test_pigLatin(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Pig latin",
			args: args{str: "The quick brown fox"},
			want: "Hetay Uickqay Rownbay Oxfay ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pigLatin(tt.args.str); got != tt.want {
				t.Errorf("pigLatin() = %v, want %v", got, tt.want)
			}
		})
	}
}
