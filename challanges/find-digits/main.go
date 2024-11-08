package main

import (
	"fmt"
	"strconv"
)

// An integer d is a divisor of an integer n if the remainder of n / d = 0.
// Given an integer, for each digit that makes up the integer determine whether it is a divisor.
// Count the number of divisors occurring within the integer.

//n = 124
// The number 124 is broken into three digits: 1, 2, and 4. When 124 is divided evenly into 124 so
//return 3.

func findDigits(n int32) int32 {
	// Write your code here
	var count int32 = 0
	num := n
	for num > 0 {
		digit := num % 10
		if digit != 0 && n%digit == 0 {
			count++
		}
		num /= 10
	}
	return count
}

// write the same code using a different approach
func findDigits1(n int32) int32 {
	// Write your code here
	count := int32(0)
	str := strconv.Itoa(int(n))
	for _, digit := range str {
		d := int32(digit - '0')
		if d != 0 && n%d == 0 {
			count++
		}
	}
	return count
}

func test() {
	fmt.Println(findDigits1(124)) // 3
	fmt.Println(findDigits1(111)) // 3
	fmt.Println(findDigits1(10))  // 1
}

func main() {
	test()
}

// Explanation of findDigits1 Function
// The findDigits1 function determines how many digits in a given number n are divisors of n. A digit is considered a divisor if n is divisible by that digit without leaving a remainder.

// Function Signature

// func findDigits1(n int32) int32

// n: The integer for which the divisors need to be counted.

// Logic
// Initialize a variable count to 0 to keep track of the number of digits that are divisors of n.
// Convert the number n to a string to iterate through each digit.
// For each digit in the string representation of n:
// 	Convert the digit back to an integer.
// 	Check if the digit is not zero and if n is divisible by the digit.
// 	If both conditions are met, increment count.
// Return the value of count.

// Algorithm
// Initialize count to 0.
// Convert n to a string to iterate through each digit.
// For each digit in the string representation of n:
// 	Convert the digit back to an integer.
// 	Check if the digit is not zero and if n is divisible by the digit.
// 	If both conditions are met, increment count.
// Return count.

// Start
//  |
//  v
// Initialize count to 0
//  |
//  v
// Convert n to string
//  |
//  v
// For each digit in string representation of n
//  |  |
//  |  v
//  |  Convert digit to integer
//  |  |
//  |  v
//  |  Is digit != 0 and n % digit == 0?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | Increment count
//  |  |
//  v  Repeat loop
//  |
//  v
// Return count
//  |
//  v
// End

// Detailed Execution
// Given the input:
// n := int32(1012)

// Initialize variables:
// count = 0

// Convert n to string:
// str := strconv.Itoa(int(n)) // "1012"

// Iterate through each digit in the string:
// First Iteration (digit = '1'):
// Convert digit to integer
// digit := int32(digit - '0') // 1
// Check if digit is not zero and n is divisible by digit
// digit != 0 && n % digit == 0 // 1 != 0 && 1012 % 1 == 0 // true
// Increment count
// count++ // 1

// Second Iteration (digit = '0'):
// Convert digit to integer
// digit := int32(digit - '0') // 0
// Check if digit is not zero and n is divisible by digit
// digit != 0 && n % digit == 0 // 0 != 0 && 1012 % 0 == 0 // false
// count remains 1.

// Third Iteration (digit = '1'):
// Convert digit to integer
// digit := int32(digit - '0') // 1
// Check if digit is not zero and n is divisible by digit
// digit != 0 && n % digit == 0 // 1 != 0 && 1012 % 1 == 0 // true
// Increment count
// count++ // 2

// Fourth Iteration (digit = '2'):
// Convert digit to integer
// digit := int32(digit - '0') // 2
// Check if digit is not zero and n is divisible by digit
// digit != 0 && n % digit == 0 // 2 != 0 && 1012 % 2 == 0 // true
// Increment count
// count++ // 3

// Return count
// count // 3
// End

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(d), where d is the number of digits in the number n.
// This is because the function iterates through each digit in the string representation of n.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the count variable and the loop variables.

// Real-Time Use Cases
// Number Analysis:

// Analyzing the properties of numbers, such as finding digits that are divisors of the number.
// Digit-Based Calculations:

// Performing calculations based on the digits of a number.
// Mathematical Puzzles:

// Solving puzzles and problems that involve digit-based properties of numbers.

// Summary
// Logic: Iterate through each digit of the number n, check if the digit is a divisor of n, and count the number of such digits.
// Algorithm: Convert n to a string, iterate through each digit, check if the digit is a divisor of n, and return the count.
// Flowchart: Visual representation of the decision-making process for counting the digits that are divisors of n.
// Time Complexity: O(d) because the function iterates through each digit in the string representation of n.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Number analysis, digit-based calculations, mathematical puzzles.
