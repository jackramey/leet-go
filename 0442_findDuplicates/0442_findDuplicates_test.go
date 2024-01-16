package findDuplicates

import (
	"reflect"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			want: []int{2, 3},
		},
		{
			args: args{
				nums: []int{1, 1, 2},
			},
			want: []int{1},
		},
		{
			args: args{
				nums: []int{1},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDuplicatesInPlace(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			},
			want: []int{2, 3},
		},
		{
			args: args{
				nums: []int{1, 1, 2},
			},
			want: []int{1},
		},
		{
			args: args{
				nums: []int{1},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicatesInPlace(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicatesInPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}
