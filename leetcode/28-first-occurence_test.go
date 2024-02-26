package main

import (
	"strings"
	"testing"
)

//func strStr(haystack string, needle string) int {
//    for i:= 0; i <= len(haystack) - len(needle); i++ {
//        if haystack[i:len(needle)+i] == needle {
//            return i
//        }
//    }
//    return -1
//}
func firstOcc(haystack string, needle string) int {

	if strings.HasPrefix(haystack, needle) {
		return 0
	}

	arr := strings.Split(haystack, needle)
	if len(arr) == 0 || arr[0] == haystack {
		return -1
	}

	return len(arr[0])

}

func TestFirstOcc(t *testing.T) {
	var tests = []struct {
		input1 string
		input2 string
		output int
	}{
		{"sadbutsad", "sad", 0},
	}

	for _, tt := range tests {
		res := firstOcc(tt.input1, tt.input2)
		if res != tt.output {
			t.Errorf("%d not valid %d is expected", res, tt.output)
		}
	}
}
