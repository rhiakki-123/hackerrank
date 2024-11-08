package main

import "fmt"

// You are given a number of sticks of varying lengths. You will iteratively cut the sticks into smaller sticks,
// discarding the shortest pieces until there are none left. At each iteration you will determine the length of the
// shortest stick remaining, cut that length from each of the longer sticks and then discard all the pieces of that
// shortest length. When all the remaining sticks are the same length, they cannot be shortened so discard them.
// Given the lengths of n sticks, print the number of sticks that are left before each iteration until there are none left.

// For example, there are n = 3 sticks of lengths arr = [1, 2, 3]. The shortest stick length is 1, so we cut that length
// from the longer two and discard the pieces of length 1. Now our lengths are arr = [1, 2]. Again, the shortest stick is
// of length 1, so we cut that amount from the longer stick and discard those pieces. There is only one stick left, arr = [1],
// so we discard that stick. Our lengths are answer = [3, 2, 1].

/*
 * Complete the 'cutTheSticks' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func cutTheSticks(arr []int32) []int32 {
	var result []int32
	for len(arr) > 0 {
		result = append(result, int32(len(arr)))
		minVal := findMin(arr)
		var newArr []int32
		for _, val := range arr {
			if val-minVal > 0 {
				newArr = append(newArr, val-minVal)
			}
		}
		arr = newArr
	}
	return result
}

func findMin(arr []int32) int32 {
	minVal := arr[0]
	for _, val := range arr {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func cutTheSticks1(arr []int32) []int32 {
	// Write your code here
	var result []int32
	for len(arr) > 0 {
		result = append(result, int32(len(arr)))
		min := arr[0]
		for _, v := range arr {
			if v < min {
				min = v
			}
		}
		var temp []int32
		for _, v := range arr {
			if v > min {
				temp = append(temp, v-min)
			}
		}
		arr = temp
	}
	return result
}

func test() {
	arr := []int32{1, 2, 3}
	arr1 := []int32{5, 4, 4, 2, 2, 8}
	fmt.Println(cutTheSticks(arr))  // [3, 2, 1]
	fmt.Println(cutTheSticks(arr1)) // [6, 4, 2, 1]
}

func main() {
	test()
}

//explain function  func cutTheSticks(arr []int32) []int32  with its logic, in details execution, algo, flowchart, timeand space complexity and use cases

// Explanation of cutTheSticks Function
// The cutTheSticks function simulates the process of cutting sticks. Given an array of stick lengths, the function repeatedly cuts the sticks by the length of the shortest stick until no sticks are left. The function returns an array representing the number of sticks that are cut in each iteration.

// Function Signature
// func cutTheSticks(arr []int32) []int32

// Algorithm
// Initialize an empty result array.
// While the array is not empty:
// Append the length of the array to the result array.
// Find the minimum value in the array.
// Subtract the minimum value from each element in the array.
// Remove all elements that are zero from the array.
// Return the result array.

// flowchart
// Start
//  |
//  v
// Initialize result as empty array
//  |
//  v
// While arr is not empty
//  |  |
//  |  v
//  |  Append len(arr) to result
//  |  |
//  |  v
//  |  Find min value in arr
//  |  |
//  |  v
//  |  Subtract min value from each element in arr
//  |  |
//  |  v
//  |  Remove all elements that are zero from arr
//  |  |
//  v  Repeat loop
//  |
//  v
// Return result
//  |
//  v
// End

// detailed Execution
// Given the input:
// arr := []int32{5, 4, 4, 2, 2, 8}

// Initialize result as empty array: result := []
// result := []int32{}

// First iteration:
// Append the length of the array to the result
// result = append(result, int32(len(arr))) // result = [6]
// Find the minimum value in the array: min = 2
// minVal := int32(2)
// Subtract the minimum value from each element in the array
// arr = []int32{3, 2, 2, 0, 0, 6}
// Remove all elements that are zero from the array
// arr = []int32{3, 2, 2, 6}

// Second iteration:
// Append the length of the array to the result
// result = append(result, int32(len(arr))) // result = [6, 4]
// Find the minimum value in the array: min = 2
// minVal := int32(2)
// Subtract the minimum value from each element in the array
// arr = []int32{1, 0, 0, 4}
// Remove all elements that are zero from the array
// arr = []int32{1, 4}

// Third iteration:
// Append the length of the array to the result
// result = append(result, int32(len(arr))) // result = [6, 4, 2]
// Find the minimum value in the array: min = 1
// minVal := int32(1)
// Subtract the minimum value from each element in the array
// arr = []int32{0, 3}
// Remove all elements that are zero from the array
// arr = []int32{3}

// Fourth iteration:
// Append the length of the array to the result
// result = append(result, int32(len(arr))) // result = [6, 4, 2, 1]

// Return result // [6, 4, 2, 1]

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n^2), where n is the length of the array.
// This is because, in the worst case, the function iterates through the array multiple times, and each iteration involves finding the minimum value and removing elements.
// Space Complexity
// The space complexity of the function is O(n).
// This is because the function uses an additional result array to store the number of sticks cut in each iteration.

// Real-Time Use Cases
// Resource Management:

// Simulating the process of resource consumption and reduction over time.
// Game Development:

// Implementing mechanics where objects are repeatedly reduced in size or quantity.
// Data Analysis:

// Analyzing the reduction of data points over multiple iterations.
// Code Implementation

// GitHub Copilot
// Explanation of cutTheSticks Function
// The cutTheSticks function simulates the process of cutting sticks. Given an array of stick lengths, the function repeatedly cuts the sticks by the length of the shortest stick until no sticks are left. The function returns an array representing the number of sticks that are cut in each iteration.

// Function Signature
// arr: An array of integers representing the lengths of the sticks.
// Logic
// Initialize an empty result array to store the number of sticks cut in each iteration.
// While there are sticks left in the array:
// Append the current number of sticks to the result array.
// Find the length of the shortest stick.
// Subtract the length of the shortest stick from each stick.
// Remove all sticks of length zero from the array.
// Return the result array.
// Algorithm
// Initialize an empty result array.
// While the array is not empty:
// Append the length of the array to the result array.
// Find the minimum value in the array.
// Subtract the minimum value from each element in the array.
// Remove all elements that are zero from the array.
// Return the result array.
// Flowchart
// Detailed Execution
// Given the input:

// Initialize variables:

// First Iteration:

// Append the length of the array to the result:
// Find the minimum value in the array:
// Subtract the minimum value from each element in the array:
// Remove all elements that are zero from the array:
// Second Iteration:

// Append the length of the array to the result:
// Find the minimum value in the array:
// Subtract the minimum value from each element in the array:
// Remove all elements that are zero from the array:
// Third Iteration:

// Append the length of the array to the result:
// Find the minimum value in the array:
// Subtract the minimum value from each element in the array:
// Remove all elements that are zero from the array:
// Fourth Iteration:

// Append the length of the array to the result:
// Find the minimum value in the array:
// Subtract the minimum value from each element in the array:
// Remove all elements that are zero from the array:
// Final State:
// result = [6, 4, 2, 1]
// Output:
// [6, 4, 2, 1]
// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n^2), where n is the length of the array.
// This is because, in the worst case, the function iterates through the array multiple times, and each iteration involves finding the minimum value and removing elements.
// Space Complexity
// The space complexity of the function is O(n).
// This is because the function uses an additional result array to store the number of sticks cut in each iteration.
// Real-Time Use Cases
// Resource Management:

// Simulating the process of resource consumption and reduction over time.
// Game Development:

// Implementing mechanics where objects are repeatedly reduced in size or quantity.
// Data Analysis:

// Analyzing the reduction of data points over multiple iterations.
// Code Implementation
// Summary
// Logic: Repeatedly cut the sticks by the length of the shortest stick until no sticks are left, and return the number of sticks cut in each iteration.
// Algorithm: Initialize an empty result array, iterate through the array while it is not empty, find the minimum value, subtract the minimum value from each element, remove zero elements, and return the result array.
// Flowchart: Visual representation of the decision-making process for cutting the sticks and counting the number of sticks cut in each iteration.
// Time Complexity: O(n^2) because the function iterates through the array multiple times, and each iteration involves finding the minimum value and removing elements.
// Space Complexity: O(n) because the function uses an additional result array to store the number of sticks cut in each iteration.
// Use Cases: Resource management, game development, data analysis.
