package main

import "fmt"

// The Utopian Tree goes through 2 cycles of growth every year. Each spring, it doubles in height.
// Each summer, its height increases by 1 meter.A Utopian Tree sapling with a height of 1 meter is
// planted at the onset of spring. How tall will the tree be after n growth cycles?

// For example, if the number of growth cycles is n = 5, the calculations are as follows:

// Period  Height
// 0          1
// 1          2
// 2          3
// 3          6
// 4          7
// 5          14

/*
 * Complete the 'utopianTree' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func utopianTree(n int32) int32 {
	// Write your code here
	var height int32 = 1
	for i := int32(1); i <= n; i++ {
		if i%2 == 0 {
			height++
		} else {
			height *= 2
		}
	}
	return height
}

func utopianTree1(n int32) int32 {
	// Write your code here
	var height int32 = 1
	for i := int32(1); i <= n; i++ {
		if i%2 == 0 {
			height++
		} else {
			height *= 2
		}
	}
	return height
}

func test() {
	fmt.Println(utopianTree1(5)) // 14
	fmt.Println(utopianTree1(4)) // 7
}

func main() {
	test()
}

// Algorithm
// Initialize a variable height to 1 to represent the initial height of the tree.
// Iterate from 1 to n using a for loop to simulate the growth cycles.
// If the current cycle is even, increment the height by 1.
// If the current cycle is odd, double the height.
// After the loop, return the final height of the tree.
// Time and Space Complexity
// Time Complexity:
// The time complexity of the function is O(n) because it iterates over the growth cycles from 1 to n.
// Space Complexity:
// The space complexity of the function is O(1) because it uses a constant amount of extra space regardless of the input size.
// Real-Time Use Cases
// Gardening:
// Calculating the height of a plant after a certain number of growth cycles.
// Forestry:
// Estimating the growth of trees in a forest over multiple seasons.
// Landscaping:
// Determining the height of trees in a garden after a specific number of years.

// Start
//  |
//  v
// Initialize height to 1
//  |
//  v
// For i = 1 to n
//  |  |
//  |  v
//  |  Is i % 2 == 1?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | Double the height  Increase height by 1
//  |  |
//  v  Repeat loop
//  |
//  v
// Return height
//  |
//  v
// End
