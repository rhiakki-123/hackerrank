package main

import "fmt"

// The distance between two array values is the number of indices between them. Given ,
// find the minimum distance between any pair of equal elements in the array.
// If no such value exists, return -1.

// Example
// arr = [7, 1, 3, 4, 1, 7]
// There are two matching pairs of values: 7 and 1.
// The indices of the 7s are i = 0 and j = 5,
// so their distance is d[i, j] = |j - i| = 5.
// The indices of the 1s are i = 1 and j = 4,
// so their distance is d[i, j] = |j - i| = 3.

//optimal solution
func minimumDistances(arr []int32) int32 {
	// Write your code here
	var minDistance int32 = -1
	distances := make(map[int32]int32)
	for i, num := range arr {
		if j, ok := distances[num]; ok {
			distance := int32(i) - j
			if minDistance == -1 || distance < minDistance {
				minDistance = distance
			}
		}
		distances[num] = int32(i)
	}
	return minDistance
}

// use other approach
func minimumDistances1(arr []int32) int32 {
	var minDistance int32 = -1
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				distance := int32(j - i)
				if minDistance == -1 || distance < minDistance {
					minDistance = distance
				}
			}
		}
	}
	return minDistance
}

func testMinimumDistances() {
	// Test case
	arr := []int32{7, 1, 3, 4, 1, 7}
	arr1 := []int32{1, 2, 3, 4, 2, 1}
	result := minimumDistances1(arr1)
	result1 := minimumDistances(arr)
	fmt.Println(result1) // Expected output: 3
	fmt.Println(result)  // Expected output: 3
}

func main() {
	testMinimumDistances()
}

// Explain in depth function func minimumDistances1(arr []int32) int32
// Logic
// Flowchart
// loop execution for each element in the array
// Algorithm
// Time Complexity
// Space Complexity
// Optimized version
// unit tests
// use cases

// minimumDistances1 finds the minimum distance between any two equal elements in the array arr.

// Logic
// Iterate through the array and store the indices of each element in a map.
// For each element, calculate the distance between its indices and keep track of the minimum distance found.

// Flowchart
// graph TD
//     A[Start] --> B[Input: arr]
//     B --> C[Initialize minDist to -1]
//     C --> D[Initialize indexMap]
//     D --> E[Loop through array]
//     E --> F[Check if element exists in indexMap]
//     F --> G[Calculate distance]
//     G --> H[Update minDist if smaller]
//     H --> I[Update indexMap with current index]
//     I --> J[End loop]
//     J --> K[Return minDist]
//     K --> L[End]

// Algorithm
// func minimumDistances1(arr []int32) int32 {
//     minDist := int32(-1)
//     indexMap := make(map[int32]int)

//     for i, num := range arr {
//         if prevIndex, found := indexMap[num]; found {
//             dist := int32(i - prevIndex)
//             if minDist == -1 || dist < minDist {
//                 minDist = dist
//             }
//         }
//         indexMap[num] = i
//     }
//     return minDist
// }

// Loop Execution for Each Element
// Initialization:

// minDist is set to -1.
// indexMap is an empty map of type map[int32]int.

// Loop through the Array:

// For each element num at index i:
// Check if num exists in indexMap.
// If it exists, calculate the distance dist between the current index i and the previous index stored in indexMap.
// Update minDist if dist is smaller than the current minDist.
// Update indexMap with the current index i for num.

// Example Execution
// arr := []int32{7, 1, 3, 4, 1, 7}

// Initial state:

// minDist = -1
// indexMap = {}
// Iteration 1 (i = 0, num = 7):

// indexMap = {7: 0}
// Iteration 2 (i = 1, num = 1):

// indexMap = {7: 0, 1: 1}
// Iteration 3 (i = 2, num = 3):

// indexMap = {7: 0, 1: 1, 3: 2}
// Iteration 4 (i = 3, num = 4):

// indexMap = {7: 0, 1: 1, 3: 2, 4: 3}
// Iteration 5 (i = 4, num = 1):

// prevIndex = 1
// dist = 4 - 1 = 3
// minDist = 3
// indexMap = {7: 0, 1: 4, 3: 2, 4: 3}
// Iteration 6 (i = 5, num = 7):

// prevIndex = 0
// dist = 5 - 0 = 5
// minDist = 3 (unchanged)
// indexMap = {7: 5, 1: 4, 3: 2, 4: 3}

// Time Complexity
// The time complexity is O(n), where n is the length of the array. This is because we are iterating through the array once.
// Space Complexity
// The space complexity is O(n) due to the storage used by indexMap.

// unit tests
// func TestMinimumDistances1(t *testing.T) {
//     tests := []struct {
//         arr  []int32
//         want int32
//     }{
//         {[]int32{7, 1, 3, 4, 1, 7}, 3},
//         {[]int32{1, 2, 3, 4, 10}, -1},
//         {[]int32{1, 1, 1, 1}, 1},
//         {[]int32{1, 2, 3, 4, 2, 1}, 1},
//         {[]int32{1, 2, 3, 4, 5}, -1},
//     }

//     for _, tt := range tests {
//         got := minimumDistances1(tt.arr)
//         if got != tt.want {
//             t.Errorf("minimumDistances1(%v) = %v; want %v", tt.arr, got, tt.want)
//         }
//     }
// }

// Use Cases
// Finding the closest pair of duplicate elements in an array.
// Analyzing data for repeated patterns.
// Optimizing search algorithms for repeated elements.
// Validating data integrity by checking for minimum distances between duplicates.
