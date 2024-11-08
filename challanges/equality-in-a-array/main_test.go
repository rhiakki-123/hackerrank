package main

import "testing"

func TestEqualizeArray(t *testing.T) {
	tests := []struct {
		arr      []int32
		expected int32
	}{
		{[]int32{3, 3, 2, 1, 3}, 2},
		{[]int32{1, 2, 2, 3}, 2},
		{[]int32{1, 1, 1, 1}, 0},
		{[]int32{1, 2, 3, 4, 5}, 4},
		{[]int32{1, 1, 2, 2, 3, 3, 3}, 4},
	}

	for _, test := range tests {
		result := equalizeArray(test.arr)
		if result != test.expected {
			t.Errorf("For input %v, expected %d but got %d", test.arr, test.expected, result)
		}
	}
}
