package removeElement

import (
	"testing"
)

func Test_removeElement(t *testing.T) {
	type check func(*testing.T, int, []int)
	checks := func(cs ...check) []check { return cs }

	// Define checks
	resultEquals := func(expected int) check {
		return func(t *testing.T, actual int, _ []int) {
			if actual != expected {
				t.Errorf("Expected result: %d, actual: %d", expected, actual)
			}
		}
	}

	beginsWith := func(prefix []int) check {
		return func(t *testing.T, _ int, actual []int) {
			for i, _ := range prefix {
				if prefix[i] != actual[i] {
					t.Errorf("Result does not begin with expected prefix: \n"+
						"expected prefix: %v\n"+
						"actual  : %v",
						prefix, actual)
				}
			}
		}
	}

	type args struct {
		nums []int
		val  int
	}

	tests := []struct {
		name   string
		args   args
		checks []check
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			checks: checks(
				resultEquals(2),
				beginsWith([]int{2, 2}),
			),
		},
		{
			name: "example 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			checks: checks(
				resultEquals(5),
				beginsWith([]int{0, 1, 3, 0, 4}),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := removeElement(tt.args.nums, tt.args.val)
			for _, ch := range tt.checks {
				ch(t, out, tt.args.nums)
			}
		})
	}
}
