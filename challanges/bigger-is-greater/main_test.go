package main

import "testing"

func TestBiggerIsGreater(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abdc", "acbd"},
		{"abcd", "abdc"},
		{"dcba", "no answer"},
		{"ab", "ba"},
		{"bb", "no answer"},
	}

	for _, test := range tests {
		result := biggerIsGreater(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %q but got %q", test.input, test.expected, result)
		}
	}
}
