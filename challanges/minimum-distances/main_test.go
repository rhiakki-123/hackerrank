package main

import "testing"

func TestMinimumDistances1(t *testing.T) {
	tests := []struct {
		arr  []int32
		want int32
	}{
		{[]int32{7, 1, 3, 4, 1, 7}, 3},
		{[]int32{1, 2, 3, 4, 10}, -1},
		{[]int32{1, 1, 1, 1}, 1},
		{[]int32{1, 2, 3, 4, 2, 1}, 3},
		{[]int32{1, 2, 3, 4, 5}, -1},
	}

	for _, tt := range tests {
		got := minimumDistances1(tt.arr)
		if got != tt.want {
			t.Errorf("minimumDistances1(%v) = %v; want %v", tt.arr, got, tt.want)
		}
	}
}
