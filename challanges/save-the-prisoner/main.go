package main

import (
	"fmt"
)

// A jail has a number of prisoners and a number of treats to pass out to them. Their
// jailer decides the fairest way to divide the treats is to seat the prisoners
// around a circular table in sequentially numbered chairs. A chair number will
// be drawn from a hat. Beginning with the prisoner in that chair, one candy will
// be handed to each prisoner sequentially around the table until all have been
// distributed. The jailer is playing a little joke, though. The last piece of
// candy looks like all the others, but it tastes awful. Determine the chair number
// occupied by the prisoner who will receive that candy.

/*
 * Complete the 'saveThePrisoner' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 *  3. INTEGER s
 */

func saveThePrisoner(n int32, m int32, s int32) int32 {
	// Write your code here
	return ((m - 1 + s - 1) % n) + 1
}

func saveThePrisoner1(n int32, m int32, s int32) int32 {
	// Write your code here
	return ((m - 1 + s - 1) % n) + 1
}

func test() {
	fmt.Println(saveThePrisoner1(5, 2, 1))  // 2
	fmt.Println(saveThePrisoner1(5, 2, 2))  // 3
	fmt.Println(saveThePrisoner1(7, 19, 2)) // 6
}

func main() {
	test()
}

// Explanation of main.go ) Function
// The main.go ) function determines which prisoner will receive the last sweet in a circular distribution. The sweets are distributed starting from a given prisoner, and the distribution wraps around when it reaches the end of the list of prisoners.

// Function Signature
// func saveThePrisoner(n int32, m int32, s int32) int32

// n: The number of prisoners.
// m: The number of sweets.
// s: The starting position for distribution.

// Logic
// Calculate the position of the last sweet using the formula:
// (s + m - 1) % n

// s + m - 1: The position of the last sweet if there were no wrapping around.
// % n: Wraps around to the beginning if the position exceeds the number of prisoners.
// If the result is 0, it means the last sweet goes to the last prisoner (n).

// Algorithm
// Calculate the position of the last sweet using the formula (s + m - 1) % n.
// If the result is 0, return n.
// Otherwise, return the result.

// Start
//  |
//  v
// Calculate position as (s + m - 1) % n
//  |
//  v
// Is position == 0?
//  / \
// Yes  No
//  |    |
//  v    v
// Return n  Return position
//  |
//  v
// End

// position = (s + m - 1) % n = (1 + 2 - 1) % 5 = 2 % 5 = 2

// return 2

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(1).
// This is because the function performs a constant number of operations regardless of the input size.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the position variable.

// Real-Time Use Cases
// Resource Allocation:

// Determining the last resource allocation in a circular distribution system.
// Round-Robin Scheduling:

// Identifying the last task or process to be executed in a round-robin scheduling algorithm.
// Circular Buffers:

// Finding the last element in a circular buffer after a series of operations.
// Game Mechanics:

// Determining the last player to receive an item in a circular distribution in games.

// Summary
// Logic: Calculate the position of the last sweet using the formula (s + m - 1) % n. If the result is 0, return n; otherwise, return the result.
// Algorithm: Calculate the position of the last sweet, check if it is 0, and return the appropriate result.
// Flowchart: Visual representation of the decision-making process for determining the prisoner who receives the last sweet.
// Time Complexity: O(1) because the function performs a constant number of operations regardless of the input size.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Resource allocation, round-robin scheduling, circular buffers, game mechanics.
