package main

import "testing"

func TestHowManyGames(t *testing.T) {
	tests := []struct {
		p, d, m, s int32
		want       int32
	}{
		{20, 3, 6, 80, 6},
		{20, 3, 6, 85, 7},
		{20, 3, 6, 100, 10},
		{20, 3, 6, 10, 0},
		{20, 3, 6, 20, 1},
	}

	for _, tt := range tests {
		got := howManyGames(tt.p, tt.d, tt.m, tt.s)
		if got != tt.want {
			t.Errorf("howManyGames(%v, %v, %v, %v) = %v; want %v",
				tt.p, tt.d, tt.m, tt.s, got, tt.want)
		}
	}
}
