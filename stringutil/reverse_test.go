package stringutil

import "testing"

func TestReverse(t *testing.T) {
	// define a simple struct that has two strings, in and want
	// create a slice of that struct with default values
	cases := []struct {
		in, want string
	}{
		{"Hello World", "dlroW olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, v := range cases {
		got := Reverse(v.in)
		if got != v.want {
			t.Errorf("Reverse(%q) == %q, want %q", v.in, got, v.want)
		}
	}
}
