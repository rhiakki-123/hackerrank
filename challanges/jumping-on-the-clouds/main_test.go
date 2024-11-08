package main

import "testing"

func TestJumpingOnClouds(t *testing.T) {
	tests := []struct {
		c        []int32
		expected int32
	}{
		{[]int32{0, 1, 0, 0, 0, 1, 0}, 3},
		{[]int32{0, 0, 0, 0, 1, 0}, 3},
		{[]int32{0, 0, 1, 0, 0, 1, 0}, 4},
		{[]int32{0, 0, 0, 1, 0, 0}, 3},
	}

	for _, test := range tests {
		result := jumpingOnClouds(test.c)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.c, test.expected, result)
		}
	}
}
