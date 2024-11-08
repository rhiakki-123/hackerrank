package main

import "testing"

func TestTimeInWords(t *testing.T) {
	tests := []struct {
		h, m int32
		want string
	}{
		{5, 0, "five o' clock"},
		{5, 1, "one minute past five"},
		{5, 10, "ten minutes past five"},
		{5, 15, "quarter past five"},
		{5, 30, "half past five"},
		{5, 40, "twenty minutes to six"},
		{5, 45, "quarter to six"},
		{5, 47, "thirteen minutes to six"},
		{12, 0, "twelve o' clock"},
		{12, 59, "one minute to one"},
	}

	for _, tt := range tests {
		got := timeInWords(tt.h, tt.m)
		if got != tt.want {
			t.Errorf("timeInWords(%v, %v) = %v; want %v", tt.h, tt.m, got, tt.want)
		}
	}
}
