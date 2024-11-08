package main

import "testing"

func TestEncryption(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"have a nice day", "hae and via ecy"},
		{"feed the dog", "fto ehg ee dd"},
		{"chill out", "clu hlt io"},
		{"if man was meant to stay on the ground god would have given us roots", "imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn sseoau"},
		{"", ""},
	}

	for _, test := range tests {
		result := encryption(test.input)
		if result != test.expected {
			t.Errorf("For input %q, expected %q but got %q", test.input, test.expected, result)
		}
	}
}
