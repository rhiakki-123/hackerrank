package main

import "testing"

func TestBeautifulTriplets(t *testing.T) {
	tests := []struct {
		d    int32
		arr  []int32
		want int32
	}{
		{3, []int32{1, 2, 4, 5, 7, 8, 10}, 3},
		{1, []int32{2, 3, 4, 5}, 2},
		{2, []int32{1, 3, 5}, 1},
		{3, []int32{1, 4, 7}, 1},
		{1, []int32{1, 1, 1}, 0},
	}

	for _, tt := range tests {
		got := beautifulTriplets(tt.d, tt.arr)
		if got != tt.want {
			t.Errorf("beautifulTriplets(%v, %v) = %v; want %v",
				tt.d, tt.arr, got, tt.want)
		}
	}
}
