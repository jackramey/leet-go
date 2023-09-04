package isSubsequence

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				s: "abc",
				t: "ahbgdc",
			},
			want: true,
		},
		{
			args: args{
				s: "axc",
				t: "ahbgdc",
			},
			want: false,
		},
		{
			args: args{
				s: "acb",
				t: "ahbgdc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
