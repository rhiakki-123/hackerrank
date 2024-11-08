package main

import "fmt"

// There is a new mobile game that starts with consecutively numbered clouds. Some of the clouds
// are thunderheads and others are cumulus. The player can jump on any cumulus cloud having a number
// that is equal to the number of the current cloud plus 1 or 2. The player must avoid the thunderheads.
// Determine the minimum number of jumps it will take to jump from the starting postion to the last cloud.
// It is always possible to win the game.

// For each game, you will get an array of clouds numbered 0 if they are safe or 1 if they must be avoided.

// Example
// c = [0, 1, 0, 0, 0, 1, 0]
// Index the array from 0...6. The number on each cloud is its index in the list so the player must avoid the clouds at indices 1 and 5.
// They could follow these two paths: 0 -> 2 -> 4 -> 6 or 0 -> 2 -> 3 -> 4 -> 6. The first path takes 3 jumps while the second takes 4.

/*
 * Complete the 'jumpingOnClouds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY c as parameter.
 */

// optimized version of jumpingOnClouds function

func jumpingOnClouds(c []int32) int32 {
	jumps := int32(0)
	i := int32(0)
	for i < int32(len(c))-1 {
		if i+2 < int32(len(c)) && c[i+2] == 0 {
			i += 2
		} else {
			i += 1
		}
		jumps++
	}
	return jumps
}

func jumpingOnClouds1(c []int32) int32 {
	// Write your code here
	var jumps int32 = 0
	var i int32 = 0
	for i < int32(len(c)-1) {
		if i+2 < int32(len(c)) && c[i+2] == 0 {
			i += 2
		} else {
			i++
		}
		jumps++
	}
	return jumps
}

func test() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0}))              // 2
	fmt.Println(jumpingOnClouds1([]int32{0, 0, 1, 0, 1, 0, 1, 1})) // 5
}

func main() {
	test()
}

//explain function jumpingOnClouds(c []int32) int32
// with its in details execution, algo, flowchart, time
// and space complexity, write its optimized version, unit tests and use cases

// Explanation of jumpingOnClouds Function
// The jumpingOnClouds function determines the minimum number of jumps required to reach the last cloud in a game where clouds are represented by an array. The player can jump on clouds represented by 0 and must avoid clouds represented by 1. The player can jump to the next cloud or skip one cloud.

// Function Signature
// func jumpingOnClouds(c []int32) int32

// c: An array of integers representing the clouds. 0 indicates a safe cloud, and 1 indicates a thundercloud that must be avoided.

// Logic
// Initialize a variable jumps to 0 to keep track of the number of jumps.
// Initialize a variable i to 0 to represent the current position.
// Iterate through the array c using a while loop until the last cloud is reached.
// For each position, check if a jump of 2 clouds is possible (i.e., i + 2 is within bounds and c[i + 2] is a safe cloud).
// If a jump of 2 clouds is possible, increment i by 2; otherwise, increment i by 1.
// Increment the jumps counter for each jump.
// Return the total number of jumps.

// Algorithm
// Initialize jumps to 0.
// Initialize i to 0.
// While i is less than the last cloud:
// If i + 2 is within bounds and c[i + 2] is a safe cloud, increment i by 2.
// Otherwise, increment i by 1.
// Increment jumps.
// Return jumps.

// flowchart
// Start
//  |
//  v
// Initialize jumps to 0
// Initialize i to 0
//  |
//  v
// While i < len(c) - 1
//  |  |
//  |  v
//  |  Is i + 2 < len(c) and c[i + 2] == 0?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | Increment i by 2  Increment i by 1
//  |  |
//  |  v
//  | Increment jumps
//  |  |
//  v  Repeat loop
//  |
//  v
// Return jumps
//  |
//  v
// End

// Detailed Execution
// given the input

// c := []int32{0, 1, 0, 0, 0, 1, 0}

// Initialize variables:
// jumps := 0
// i := 0

// First Iteration (i = 0):
// i + 2 < len(c) (2 < 7) and c[i + 2] == 0 (c[2] == 0): Yes
// Increment i by 2: i = 2

// Increment jumps:
// jumps = 1

// Second Iteration (i = 2):
// i + 2 < len(c) (4 < 7) and c[i + 2] == 0 (c[4] == 0): Yes
// Increment i by 2: i = 4

// Increment jumps:
// jumps = 2

// Third Iteration (i = 4):
// i + 2 < len(c) (6 < 7) and c[i + 2] == 0 (c[6] == 0): Yes
// Increment i by 2: i = 6

// Increment jumps:
// jumps = 3

// End of loop: Return jumps = 3

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n), where n is the length of the array c.
// This is because the function iterates through the array c once, making jumps based on the conditions.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the jumps and i variables.

// Real-Time Use Cases
// Pathfinding:

// Finding the shortest path in a grid or maze where certain cells are blocked.
// Game Development:

// Implementing jump mechanics in platformer games where the player must avoid obstacles.
// Robotics:

// Programming robots to navigate through environments with obstacles.
// Network Routing:

// Determining the shortest path in a network with certain nodes or links blocked.

// Optimized Version of jumpingOnClouds Function
// package main

// import "fmt"

// func jumpingOnClouds(c []int32) int32 {
//     jumps := int32(0)
//     i := int32(0)
//     for i < int32(len(c))-1 {
//         if i+2 < int32(len(c)) && c[i+2] == 0 {
//             i += 2
//         } else {
//             i += 1
//         }
//         jumps++
//     }
//     return jumps
// }

// func testJumpingOnClouds() {
//     c := []int32{0, 1, 0, 0, 0, 1, 0}
//     fmt.Println(jumpingOnClouds(c)) // Output: 3
// }

// func main() {
//     testJumpingOnClouds()
// }

// Unit Tests
// package main

// import "testing"

// func TestJumpingOnClouds(t *testing.T) {
//     tests := []struct {
//         c        []int32
//         expected int32
//     }{
//         {[]int32{0, 1, 0, 0, 0, 1, 0}, 3},
//         {[]int32{0, 0, 0, 0, 1, 0}, 3},
//         {[]int32{0, 0, 1, 0, 0, 1, 0}, 4},
//         {[]int32{0, 0, 0, 1, 0, 0}, 3},
//     }

//     for _, test := range tests {
//         result := jumpingOnClouds(test.c)
//         if result != test.expected {
//             t.Errorf("For input %v, expected %d but got %d", test.c, test.expected, result)
//         }
//     }
// }

// Summary
// Logic: Iterate through the array c, making jumps of 1 or 2 clouds based on the conditions, and count the number of jumps required to reach the last cloud.
// Algorithm: Initialize jumps and i, iterate through the array c using a while loop, make jumps based on the conditions, and return the total number of jumps.
// Flowchart: Visual representation of the decision-making process for calculating the minimum number of jumps required to reach the last cloud.
// Time Complexity: O(n) because the function iterates through the array c once.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Pathfinding, game development, robotics, network routing.
