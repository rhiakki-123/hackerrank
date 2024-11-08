package main

import (
	"testing"
)

func TestRepeatedString(t *testing.T) {
	tests := []struct {
		s        string
		n        int64
		expected int64
	}{
		{"abcac", 10, 4},
		{"a", 1000000000000, 1000000000000},
		{"abc", 7, 3},
		{"abca", 10, 5},
		{"", 10, 0},
	}

	for _, test := range tests {
		result := repeatedString(test.s, test.n)
		if result != test.expected {
			t.Errorf("For input s=%s and n=%d, expected %d but got %d", test.s, test.n, test.expected, result)
		}
	}
}
