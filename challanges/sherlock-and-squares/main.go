package main

import (
	"fmt"
	"math"
)

// Watson likes to challenge Sherlocks math ability. He will provide a starting and ending value that describe
// a range of integers, inclusive of the endpoints. Sherlock must determine the number of square integers within that range.

// Note: A square integer is an integer which is the square of an integer, e.g.  1, 4, 9, 16, 25.

// For example, the range is a=24 and b=49, inclusive. There are three square integers in the range: 25, 36 and 49.

// int a: the lower range boundary
// int b: the upper range boundary

// Returns
// int: the number of square integers in the range

/*
 * Complete the 'squares' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER a
 *  2. INTEGER b
 */

// Optimized Code Implementation
func squares(a int32, b int32) int32 {
	start := int32(math.Ceil(math.Sqrt(float64(a))))
	end := int32(math.Floor(math.Sqrt(float64(b))))
	if start > end {
		return 0
	}
	return end - start + 1
}

func squares1(a int32, b int32) int32 {
	// Write your code here
	var count int32
	for i := a; i <= b; i++ {
		if isSquare(i) {
			count++
		}
	}
	return count
}

func isSquare(n int32) bool {
	sqrt := int32(1)
	for sqrt*sqrt < n {
		sqrt++
	}
	return sqrt*sqrt == n
}

func test() {
	fmt.Println(squares(3, 9))   // 2
	fmt.Println(squares(17, 24)) // 0
	fmt.Println(squares(24, 49)) // 3
}

func main() {
	test()
}

//explain function  func squares(a int32, b int32) int32 with its logic, in details execution, algo, flowchart, timeand space complexity and use cases

// Explanation of squares Function
// The squares function calculates the number of perfect squares between two given integers a and b (inclusive). A perfect square is an integer that is the square of another integer.

// Function Signature
// func squares(a int32, b int32) int32

// a: The lower bound of the range.
// b: The upper bound of the range.

// Logic
// Calculate the smallest integer start whose square is greater than or equal to a.
// Calculate the largest integer end whose square is less than or equal to b.
// The number of perfect squares between a and b is the difference between end and start, plus one.

// Algorithm
// Calculate start as the ceiling of the square root of a.
// Calculate end as the floor of the square root of b.
// If start is greater than end, return 0.
// Otherwise, return end - start + 1.

// flowchart
// Start
//  |
//  v
// Calculate start as ceil(sqrt(a))
//  |
//  v
// Calculate end as floor(sqrt(b))
//  |
//  v
// Is start > end?
//  / \
// Yes  No
//  |    |
//  v    v
// Return 0  Return end - start + 1
//  |
//  v
// End

// Detailed Execution
// Given the inputs:
// a := int32(3)
// b := int32(9)

// Calculate start:
// start := int32(math.Ceil(math.Sqrt(float64(a)))) // ceil(sqrt(3)) = 2

// Calculate end:
// end := int32(math.Floor(math.Sqrt(float64(b)))) // floor(sqrt(9)) = 3

// Check if start is greater than end:
// No

// Return end - start + 1:
// 3 - 2 + 1 = 2

// Final State:
// start = 2
// end = 3
// Output:
// 2

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(1).
// This is because the function performs a constant number of operations regardless of the input size.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the start and end variables.

// Real-Time Use Cases
// Mathematical Analysis:

// Counting the number of perfect squares within a given range for mathematical studies.
// Data Analysis:

// Analyzing datasets to find the number of perfect squares within specific intervals.
// Game Development:

// Implementing game mechanics that involve counting perfect squares within a range.

// Summary
// Logic: Calculate the smallest integer whose square is greater than or equal to a and the largest integer whose square is less than or equal to b, then return the difference plus one.
// Algorithm: Calculate start and end, check if start is greater than end, and return the number of perfect squares.
// Flowchart: Visual representation of the decision-making process for calculating the number of perfect squares between a and b.
// Time Complexity: O(1) because the function performs a constant number of operations regardless of the input size.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Mathematical analysis, data analysis, game development.
