package main

import "testing"

func TestRectangularFrame(t *testing.T) {
	result := "*********\n* Hello *\n* World *\n* in    *\n* a     *\n* frame *\n*********"

	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "RectangularFrame",
			args: args{s: []string{"Hello", "World", "in", "a", "frame"}},
			want: result},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RectangularFrame(tt.args.s); got != tt.want {
				t.Errorf("RectangularFrame() = %v, want %v", got, tt.want)
			}
		})
	}
}
