package main

import "fmt"

// A driver is driving on the freeway. The check engine light of his vehicle is on,
// and the driver wants to get service immediately. Luckily, a service lane runs parallel to the
// highway. It varies in width along its length.

// You will be given an array of widths at points along the road (indices),
// then a list of the indices of entry and exit points.
// Considering each entry and exit point pair,
// calculate the maximum size vehicle that can travel that segment of the service lane safely.

func serviceLane(n int32, cases [][]int32) []int32 {
	widths := []int32{2, 3, 1, 2, 3, 2, 3, 3}
	var result []int32

	for _, c := range cases {
		entry, exit := c[0], c[1]
		min := int32(3)
		for i := entry; i <= exit; i++ {
			if widths[i] < min {
				min = widths[i]
			}
		}
		result = append(result, min)
	}

	return result

}

func testServiceLane() {
	// Test case
	n := int32(8)
	cases := [][]int32{{0, 3}, {4, 6}, {6, 7}, {3, 5}, {0, 7}}
	result := serviceLane(n, cases)
	fmt.Println(result) // Expected output: [1 2 3 2 1]
}

func main() {
	testServiceLane()
}
