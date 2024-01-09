package main

import (
	"reflect"
	"testing"
)

func Test_leapYearRule(t *testing.T) {
	type args struct {
		yearDuration float32
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Leap year generator",
			args: args{365.26},
			want: []int{2024, 2028, 2032, 2036, 2040},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leapYearRule(tt.args.yearDuration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leapYearRule() = %v, want %v", got, tt.want)
			}
		})
	}
}
