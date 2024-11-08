package main

import "fmt"

// John Watson knows of an operation called a right circular rotation on an array of integers.
// One rotation operation moves the last array element to the first position and shifts all
// emaining elements right one. To test Sherlocks abilities, Watson provides Sherlock with an array of
// integers. Sherlock is to perform the rotation operation a number of times then determine the
// value of the element at a given position.
// For each array, perform a number of right circular rotations and return the values of the elements
// at the given indices.

/*
 * Complete the 'circularArrayRotation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER k
 *  3. INTEGER_ARRAY queries
 */

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	var result []int32
	for _, q := range queries {
		result = append(result, a[(len(a)-int(k)+int(q))%len(a)])
	}
	return result
}

func circularArrayRotation1(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	effectiveRotations := k % int32(len(a))
	result := []int32{}
	for _, q := range queries {
		originalIndex := (len(a) - int(effectiveRotations) + int(q)) % len(a)
		result = append(result, a[originalIndex])
	}
	return result
}

func test() {
	fmt.Println(circularArrayRotation([]int32{3, 4, 5}, 2, []int32{1, 2}))     // [5, 3]
	fmt.Println(circularArrayRotation1([]int32{1, 2, 3}, 2, []int32{0, 1, 2})) // [2, 3, 1]
}

func test1() {
	a := []int32{1, 2, 3}
	k := 2
	queries := []int32{0, 1, 2}
	fmt.Println(circularArrayRotation(a, int32(k), queries)) // Output: [2, 3, 1]
}

func main() {
	test()
}

//1 2 3 - 1 - 3 1 2 - 2 - 2 3 1

// Explanation of circularArrayRotation Function
// The circularArrayRotation function performs a specified number of rotations on an array and then returns the values at specified indices after the rotations.

// Function Signature

// func circularArrayRotation(a []int32, k int32, queries []int32) []int32

// a: The array to be rotated.
// k: The number of rotations to perform.
// queries: An array of indices for which the values need to be returned after the rotations.

// Logic
// Calculate the effective number of rotations using k % len(a). This is because rotating an array by its length results in the same array.
// For each query, calculate the index in the original array that corresponds to the queried index after the rotations.
// Use the formula (len(a) - int(k) + int(q)) % len(a) to find the original index for each query.
// Return the values at the calculated indices.

// Algorithm
// Calculate the effective number of rotations: effectiveRotations = k % len(a).
// Initialize an empty result array.
// Iterate through each query in the queries array:
// Calculate the original index using the formula (len(a) - int(effectiveRotations) + int(q)) % len(a).
// Append the value at the calculated index to the result array.
// Return the result array.

// Flowchart
// Start
//  |
//  v
// Calculate effectiveRotations as k % len(a)
//  |
//  v
// Initialize result as empty array
//  |
//  v
// For each query in queries
//  |  |
//  |  v
//  |  Calculate originalIndex as (len(a) - int(effectiveRotations) + int(query)) % len(a)
//  |  |
//  |  v
//  |  Append a[originalIndex] to result
//  |  |
//  v  Repeat loop
//  |
//  v
// Return result
//  |
//  v
// End

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n + q), where n is the length of the array a and q is the number of queries.
// This is because the function iterates through each query and performs a constant-time calculation for each query.
// Space Complexity
// The space complexity of the function is O(q).
// This is because the function uses an additional result array to store the values for each query.
// Real-Time Use Cases
// Data Shuffling:

// Rotating elements in a dataset for shuffling purposes.
// Circular Buffers:

// Accessing elements in a circular buffer after a certain number of rotations.
// Game Mechanics:

// Implementing circular rotations in game mechanics, such as rotating a game board or elements.
// Scheduling:

// Rotating schedules or tasks in a circular manner.

// Summary
// Logic: Calculate the effective number of rotations, then use the formula (len(a) - int(effectiveRotations) + int(q)) % len(a) to find the original index for each query.
// Algorithm: Calculate the effective number of rotations, iterate through each query, calculate the original index, and return the values at the calculated indices.
// Flowchart: Visual representation of the decision-making process for calculating the values at specified indices after rotations.
// Time Complexity: O(n + q) because the function iterates through each query and performs a constant-time calculation for each query.
// Space Complexity: O(q) because the function uses an additional result array to store the values for each query.
// Use Cases: Data shuffling, circular buffers, game mechanics, scheduling.
