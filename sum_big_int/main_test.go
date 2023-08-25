package main

import (
	"reflect"
	"testing"
)

func Test_sum(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name  string
		args  args
		want1 []int
		want2 []int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				arr1: []int{1, 2, 3},
				arr2: []int{4, 5, 6},
			},
			want1: []int{0, 5, 7, 9},
			want2: []int{5, 7, 9},
		},
		{
			name: "2",
			args: args{
				arr1: []int{5, 4, 4},
				arr2: []int{4, 5, 6},
			},
			want1: []int{1, 0, 0, 0},
			want2: []int{1, 0, 0, 0},
		},
		{
			name: "3",
			args: args{
				arr1: []int{5, 4, 4},
				arr2: []int{5, 6},
			},
			want1: []int{6, 0, 0},
			want2: []int{0, 6, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want1) && !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("sum() = %v, want %v or %v", got, tt.want1, tt.want2)
			}
		})
	}
}
