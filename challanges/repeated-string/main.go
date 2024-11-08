package main

import (
	"fmt"
)

// There is a string, s , of lowercase English letters that is repeated infinitely many times. Given an integer, n, find and print
// the number of letter as in the first n letters of the infinite string.

// Example
// s = 'abcac'
// n = 10
// The substring we consider is  abcacabcac, the first 10 characters of the infinite string.
// There are 4 occurrences of a in the substring.

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

// optimized version of repeatedString function
func repeatedString(s string, n int64) int64 {

	if len(s) == 0 {
		return 0
	}

	lenS := int64(len(s))
	fullRepetitions := n / lenS
	remainingChars := n % lenS

	countA := int64(0)
	for _, char := range s {
		if char == 'a' {
			countA++
		}
	}

	totalCount := countA * fullRepetitions

	for i := int64(0); i < remainingChars; i++ {
		if s[i] == 'a' {
			totalCount++
		}
	}

	return totalCount
}

// Code Implementation
// func repeatedString1(s string, n int64) int64 {
// 	// Write your code here
// 	var count int64
// 	for _, v := range s {
// 		if v == 'a' {
// 			count++
// 		}
// 	}
// 	count *= n / int64(len(s))
// 	for i := int64(0); i < n%int64(len(s)); i++ {
// 		if s[i] == 'a' {
// 			count++
// 		}
// 	}
// 	return count
// }

// func repeatedString2(s string, n int64) int64 {
// 	// Write your code here
// 	var count int64
// 	for i := int64(0); i < n; i++ {
// 		if s[i%int64(len(s))] == 'a' {
// 			count++
// 		}
// 	}
// 	return count
// }

func testRepeatedString() {
	s := "abcac"
	n := int64(10)
	fmt.Println(repeatedString(s, n)) // 4	// abcacabcac
	s = "a"
	n = int64(1000000000000)
	fmt.Println(repeatedString(s, n)) // 1000000000000
	s = "abc"
	n = int64(7)
	fmt.Println(repeatedString(s, n)) // 7	// abaabaabaa

}

func main() {
	testRepeatedString()
	// TestRepeatedString() // This should be run using `go test`

}

//explain function  func repeatedString(s string, n int64) int64
// with its in details execution, algo, flowchart, time
// and space complexity, write its optimized version, unit tests and use cases

// Explanation of repeatedString Function
// The repeatedString function calculates the number of occurrences of the character 'a' in the first
// n characters of an infinitely repeated string s.

// Function Signature
// func repeatedString(s string, n int64) int64

// Function Parameters
// s: A string representing the infinite string.
// n: An integer representing the number of characters to consider.

// Function Logic
// Calculate the number of times the string s can be fully repeated within the first n characters.
// Calculate the number of remaining characters after the full repetitions.
// Count the occurrences of 'a' in the string s.
// Multiply the count of 'a' by the number of full repetitions.
// Count the occurrences of 'a' in the remaining characters.
// Add the counts from steps 4 and 5 to get the total number of occurrences of 'a' in the first n characters.

// Algorithm
// Calculate the length of the string s.
// Calculate the number of full repetitions of s within n characters: fullRepetitions = n / int64(len(s)).
// Calculate the number of remaining characters after the full repetitions: remainingChars = n % int64(len(s)).
// Count the occurrences of 'a' in the string s.
// Multiply the count of 'a' by fullRepetitions to get the total count of 'a' in the full repetitions.
// Count the occurrences of 'a' in the first remainingChars characters of s.
// Add the counts from steps 5 and 6 to get the total number of occurrences of 'a' in the first n characters.
// Return the total count.

// Flowchart
// Start
//  |
//  v
// Calculate length of s
//  |
//  v
// Calculate fullRepetitions as n / len(s)
//  |
//  v
// Calculate remainingChars as n % len(s)
//  |
//  v
// Count occurrences of 'a' in s
//  |
//  v
// Multiply count of 'a' by fullRepetitions
//  |
//  v
// Count occurrences of 'a' in first remainingChars characters of s
//  |
//  v
// Add counts from fullRepetitions and remainingChars
//  |
//  v
// Return total count
//  |
//  v
// End

// Detailed Execution
// Given the inputs:
// s := "abcac"
// n := int64(10)

// Calculate the length of s:
// lenS := int64(len(s)) // lenS = 5

// Calculate the number of full repetitions of s within n characters:
// fullRepetitions := n / lenS // fullRepetitions = 10 / 5 = 2

// Calculate the number of remaining characters after the full repetitions:
// remainingChars := n % lenS // remainingChars = 10 % 5 = 0

// Count the occurrences of 'a' in s:
// countA := int64(0)
// for _, char := range s {
//     if char == 'a' {
//         countA++
//     }
// }
// // countA = 2	// 'a' occurs twice in s = "abcac"

// Multiply the count of 'a' by fullRepetitions:
// totalCount := countA * fullRepetitions // totalCount = 2 * 2 = 4

// Count the occurrences of 'a' in the first remainingChars characters of s:
// for i := int64(0); i < remainingChars; i++ {
// 	if s[i] == 'a' {
// 		totalCount++
// 	}
// }
//  //remainingChars = 0, so no additional 'a' characters to count

//  return the total count:
//  return totalCount // 4

//  Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(len(s)).
// This is because the function iterates through the string s to count the occurrences of 'a'.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the variables lenS, fullRepetitions, remainingChars, countA, and totalCount.

// Optimized Version of repeatedString Function
// package main

// import "fmt"

// //Optimized version of repeatedString function
// func repeatedString(s string, n int64) int64 {
//     lenS := int64(len(s))
//     fullRepetitions := n / lenS
//     remainingChars := n % lenS

//     countA := int64(0)
//     for _, char := range s {
//         if char == 'a' {
//             countA++
//         }
//     }

//     totalCount := countA * fullRepetitions

//     for i := int64(0); i < remainingChars; i++ {
//         if s[i] == 'a' {
//             totalCount++
//         }
//     }

//     return totalCount
// }

// func test() {
//     s := "abcac"
//     n := int64(10)
//     fmt.Println(repeatedString(s, n)) // Output: 4
// }

// func main() {
//     test()
// }

// // Unit tests for repeatedString function
// func TestRepeatedString(t *testing.T) {
//     tests := []struct {
//         s        string
//         n        int64
//         expected int64
//     }{
//         {"abcac", 10, 4},
//         {"a", 1000000000000, 1000000000000},
//         {"abc", 7, 2},
//         {"abca", 10, 5},
//         {"", 10, 0},
//     }

//     for _, test := range tests {
//         result := repeatedString(test.s, test.n)
//         if result != test.expected {
//             t.Errorf("For input s=%s and n=%d, expected %d but got %d", test.s, test.n, test.expected, result)
//         }
//     }
// }

// Summary
// Logic: Calculate the number of occurrences of 'a' in the first n characters of an infinitely repeated string s by counting the occurrences in full repetitions and remaining characters.
// Algorithm: Calculate the length of s, the number of full repetitions, the number of remaining characters, count the occurrences of 'a' in s, multiply by the number of full repetitions, count the occurrences in the remaining characters, and return the total count.
// Flowchart: Visual representation of the decision-making process for calculating the number of occurrences of 'a' in the first n characters of an infinitely repeated string s.
// Time Complexity: O(len(s)) because the function iterates through the string s to count the occurrences of 'a'.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Text analysis, data processing, pattern matchin
