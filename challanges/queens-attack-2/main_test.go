package main

import "testing"

func TestQueensAttack(t *testing.T) {
	tests := []struct {
		n         int32
		k         int32
		r_q       int32
		c_q       int32
		obstacles [][]int32
		expected  int32
	}{
		{8, 1, 4, 4, [][]int32{{3, 5}}, 24},
		{8, 0, 4, 4, [][]int32{}, 27},
		{8, 2, 4, 4, [][]int32{{3, 5}, {5, 4}}, 20},
		{5, 3, 4, 3, [][]int32{{5, 5}, {4, 2}, {2, 3}}, 10},
	}

	for _, test := range tests {
		result := queensAttack(test.n, test.k, test.r_q, test.c_q, test.obstacles)
		if result != test.expected {
			t.Errorf("For input n=%d, k=%d, r_q=%d, c_q=%d, obstacles=%v, expected %d but got %d", test.n, test.k, test.r_q, test.c_q, test.obstacles, test.expected, result)
		}
	}
}
