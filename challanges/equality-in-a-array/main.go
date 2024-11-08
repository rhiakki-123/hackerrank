package main

import "fmt"

// Given an array of integers, determine the minimum number of elements to delete to leave only
// elements of equal value.

// Example
// arr = [1, 2, 2, 3]
// Delete the 2 elements 1 and 3 leaving arr = [2, 2]. If both twos plus either the 2 or 3 are deleted,
// it takes 3 deletions to leave either [3] or [1].

/*
 * Complete the 'equalizeArray' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func equalizeArray(arr []int32) int32 {
	// Write your code here
	freq := make(map[int32]int32)
	maxFreq := int32(0)
	for _, num := range arr {
		freq[num]++
		if freq[num] > maxFreq {
			maxFreq = freq[num]
		}
	}
	return int32(len(arr)) - maxFreq // 4 - 1 = 3
}

func testEqualizeArray() {
	arr := []int32{1, 2, 2, 3}
	fmt.Println(equalizeArray(arr)) // Output: 3
}

func main() {
	testEqualizeArray()
}

//explain function equalizeArray(arr []int32) int32
// with its in details execution, algo, flowchart, time
// and space complexity, write its optimized version, unit tests and use cases

// Explanation of equalizeArray Function
// The equalizeArray function determines the minimum number of deletions required to make all elements in an array equal. The goal is to find the most frequent element in the array and then calculate how many elements need to be deleted to leave only that element.

// Function Signature
// func equalizeArray(arr []int32) int32

// arr: An array of integers.

// Logic
// Use a map to count the frequency of each element in the array.
// Find the maximum frequency from the map.
// The minimum number of deletions required is the total number of elements minus the maximum frequency.

// Algorithm
// Initialize a map to count the frequency of each element in the array.
// Iterate through the array and update the frequency map.
// Find the maximum frequency from the map.
// Calculate the minimum number of deletions as the total number of elements minus the maximum frequency.
// Return the minimum number of deletions.

// flowchart

// Start
//  |
//  v
// Initialize frequency map
//  |
//  v
// For each element in arr
//  |  |
//  |  v
//  |  Update frequency map
//  |  |
//  v  Repeat loop
//  |
//  v
// Find maximum frequency from map
//  |
//  v
// Calculate deletions as len(arr) - max frequency
//  |
//  v
// Return deletions
//  |
//  v
// End

// Detailed Execution
// Given the input:

// arr := []int32{3, 3, 2, 1, 3}

// initialize frequency map: {3: 3, 2: 1, 1: 1}
// freq := make(map[int32]int32)
// maxFreq := int32(0)

// Update frequency map:
// First Element (3)
// freq[3]++       // freq = {3: 1}
// freq = {3: 1}   // maxFreq = 1

// Second Element (3)
// freq[3]++       // freq = {3: 2}
// freq = {3: 2}   // maxFreq = 2

// Third Element (2)
// freq[2]++       // freq = {3: 2, 2: 1}
// freq = {3: 2, 2: 1} // maxFreq = 2

// Fourth Element (1)
// freq[1]++       // freq = {3: 2, 2: 1, 1: 1}
// freq = {3: 2, 2: 1, 1: 1} // maxFreq = 2

// Find maximum frequency:
// maxFreq := int32(0)
// for _, count := range freq {
//     if count > maxFreq {
//         maxFreq = count
//     }
// }
// // maxFreq = 3

// Calculate deletions:
// deletions := int32(len(arr)) - maxFreq // deletions = 5 - 3 = 2

// Return deletions: 2

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n), where n is the length of the array.
// This is because the function iterates through the array to update the frequency map and then iterates through the map to find the maximum frequency.
// Space Complexity
// The space complexity of the function is O(n).
// This is because the function uses an additional map to store the frequency of each element in the array.
// Real-Time Use Cases
// Data Analysis:

// Finding the most frequent element in a dataset and determining how many elements need to be removed to make all elements equal.
// Inventory Management:

// Identifying the most common item in an inventory and determining how many items need to be removed to standardize the inventory.
// Game Development:

// Implementing mechanics where players need to equalize elements by removing the least frequent ones.

// Optimized Code Implementation
// The original implementation is already optimized with a time complexity of O(n) and a space complexity of O(n). Here is the implementation for reference:

// package main

// import "fmt"

// func equalizeArray(arr []int32) int32 {
//     freq := make(map[int32]int32)
//     for _, num := range arr {
//         freq[num]++   // freq[num] = freq[num] + 1
//     }

//     maxFreq := int32(0)
//     for _, count := range freq {
//         if count > maxFreq {
//             maxFreq = count
//         }
//     }

//     deletions := int32(len(arr)) - maxFreq
//     return deletions
// }

// func testEqualizeArray() {
//     arr := []int32{3, 3, 2, 1, 3}
//     fmt.Println(equalizeArray(arr)) // Output: 2
// }

// func main() {
//     testEqualizeArray()
// }

// Unit Tests

// package main

// import "testing"

// func TestEqualizeArray(t *testing.T) {
//     tests := []struct {
//         arr      []int32
//         expected int32
//     }{
//         {[]int32{3, 3, 2, 1, 3}, 2},
//         {[]int32{1, 2, 2, 3}, 2},
//         {[]int32{1, 1, 1, 1}, 0},
//         {[]int32{1, 2, 3, 4, 5}, 4},
//         {[]int32{1, 1, 2, 2, 3, 3, 3}, 4},
//     }

//     for _, test := range tests {
//         result := equalizeArray(test.arr)
//         if result != test.expected {
//             t.Errorf("For input %v, expected %d but got %d", test.arr, test.expected, result)
//         }
//     }
// }

// Summary
// Logic: Use a map to count the frequency of each element in the array, find the maximum frequency, and calculate the minimum number of deletions required to make all elements equal.
// Algorithm: Initialize a frequency map, update the map with element frequencies, find the maximum frequency, calculate deletions, and return the result.
// Flowchart: Visual representation of the decision-making process for calculating the minimum number of deletions required to make all elements in the array equal.
// Time Complexity: O(n) because the function iterates through the array to update the frequency map and then iterates through the map to find the maximum frequency.
// Space Complexity: O(n) because the function uses an additional map to store the frequency of each element in the array.
// Use Cases: Data analysis, inventory management, game development.
