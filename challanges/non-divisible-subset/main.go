package main

// Given a set of distinct integers, print the size of a maximal subset of S where the sum of any 2 numbers in
// S is not evenly divisible by k.

// For example, the array S = [19, 10, 12, 10, 24, 25, 22] and k = 4. One of the arrays that can be created is
// S = [10, 12, 25]. Another is
// S = [19, 22, 24]. After testing all permutations, the maximum length solution array has 3 elements.

/*
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */

func nonDivisibleSubset(k int32, s []int32) int32 {
	freq := make([]int32, k)
	for _, num := range s {
		remainder := num % k
		freq[remainder]++
	}

	subsetSize := min(freq[0], 1)
	for i := int32(1); i <= k/2; i++ {
		if i == k-i {
			subsetSize += min(freq[i], 1)
		} else {
			subsetSize += max(freq[i], freq[k-i])
		}
	}
	return subsetSize
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// Code Implementation
func nonDivisibleSubset1(k int32, s []int32) int32 {
	// Write your code here
	rem := make([]int32, k)
	for _, v := range s {
		rem[v%k]++
	}
	count := int32(0)
	if rem[0] > 0 {
		count++
	}
	for i := int32(1); i <= k/2; i++ {
		if i != k-i {
			count += max1(rem[i], rem[k-i])
		} else {
			count++
		}
	}
	return count
}

func max1(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func test() {
	k := int32(4)
	s := []int32{19, 10, 12, 10, 24, 25, 22}
	println(nonDivisibleSubset(k, s)) // 3
}

func main() {
	test()
}

//explain function  func func nonDivisibleSubset(k int32, s []int32) int32   with its logic, in details execution, algo, flowchart, timeand space complexity and use cases

// Explanation of nonDivisibleSubset Function
// The nonDivisibleSubset function finds the maximum size of a subset of array s such that the sum of any two numbers in the subset is not divisible by k.

// Function Signature
// func nonDivisibleSubset(k int32, s []int32) int32

// k: The divisor.
// s: An array of integers.

// Algorithm
// Initialize a frequency array freq of size k to 0.
// Iterate through the array s and update the frequency array with the remainders.
// Initialize the subset size to include at most one element with a remainder of 0.
// For each remainder i from 1 to k/2:
// If i is equal to k - i, include at most one element with this remainder.
// Otherwise, include the maximum count of elements with remainder i or k - i.
// Return the subset size.

// flowchart
// Start
//  |
//  v
// Initialize freq array of size k to 0
//  |
//  v
// For each element in s
//  |  |
//  |  v
//  |  Calculate remainder and update freq array
//  |  |
//  v  Repeat loop
//  |
//  v
// Initialize subset size to min(freq[0], 1)
//  |
//  v
// For i = 1 to k/2
//  |  |
//  |  v
//  |  Is i == k - i?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | Include min(freq[i], 1)  Include max(freq[i], freq[k - i])
//  |  |
//  v  Repeat loop
//  |
//  v
// Return subset size
//  |
//  v
// End

// Detailed Execution
// Given the inputs:
// k := int32(3)
// s := []int32{1, 7, 2, 4}

// Initialize variables:
// freq := make([]int32, k) // freq = [0, 0, 0]

// Calculate remainders and update frequency array:
// First Element (1)
// remainder = 1 % 3 = 1
// freq[1]++
// freq = [0, 1, 0]

// Second Element (7)
// remainder = 7 % 3 = 1
// freq[1]++
// freq = [0, 2, 0]

// Third Element (2)
// remainder = 2 % 3 = 2
// freq[2]++
// freq = [0, 2, 1]

// Fourth Element (4)
// remainder = 4 % 3 = 1
// freq[1]++
// freq = [0, 3, 1]

// Initialize subset size to min(freq[0], 1) = min(0, 1) = 0
// subsetSize := min(freq[0], 1) // subsetSize = 0

// Iterate through remainders:
// First Iteration (i = 1)
// subsetSize += max(freq[1], freq[3 - 1]) // subsetSize += max(3, 1) = 3
// subsetSize = 3

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n + k), where n is the length of the array s and k is the divisor.
// This is because the function iterates through the array s to calculate remainders and then iterates through the frequency array.
// Space Complexity
// The space complexity of the function is O(k).
// This is because the function uses an additional frequency array of size k.

// Real-Time Use Cases
// Data Analysis:

// Finding subsets of data that meet specific criteria, such as non-divisibility by a given number.
// Cryptography:

// Analyzing subsets of numbers for cryptographic algorithms that require non-divisibility properties.
// Game Development:

// Implementing game mechanics that involve selecting subsets of elements based on specific rules.

// Summary
// Logic: Calculate the remainder of each element in s when divided by k, use a frequency array to count the occurrences of each remainder, and determine the maximum size of a subset where the sum of any two numbers is not divisible by k.
// Algorithm: Initialize a frequency array, update the frequency array with remainders, initialize the subset size, iterate through remainders, and return the subset size.
// Flowchart: Visual representation of the decision-making process for finding the maximum size of a non-divisible subset.
// Time Complexity: O(n + k) because the function iterates through the array s and the frequency array.
// Space Complexity: O(k) because the function uses an additional frequency array of size k.
// Use Cases: Data analysis, cryptography, game development
