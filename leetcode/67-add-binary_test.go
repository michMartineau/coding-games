package main

import "testing"

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	acc := '0'
	str := ""

	for true {
		if i < 0 && j < 0 {
			break
		}

		va := '0'
		if i >= 0 {
			va = rune(a[i])
		}

		vb := '0'
		if j >= 0 {
			vb = rune(b[j])
		}
		i--
		j--

		if acc == '0' && va == '0' && vb == '0' {
			str = "0" + str
			acc = '0'
		} else if acc == '0' && va == '1' && vb == '1' {
			str = "0" + str
			acc = '1'
		} else if acc == '0' && ((va == '0' && vb == '1') || (va == '1' && vb == '0')) {
			str = "1" + str
			acc = '0'
		} else if acc == '1' && va == '0' && vb == '0' {
			str = "1" + str
			acc = '0'
		} else if acc == '1' && va == '1' && vb == '1' {
			str = "1" + str
			acc = '1'
		} else if acc == '1' && ((va == '0' && vb == '1') || (va == '1' && vb == '0')) {
			str = "0" + str
			acc = '1'
		}
	}
	if acc == '1' {
		str = string(acc) + str
	}

	return str
}

func TestAddBinary(t *testing.T) {
	var tests = []struct {
		input1 string
		input2 string
		output string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, tt := range tests {
		resp := addBinary(tt.input1, tt.input2)
		if resp != tt.output {
			t.Errorf("for %s not valid %s is expected", resp, tt.output)
		}
	}
}
