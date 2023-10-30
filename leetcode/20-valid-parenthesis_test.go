package main

import "testing"

var chars = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	var tmp []rune
	for _, r := range s {

		v, ok := chars[r]
		if !ok {
			tmp = append(tmp, r)
			continue
		}

		l := len(tmp)
		if l == 0 && ok {
			return false
		}

		if l >= 1 && v != tmp[l-1] {
			return false
		}

		tmp = tmp[:l-1]
	}

	return len(tmp) == 0
}

func TestIsValid(t *testing.T) {
	var tests = []struct {
		input  string
		output bool
	}{
		{"", false},
		{"[]", true},
		{"[}]", false},
	}

	for _, tt := range tests {
		valid := isValid(tt.input)
		if valid != tt.output {
			t.Errorf("for %s not valid %t is expected", tt.input, tt.output)
		}

	}

}
