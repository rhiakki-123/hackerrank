package main

import "testing"

func TestAcmTeam(t *testing.T) {
	tests := []struct {
		topic    []string
		expected []int32
	}{
		{[]string{"10101", "11100", "11010", "00101"}, []int32{5, 2}},
		{[]string{"101", "110", "001"}, []int32{3, 2}},
		{[]string{"111", "111", "111"}, []int32{3, 3}},
		{[]string{"000", "000", "000"}, []int32{0, 3}},
		{[]string{"10101", "11100", "11010"}, []int32{5, 1}},
	}

	for _, test := range tests {
		result := acmTeam(test.topic)
		if result[0] != test.expected[0] || result[1] != test.expected[1] {
			t.Errorf("For input %v, expected %v but got %v", test.topic, test.expected, result)
		}
	}
}
