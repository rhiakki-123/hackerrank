package main

import (
	"fmt"
	"strconv"
)

func kaprekarNumbers(p int32, q int32) {
	found := false
	for i := p; i <= q; i++ {
		square := int64(i) * int64(i)
		squareStr := strconv.FormatInt(square, 10)
		d := len(strconv.FormatInt(int64(i), 10))
		leftStr := "0"
		if len(squareStr) > d {
			leftStr = squareStr[:len(squareStr)-d]
		}
		rightStr := squareStr[len(squareStr)-d:]

		left, _ := strconv.Atoi(leftStr)
		right, _ := strconv.Atoi(rightStr)

		if left+right == int(i) {
			fmt.Printf("%d ", i)
			found = true
		}
	}
	if !found {
		fmt.Println("INVALID RANGE")
	} else {
		fmt.Println()
	}
}

func main() {
	p := int32(400)
	q := int32(700)
	kaprekarNumbers(p, q)
}

// Explain function kaprekarNumbers(p int32, q int32)
// with its in details execution, algo, flowchart, time
// and space complexity, write its optimized version, unit tests and use cases

// To understand why the test case kaprekarNumbers(p, q) is failing for the input p := int32(400) and q := int32(700), lets first review the logic of the kaprekarNumbers function and then identify potential issues.

// Explanation of kaprekarNumbers Function
// A Kaprekar number for a given base is a non-negative integer, the representation of whose square can be split into two parts that add up to the original number again. For example, 45 is a Kaprekar number because 45² = 2025 and 20 + 25 = 45.

// Function Signature

// func kaprekarNumbers(p int32, q int32)

// Parameters:
// p: The starting range value.
// q: The ending range value.

// Logic
// Iterate through all numbers from p to q.
// For each number, calculate its square.
// Split the square into two parts.
// Check if the sum of the two parts equals the original number.
// Print the Kaprekar numbers found in the range.

// Algorithm
// Iterate through the Range:

// Loop through all numbers from p to q.
// Calculate the Square:

// Calculate the square of the current number.
// Convert to String:

// Convert the square to a string to facilitate splitting.
// Split the Square:

// Split the string representation of the square into two parts.
// Convert Parts to Integers:

// Convert the two parts back to integers.
// Check the Kaprekar Condition:

// Check if the sum of the two parts equals the original number.
// Print the Kaprekar Numbers:

// Print the numbers that satisfy the Kaprekar condition.

// Detailed Execution
// Step-by-Step Execution for p := int32(400) and q := int32(700)

// Iterate through the Range:

// Loop through all numbers from 400 to 700.
// Calculate the Square:

// For i = 400, square = 400 * 400 = 160000.
// Convert to String:

// squareStr = "160000".
// Split the Square:

// d = 3 (number of digits in 400).
// leftStr = "160", rightStr = "000".
// Convert Parts to Integers:

// left = 160, right = 0.
// Check the Kaprekar Condition:

// left + right = 160 + 0 = 160 (not equal to 400).
// Continue Iteration:

// Repeat the above steps for all numbers from 400 to 700.

// Summary
// Logic: Iterate through the range, calculate the square, split the square, and check the Kaprekar condition.
// Algorithm: Loop through the range, calculate the square, convert to string, split, convert back to integers, and check the condition.
// Time Complexity: O(n) because the function iterates through the range and performs constant-time operations for each number.
// Space Complexity: O(1) because the function uses a constant amount of extra space.s
