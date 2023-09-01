package rotateArray

import (
	"reflect"
	"testing"

	"github.com/jackramey/leet-go/0189_rotateArray/march"
	queueing "github.com/jackramey/leet-go/0189_rotateArray/queue"
	"github.com/jackramey/leet-go/0189_rotateArray/sidecar"
)

func Test_all(t *testing.T) {
	type check func(*testing.T, []int)
	checks := func(cs ...check) []check { return cs }

	// Define checks
	equal := func(expected []int) check {
		return func(t *testing.T, actual []int) {
			if !reflect.DeepEqual(expected, actual) {
				t.Errorf("expected: %v, actual %v", expected, actual)
			}
		}
	}

	algorithms := []struct {
		name string
		alg  func([]int, int)
	}{
		{
			name: "march",
			alg:  march.Rotate,
		},
		{
			name: "queue",
			alg:  queueing.Rotate,
		},
		{
			name: "sidecar",
			alg:  sidecar.Rotate,
		},
	}

	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name   string
		args   args
		checks []check
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			checks: checks(
				equal([]int{5, 6, 7, 1, 2, 3, 4}),
			),
		},
		{
			name: "example 2",
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			checks: checks(
				equal([]int{3, 99, -1, -100}),
			),
		},
		{
			name: "example 3",
			args: args{
				nums: []int{-1},
				k:    1,
			},
			checks: checks(
				equal([]int{-1}),
			),
		},
		{
			name: "example 4",
			args: args{
				nums: []int{1, 2},
				k:    3,
			},
			checks: checks(
				equal([]int{2, 1}),
			),
		},
		{
			name: "example 5",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    2,
			},
			checks: checks(
				equal([]int{5, 6, 1, 2, 3, 4}),
			),
		},
		{
			name: "example 6",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    4,
			},
			checks: checks(
				equal([]int{3, 4, 5, 6, 1, 2}),
			),
		},
		{
			name: "example 7",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    3,
			},
			checks: checks(
				equal([]int{4, 5, 6, 1, 2, 3}),
			),
		},
		{
			name: "example 8",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
				k:    6,
			},
			checks: checks(
				equal([]int{11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			),
		},
	}

	for _, alg := range algorithms {
		for _, tt := range tests {
			t.Run(alg.name+": "+tt.name, func(t *testing.T) {
				nums := append([]int(nil), tt.args.nums...)
				alg.alg(nums, tt.args.k)
				for _, ch := range tt.checks {
					ch(t, nums)
				}
			})
		}
	}
}
