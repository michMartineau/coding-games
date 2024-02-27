package main

import "testing"

func lengthOfLastWord(s string) int {
	l := len(s)
	count := 0
	started := false
	for i := l - 1; i >= 0; i-- {
		if rune(s[i]) != ' ' {
			count++
			started = true
			continue
		}
		if rune(s[i]) == ' ' && started {
			break
		}
	}
	return count
}

func TestLengthOfLastWord(t *testing.T) {
	var tests = []struct {
		input1 string
		output int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	for _, tt := range tests {
		resp := lengthOfLastWord(tt.input1)
		if resp != tt.output {
			t.Errorf("for %d not valid %d is expected", resp, tt.output)
		}
	}
}
