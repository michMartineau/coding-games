package main

import (
	"testing"
)

func longestCommonPrefix(strs []string) string {
	idx := 0
	res := ""

main:
	for {
		var tmp = ' '
		for idxL, val := range strs {
			if len(val) == idx {
				break main
			}
			if idxL == 0 {
				tmp = rune(val[idx])
				continue
			}
			if tmp != rune(val[idx]) {
				break main
			}
		}
		if tmp != ' ' {
			res += string(tmp)
		}
		idx++
	}

	return res
}
func TestLongestCommonPrefix(t *testing.T) {
	var tests = []struct {
		input1 []string
		output string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
	}

	for _, tt := range tests {
		res := longestCommonPrefix(tt.input1)
		if res != tt.output {
			t.Errorf("%s not valid %s is expected", res, tt.output)
		}
	}
}
