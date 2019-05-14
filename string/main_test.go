package string

import (
	"strconv"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type expect struct {
		Input string
		Exp   string
	}
	ss := []expect{
		{
			Input: "babad",
			Exp:   "bab",
		},
		{
			Input: "cbbd",
			Exp:   "bb",
		},
		{
			Input: "a",
			Exp:   "a",
		},
		{Input: "ccc", Exp: "ccc"},
	}
	for _, s := range ss {
		//m := "123"
		//println(m[0:3])
		got := longestPalindrome(s.Input)
		if got != s.Exp {
			t.Errorf("Input %v, got: %v, Exp: %v", s.Input, got, s.Exp)
		}
	}

}

func Test_m(t *testing.T) {
	type args struct {
		s string
		L int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				"01234567890123456789",
				1,
			},
			"01234567890123456789",
		},
		{
			"2",
			args{
				"01234567890123456789",
				2,
			},
			"02468024681357913579",
		},
		{
			"test1",
			args{
				"01234567890123456789",
				3,
			},
			"04826135791357926048",
		},
		{
			"test1",
			args{
				"01234567890123456789",
				4,
			},
			"06281571379248046395",
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.args.L), func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.L); got != tt.want {
				t.Errorf("m() = %v, want %v", got, tt.want)
			}
		})
	}
}
