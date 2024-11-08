package main

import "fmt"

// You have two strings of lowercase English letters. You can perform two types of operations on the first string:

// Append a lowercase English letter to the end of the string.
// Delete the last character of the string. Performing this operation on an empty string results in an empty string.
// Given an integer k and two strings, s and t, determine whether or not you can convert s to t by performing exactly
// k of the above operations on s. If its possible, print Yes; otherwise, print No.

// For example, strings s = [a, b, c] and t = [d, e, f].  k = 6

// To convert s to t, we first delete all of the characters in 3 operations.
// Then we perform k = 3 more operations to append d, e, and f.
// on the 6th operation, we have the string t. return Yes.

/*
 * Complete the 'appendAndDelete' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING t
 *  3. INTEGER k
 */

func appendAndDelete(s string, t string, k int32) string {
	// Write your code here
	sLen := len(s)
	tLen := len(t)
	if sLen+tLen <= int(k) {
		return "Yes"
	}
	i := 0
	for i < sLen && i < tLen && s[i] == t[i] {
		i++
	}
	ops := int32(sLen + tLen - 2*i)
	if ops <= k && (k-ops)%2 == 0 {
		return "Yes"
	}
	return "No"
}

func test() {
	s := "hackerhappy"
	t := "hackerrank"
	k := int32(9)
	fmt.Println(appendAndDelete(s, t, k)) // Yes

	s = "aba"
	t = "aba"
	k = int32(7)
	fmt.Println(appendAndDelete(s, t, k)) // Yes

	s = "ashley"
	t = "ash"
	k = int32(2)
	fmt.Println(appendAndDelete(s, t, k)) // No
}

func main() {
	test()
}

//explain function func func appendAndDelete(s string, t string, k int32) string with its logic, in details execution, algo, flowchart, timeand space complexity and use cases

// Explanation of appendAndDelete Function
// The appendAndDelete function determines whether it is possible to convert string s into string t using exactly k operations. The operations allowed are appending a character to the end of the string or deleting the last character of the string.

// Function Signature
// func appendAndDelete(s string, t string, k int32) string

// s: The initial string.
// t: The target string.
// k: The number of operations allowed.

// Logic
// Find the common prefix length of s and t.
// Calculate the number of operations needed to delete the non-common part of s and append the non-common part of t.
// Check if the total number of operations is less than or equal to k and if the difference between k and the total number of operations is even.
// If the conditions are met, return "Yes"; otherwise, return "No".

// Algorithm
// Initialize a variable commonLength to 0.
// Iterate through the characters of s and t to find the common prefix length.
// Calculate the number of operations needed to delete the non-common part of s and append the non-common part of t:
// totalOperations = len(s) - commonLength + len(t) - commonLength
// Check if totalOperations is less than or equal to k and if (k - totalOperations) % 2 == 0 or if len(s) + len(t) <= k.
// If the conditions are met, return "Yes"; otherwise, return "No".

// Flowchart
// Start
//  |
//  v
// Initialize commonLength to 0
//  |
//  v
// For i = 0 to min(len(s), len(t)) - 1
//  |  |
//  |  v
//  |  Is s[i] == t[i]?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | Increment commonLength
//  |  |
//  v  Repeat loop
//  |
//  v
// Calculate totalOperations as (len(s) - commonLength) + (len(t) - commonLength)
//  |
//  v
// Is totalOperations <= k and (k - totalOperations) % 2 == 0 or len(s) + len(t) <= k?
//  / \
// Yes  No
//  |    |
//  v    v
// Return "Yes"  Return "No"
//  |
//  v
// End

// Detailed Execution
// Given the input:
// s := "hackerhappy"
// t := "hackerrank"
// k := int32(9)

// Initialize commonLength to 0.
// commonLength := 0

// Find common prefix length:
// First Iteration (i = 0):
// s[0] == t[0] // "h" == "h"  //true  commonLength = 1

// Second Iteration (i = 1):
// s[1] == t[1] // "a" == "a"  //true  commonLength = 2

// Third Iteration (i = 2):
// s[2] == t[2] // "c" == "c"  //true  commonLength = 3

