package main

import (
	"fmt"
	"testing"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r) - 1; i < len(r)/2; i, j = i+1, j-1{
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Test(t *testing.T) {
	cases := []struct{
		in, want string
	}{
		{"Hello, world", "dlrow, olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c:= range cases{
		got := Reverse(c.in)
		if got != c.want{
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
func main() {
	fmt.Print("Hihi")
	Test("Hello, world")
}
