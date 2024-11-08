package main

import "fmt"

// Given a sequence of n integers , p(1), p(2), p(3), p(n) where each element is distinct and satisfies 1 <= p(x) <= n.
// For each x where 1 <= x <= n, find any integer y such that p(p(y)) = x and print the value of y on a new line.

// For example, assume the sequence p = [5, 2, 1, 3, 4]. Each value of x between 1 and 5, the length of the sequence, is analyzed as follows:

// x = 1: p[3], p[4] = 3, p[p[4]] = 1
// x = 2: p[2], p[2] = 2, p[p[2]] = 2
// x = 3: p[4], p[5] = 4, p[p[5]] = 3
// x = 4: p[5], p[1] = 5, p[p[1]] = 4
// x = 5: p[1], p[3] = 1, p[p[3]] = 5

// The values for y are [4, 2, 5, 1, 3].

/*
 * Complete the 'permutationEquation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY p as parameter.
 */

func permutationEquation(p []int32) []int32 {
	// Write your code here
	positionMap := map[int32]int32{}
	for i, value := range p {
		positionMap[value] = int32(i + 1)
	}
	result := []int32{}
	for x := int32(1); x <= int32(len(p)); x++ {
		y := positionMap[positionMap[x]]
		result = append(result, y)
	}
	return result
}

func permutationEquation1(p []int32) []int32 {
	// Write your code here
	var result []int32
	for i := int32(1); i <= int32(len(p)); i++ {
		result = append(result, int32(p[int32(p[int32(p[i-1])-1])-1]))
	}
	return result
}

func test() {
	fmt.Println(permutationEquation([]int32{2, 3, 1}))        // [2, 3, 1]
	fmt.Println(permutationEquation1([]int32{4, 3, 5, 1, 2})) // [1, 3, 5, 4, 2]
}

func main() {
	test()
}

// Explanation of permutationEquation Function
// The permutationEquation function solves a problem where you are given a sequence of integers, and you need to find the value of y for each x such that p(p(y)) = x. The function returns an array of integers representing the values of y for each x from 1 to n.

// Function Signature
// func permutationEquation(p []int32) []int32

// Logic
// Create a map to store the positions of each value in the array p.
// Iterate through each value x from 1 to n.
// For each x, find the position y such that p(p(y)) = x.
// Append the value of y to the result array.
// Algorithm
// Create a map positionMap to store the positions of each value in the array p.
// Iterate through the array p and populate positionMap with the positions of each value.
// Initialize an empty result array.
// Iterate through each value x from 1 to n:
// Find the position y such that p(p(y)) = x using the positionMap.
// Append the value of y to the result array.
// Return the result array.

// Start
//  |
//  v
// Create positionMap
//  |
//  v
// For i = 0 to len(p)-1
//  |  |
//  |  v
//  |  positionMap[p[i]] = i + 1
//  |  |
//  v  Repeat loop
//  |
//  v
// Initialize result as empty array
//  |
//  v
// For x = 1 to len(p)
//  |  |
//  |  v
//  |  y = positionMap[positionMap[x]]
//  |  |
//  |  v
//  |  Append y to result
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
// The time complexity of the function is O(n), where n is the length of the array p.
// This is because the function iterates through the array p twice: once to populate the positionMap and once to calculate the result array.
// Space Complexity
// The space complexity of the function is O(n).
// This is because the function uses an additional map positionMap to store the positions of each value in the array p.
// Real-Time Use Cases
// Index Mapping:

// Mapping indices in one array to another array based on specific rules.
// Permutation Problems:

// Solving problems related to permutations and their properties.
// Data Rearrangement:

// Rearranging data based on specific index mappings.

// Detailed Execution
// Given the input:

// p := []int32{5, 2, 1, 3, 4}

// Create positionMap:
// positionMap := map[int32]int32{}

// Populate positionMap:

// positionMap[5] = 1
// positionMap[2] = 2
// positionMap[1] = 3
// positionMap[3] = 4
// positionMap[4] = 5

// Initialize result:result := []int32{}

// First Iteration (x = 1):

// y = positionMap[positionMap[1]] = positionMap[3] = 4
// Append y to result:
// result = append(result, 4) = [4]

// Second Iteration (x = 2):

// y = positionMap[positionMap[2]] = positionMap[2] = 2
// Append y to result
// result = append(result, 2) = [4, 2]

// Third Iteration (x = 3):
// y = positionMap[positionMap[3]] = positionMap[4] = 5
// Append y to result
// result = append(result, 5) = [4, 2, 5]

// Fourth Iteration (x = 4):
// y = positionMap[positionMap[4]] = positionMap[5] = 1
// Append y to result
// result = append(result, 1) = [4, 2, 5, 1]

// Fifth Iteration (x = 5):
// y = positionMap[positionMap[5]] = positionMap[1] = 3
// Append y to result
// result = append(result, 3) = [4, 2, 5, 1, 3]

// Return result: [4, 2, 5, 1, 3]
