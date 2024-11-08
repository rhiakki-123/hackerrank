package main

import "fmt"

// Given a sequence of integers a, a triplet (a[i], a[j], a[k]) is beautiful if:
// i<j<k
// a[j] - a[i] = a[k] - a[j] = d
// Given an increasing sequence of integers and the value of d,
// count the number of beautiful triplets in the sequence.

/*
 * Complete the 'beautifulTriplets' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER d
 *  2. INTEGER_ARRAY arr
 */

//Optimized version
func optbeautifulTriplets(d int32, arr []int32) int32 {
	count := int32(0)
	numMap := make(map[int32]bool)

	// Create hashmap for O(1) lookups
	for _, num := range arr {
		numMap[num] = true
	}

	// Check each number if it forms beautiful triplet
	for _, num := range arr {
		if numMap[num+d] && numMap[num+2*d] {
			count++
		}
	}
	return count
}

func beautifulTriplets(d int32, arr []int32) int32 {
	// Write your code here
	var count int32
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]-arr[i] == d {
				for k := j + 1; k < len(arr); k++ {
					if arr[k]-arr[j] == d {
						count++
					}
				}
			}
		}
	}
	return count
}

func testBeautifulTriplets() {
	// Test case
	d := int32(3)
	arr := []int32{1, 2, 4, 5, 7, 8, 10}
	fmt.Println(optbeautifulTriplets(d, arr)) // Expected output: 3
}

func main() {
	testBeautifulTriplets()
}

//Explain function beautifulTriplets(d int32, arr []int32)
//Logic
//Flowchart
//In depts loop executions
//Algorithm
//Time Complexity
//Space Complexity
//Optimized version
//use cases
//unit tests

// Function Purpose
// beautifulTriplets finds the number of beautiful triplets (i,j,k) where:

// arr[j] - arr[i] = d
// arr[k] - arr[j] = d
// i < j < k

// graph TD
//     A[Start] --> B[Input: d and arr]
//     B --> C[Initialize count = 0]
//     C --> D[Triple nested loops]
//     D --> E[Check if differences = d]
//     E --> F[Increment count if beautiful]
//     F --> G[Return count]
//     G --> H[End]

// Algorithm
// func beautifulTriplets(d int32, arr []int32) int32 {
//     count := int32(0)
//     n := len(arr)

//     for i := 0; i < n-2; i++ {
//         for j := i+1; j < n-1; j++ {
//             if arr[j]-arr[i] == d {
//                 for k := j+1; k < n; k++ {
//                     if arr[k]-arr[j] == d {
//                         count++
//                     }
//                 }
//             }
//         }
//     }
//     return count
// }

// optimized version
// func beautifulTriplets(d int32, arr []int32) int32 {
//     count := int32(0)
//     numMap := make(map[int32]bool)

//     // Create hashmap for O(1) lookups
//     for _, num := range arr {
//         numMap[num] = true
//     }

//     // Check each number if it forms beautiful triplet
//     for _, num := range arr {
//         if numMap[num+d] && numMap[num+2*d] {
//             count++
//         }
//     }
//     return count
// }

// complexity analysis
// Complexity Analysis
// Original version:
// Time: O(n³)
// Space: O(1)
// Optimized version:
// Time: O(n)
// Space: O(n)

// unit tests
// func TestBeautifulTriplets(t *testing.T) {
//     tests := []struct {
//         d    int32
//         arr  []int32
//         want int32
//     }{
//         {3, []int32{1, 2, 4, 5, 7, 8, 10}, 3},
//         {1, []int32{2, 3, 4, 5}, 2},
//         {2, []int32{1, 3, 5}, 1},
//         {3, []int32{1, 4, 7}, 1},
//         {1, []int32{1, 1, 1}, 0},
//     }

//     for _, tt := range tests {
//         got := beautifulTriplets(tt.d, tt.arr)
//         if got != tt.want {
//             t.Errorf("beautifulTriplets(%v, %v) = %v; want %v",
//                      tt.d, tt.arr, got, tt.want)
//         }
//     }
// }

// Use Cases
// Finding arithmetic sequences with fixed difference
// Pattern matching in numerical sequences
// Data analysis for equidistant elements
// Mathematical sequence validation

// Step-by-Step Execution
// Initialization:

// count is initialized to 0.
// numMap is an empty map of type map[int32]bool.
// First Loop: Creating the HashMap

// This loop iterates over each element in arr and adds it to numMap with the value true.

// arr := []int32{1, 2, 4, 5, 7, 8, 10}

// After the loop, numMap will be:
// numMap = {1: true, 2: true, 4: true, 5: true, 7: true, 8: true, 10: true}

// Second Loop: Checking for Beautiful Triplets

// This loop iterates over each element in arr and checks if the elements num + d and num + 2*d exist in numMap.
// If both elements exist, it means a beautiful triplet is found, and count is incremented.
// Example Execution:
// For num = 1:
// Check if numMap[1 + 3] (i.e., numMap[4]) exists: true
// Check if numMap[1 + 2*3] (i.e., numMap[7]) exists: true
// Both conditions are true, so count is incremented to 1.
// For num = 2:
// Check if numMap[2 + 3] (i.e., numMap[5]) exists: true
// Check if numMap[2 + 2*3] (i.e., numMap[8]) exists: true
// Both conditions are true, so count is incremented to 2.
// For num = 4:
// Check if numMap[4 + 3] (i.e., numMap[7]) exists: true
// Check if numMap[4 + 2*3] (i.e., numMap[10]) exists: true
// Both conditions are true, so count is incremented to 3.
// For num = 5:
// Check if numMap[5 + 3] (i.e., numMap[8]) exists: true
// Check if numMap[5 + 2*3] (i.e., numMap[11]) exists: false
// One condition is false, so count remains 3.
// For num = 7, 8, and 10:
// The conditions will not be met for these values, so count remains 3.
// Return the Result:

// The final value of count is returned, which is the number of beautiful triplets found.
// Summary
// The function first creates a hashmap for quick lookups.
// It then checks each element in the array to see if it can form a beautiful triplet with the given difference d.
// The time complexity of this optimized version is O(n), and the space complexity is O(n).
