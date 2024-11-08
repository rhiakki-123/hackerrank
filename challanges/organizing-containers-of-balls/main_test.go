package main

import (
	"testing"
)

func TestOrganizingContainers(t *testing.T) {
	tests := []struct {
		container [][]int32
		expected  string
	}{
		{
			container: [][]int32{
				{1, 1},
				{1, 1},
			},
			expected: "Possible",
		},
		{
			container: [][]int32{
				{0, 2, 1},
				{1, 1, 1},
				{2, 0, 0},
			},
			expected: "Possible",
		},
		{
			container: [][]int32{
				{2, 0, 0},
				{0, 2, 0},
				{0, 0, 2},
			},
			expected: "Possible",
		},
		{
			container: [][]int32{
				{999336263, 998799923},
				{998799923, 999763019},
			},
			expected: "Possible",
		},
		{
			container: [][]int32{
				{997612619, 934920795, 998879231, 999926463},
				{960369681, 997828120, 999792735, 979622676},
				{999013654, 998634077, 997988323, 958769423},
				{997409523, 999301350, 940952923, 993020546},
			},
			expected: "Possible",
		},
	}

	for _, test := range tests {
		result := organizingContainers(test.container)
		if result != test.expected {
			t.Errorf("For input %v, expected %s but got %s", test.container, test.expected, result)
		}
	}
}
