package main

import "fmt"

// A video player plays a game in which the character competes in a hurdle race. Hurdles
// are of varying heights, and the characters have a maximum height they can jump. There is
// a magic potion they can take that will increase their maximum jump height by 1 unit for
//  each dose. How many doses of the potion must the character take to be able to jump
//  all of the hurdles. If the character can already clear all of the hurdles, return 0.

//  height = [1, 2, 3, 3, 2]
//  k = 1

//  The character can jump 1 unit high initially and must take 3-1 = 2 doses of potion to be able to jump all of the hurdles.

/*
 * Complete the 'hurdleRace' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY height
 */

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	var max int32 = 0
	for _, h := range height {
		if h > max {
			max = h
		}
	}
	if max > k {
		return max - k
	}
	return 0
}

func hurdleRace1(k int32, height []int32) int32 {
	// Write your code here
	var max int32 = 0
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	if max > k {
		return max - k
	}
	return 0
}

func test() {
	fmt.Println(hurdleRace1(1, []int32{1, 2, 3, 3, 2})) // 2
	fmt.Println(hurdleRace1(4, []int32{1, 6, 3, 5, 2})) // 2
	fmt.Println(hurdleRace1(7, []int32{2, 5, 4, 5, 2})) // 0
}

func main() {
	test()
}

// Time and Space Complexity
// Time Complexity:
// The time complexity of the function is O(n) because it iterates over the height array once to find the maximum height of the hurdles.
// Space Complexity:
// The space complexity of the function is O(1) because it uses a constant amount of extra space regardless of the input size.
// Real-Time Use Cases
// Use Cases
// Gaming:
// Determining the number of power-ups needed for a character to overcome obstacles in a game.
// Sports:
// Calculating the number of training sessions needed for an athlete to clear a certain height in high jump.
// Robotics:
// Determining the number of adjustments needed for a robot to overcome obstacles of varying heights.
// Construction:
// Calculating the number of additional materials needed to build a structure to a certain heigh

// Algorithm
// Initialize a variable max to 0 to keep track of the maximum height of the hurdles.
// Iterate through each element in the height array using a for loop.
// For each element, check if the current height is greater than max.
// If the current height is greater than max, update max to the current height.
// After the loop, check if max is greater than k.
// If max is greater than k, return max - k (the number of doses needed).
// If max is not greater than k, return 0 (no doses needed).

// Start
//  |
//  v
// Initialize max to 0
//  |
//  v
// For i = 0 to len(height)-1
//  |  |
//  |  v
//  |  Is height[i] > max?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  |  Update max to height[i]
//  |  |
//  v  Repeat loop
//  |
//  v
// Is max > k?
//  / \
// Yes  No
//  |    |
//  v    v
// Return max - k  Return 0
//  |
//  v
// End
