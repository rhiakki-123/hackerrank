package main

import (
	"fmt"
	"math/big"
)

// The factorial of the integer  n written as n! is the product of all positive integers less than or equal to n.
// !5 = 5 * 4 * 3 * 2 * 1 = 120

/*
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func extraLongFactorials(n int32) { //25
	fact := big.NewInt(1)
	// Loop from 1 to n
	for i := int32(1); i <= n; i++ {
		// Multiply fact by i
		fact.Mul(fact, big.NewInt(int64(i)))
	}
	fmt.Println(fact) // 15511210043330985984000000
}

func extraLongFactorials3(n int32) {
	// Write your code here
	var fact int64 = 1
	for i := int32(1); i <= n; i++ {
		fact *= int64(i)
	}
}

// write using recursion
func extraLongFactorials1(n int32) int64 {
	// Write your code here
	if n == 0 {
		return 1
	}
	return int64(n) * extraLongFactorials1(n-1)
}

// write using other approach
func extraLongFactorials2(n int32) int64 {
	// Write your code here
	var fact int64 = 1
	for i := int32(1); i <= n; i++ {
		fact *= int64(i)
	}
	return fact
}

func test() {
	fmt.Printf("Factorial %v\n", extraLongFactorials1(5))
	fmt.Printf("Factorial %v", extraLongFactorials2(25))
}

func main() {
	test()
}

// ////explain function func extraLongFactorials(n int32) with its logic, in details execution, algo, flowchart, timeand space complexity and use cases

// Explanation of extraLongFactorials Function
// The extraLongFactorials function calculates the factorial of a given number n using arbitrary-precision arithmetic provided by the math/big package in Go. This is necessary because the factorial of large numbers can exceed the storage capacity of standard integer types.

// Function Signature
// func extraLongFactorials(n int32)

// n: The input number for which the factorial is to be calculated.

// Logic
// Initialize a big integer fact with the value 1.
// Loop from 1 to n and multiply fact by each number in the loop.
// Print the result.

// Algorithm
// Initialize fact to 1 using big.NewInt(1).
// Loop from 1 to n:
// Multiply fact by the current loop variable i using fact.Mul(fact, big.NewInt(int64(i))).
// Print the result using fmt.Println(fact).

// Flowchart
// Start
//  |
//  v
// Initialize fact to 1
//  |
//  v
// For i = 1 to n
//  |  |
//  |  v
//  |  Multiply fact by i
//  |  |
//  v  Repeat loop
//  |
//  v
// Print fact
//  |
//  v
// End

// Detailed Execution
// Given the input:
// n := int32(25)

// Initialize variables:
// fact := big.NewInt(1)

// Loop from 1 to 25:
// First Iteration (i = 1)
// fact = fact * 1 = 1

// Second Iteration (i = 2)
// fact = fact * 2 = 2

// Third Iteration (i = 3)
// fact = fact * 3 = 6

// Twenty-Fifth Iteration (i = 25):

// fact = fact * 25 = 15511210043330985984000000

// Twenty-Fifth Iteration (i = 25):
// fmt.Println(fact) // Output: 15511210043330985984000000

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n), where n is the input number.
// This is because the function iterates from 1 to n to calculate the factorial.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the fact variable.
// Real-Time Use Cases
// Scientific Calculations:

// Calculating large factorials for scientific computations.
// Cryptography:

// Using large numbers in cryptographic algorithms.
// Combinatorics:

// Calculating combinations and permutations that involve large factorials.

// Summary
// Logic: Use the math/big package to handle arbitrary-precision arithmetic and calculate the factorial of a large number.
// Algorithm: Initialize a big integer, loop from 1 to n, multiply the big integer by each number in the loop, and print the result.
// Flowchart: Visual representation of the decision-making process for calculating the factorial of a large number.
// Time Complexity: O(n) because the function iterates from 1 to n to calculate the factorial.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Scientific calculations, cryptography, combinatorics.
