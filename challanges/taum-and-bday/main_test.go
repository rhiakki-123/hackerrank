package main

import "testing"

func TestTaumBday(t *testing.T) {
	tests := []struct {
		b        int32
		w        int32
		bc       int32
		wc       int32
		z        int32
		expected int64
	}{
		{10, 10, 1, 1, 1, 20},
		{5, 9, 2, 3, 4, 37},
		{3, 6, 9, 1, 1, 12},
		{7, 7, 4, 2, 1, 35},
		{3, 3, 1, 9, 2, 12},
	}

	for _, test := range tests {
		result := taumBday(test.b, test.w, test.bc, test.wc, test.z)
		if result != test.expected {
			t.Errorf("For input b=%d, w=%d, bc=%d, wc=%d, z=%d, expected %d but got %d", test.b, test.w, test.bc, test.wc, test.z, test.expected, result)
		}
	}
}