// Fourth Iteration (i = 3):
// s[3] == t[3] // "k" == "k"  //true  commonLength = 4

// Fifth Iteration (i = 4):
// s[4] == t[4] // "e" == "e"  //true  commonLength = 5

// Sixth Iteration (i = 5):
// s[5] == t[5] // "r" == "r"  //true  commonLength = 6

// Seventh Iteration (i = 6):
// s[6] == t[6] // "h" == "r"  //false  commonLength remains 6

// Calculate totalOperations:

// totalOperations = (len(s) - commonLength) + (len(t) - commonLength)
// totalOperations = (11 - 6) + (10 - 6)
// totalOperations = 5 + 4
// totalOperations = 9

// Check conditions:
// totalOperations <= k and (k - totalOperations) % 2 == 0 or len(s) + len(t) <= k
// 9 <= 9 and (9 - 9) % 2 == 0 or 11 + 10 <= 9
// true and true or false
// true

// Return "Yes"

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(min(len(s), len(t))).
// This is because the function iterates through the characters of s and t to find the common prefix length.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the commonLength and totalOperations variables.

// Real-Time Use Cases
// Text Editing:

// Determining if a series of text edits can transform one string into another within a given number of operations.
// Version Control:

// Checking if a series of changes can convert one version of a file into another within a specified number of operations.
// String Manipulation:

// Validating if a sequence of append and delete operations can achieve a desired string transformation.

// GitHub Copilot
// Explanation of appendAndDelete Function
// The appendAndDelete function determines whether it is possible to convert string s into string t using exactly k operations. The operations allowed are appending a character to the end of the string or deleting the last character of the string.

// Function Signature
// s: The initial string.
// t: The target string.
// k: The number of operations allowed.
// Logic
// Find the common prefix length of s and t.
// Calculate the number of operations needed to delete the non-common part of s and append the non-common part of t.
// Check if the total number of operations is less than or equal to k and if the difference between k and the total number of operations is even.
// If the conditions are met, return "Yes"; otherwise, return "No".
// Algorithm
// Initialize a variable commonLength to 0.
// Iterate through the characters of s and t to find the common prefix length.
// Calculate the number of operations needed to delete the non-common part of s and append the non-common part of t:
// totalOperations = len(s) - commonLength + len(t) - commonLength
// Check if totalOperations is less than or equal to k and if (k - totalOperations) % 2 == 0 or if len(s) + len(t) <= k.
// If the conditions are met, return "Yes"; otherwise, return "No".
// Flowchart
// Detailed Execution
// Given the inputs:

// 1 vulnerability
// Initialize variables:

// Find common prefix length:

// First Iteration (i = 0):
// Second Iteration (i = 1):
// Third Iteration (i = 2):
// Fourth Iteration (i = 3):
// Fifth Iteration (i = 4):
// Sixth Iteration (i = 5):
// Seventh Iteration (i = 6):
// Calculate total operations:

// Check conditions:

// Return result:

// Final State:
// commonLength = 6
// totalOperations = 9
// Output:
// "Yes"
// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(min(len(s), len(t))).
// This is because the function iterates through the characters of s and t to find the common prefix length.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the commonLength and totalOperations variables.
// Real-Time Use Cases
// Text Editing:

// Determining if a series of text edits can transform one string into another within a given number of operations.
// Version Control:

// Checking if a series of changes can convert one version of a file into another within a specified number of operations.
// String Manipulation:

// Validating if a sequence of append and delete operations can achieve a desired string transformation.
// Code Implementation
// Summary
// Logic: Find the common prefix length of s and t, calculate the number of operations needed to delete the non-common part of s and append the non-common part of t, and check if the total number of operations is less than or equal to k and if the difference between k and the total number of operations is even.
// Algorithm: Initialize commonLength, find the common prefix length, calculate totalOperations, check the conditions, and return the result.
// Flowchart: Visual representation of the decision-making process for determining if it is possible to convert s into t using exactly k operations.
// Time Complexity: O(min(len(s), len(t))) because the function iterates through the characters of s and t to find the common prefix length.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Text editing, version control, string manipulation.
