package main

import (
	"reflect"
	"testing"
)

func TestTree_addNode(t *testing.T) {
	type fields struct {
		root *TreeNode
	}
	type args struct {
		root  *TreeNode
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TreeNode
	}{
		{
			name:   "Empty root",
			fields: fields{root: nil},
			args: args{
				value: 8,
			},
			want: &TreeNode{value: 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Tree{
				root: tt.fields.root,
			}
			if got := b.addNode(tt.args.root, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
