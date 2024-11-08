package main

import (
	"fmt"
)

// Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero.
// Print the decimal value of each fraction on a new line with  places after the decimal.

// Note: This challenge introduces precision problems. The test cases are scaled to six decimal places,
//  though answers with absolute error of up to  10^-4 are acceptable.

// Example
// arr = [1,1,0,-1,-1]
// There are  elements, two positive, two negative and one zero. Their ratios are ,  and . Results are printed as:
// 2/5=0.400000, 2/5=0.40000,1/5=0.20000
// 0.400000
// 0.400000
// 0.200000

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	var positive, negative, zero int32
	for _, v := range arr {
		if v > 0 {
			positive++
		} else if v < 0 {
			negative++
		} else {
			zero++
		}
	}
	fmt.Printf("%.6f\n", float64(positive)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(negative)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(zero)/float64(len(arr)))
}

func plusMinus1(arr []int32) {
	// Write your code here
	var positive, negative, zero int32
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positive++
		} else if arr[i] < 0 {
			negative++
		} else {
			zero++
		}
	}
	fmt.Printf("%.6f\n", float64(positive)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(negative)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(zero)/float64(len(arr)))
}

func test() {
	plusMinus([]int32{1, 1, 0, -1, -1})
	plusMinus1([]int32{1, 1, 0, -1, -1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, -9, -6, 1, 1, 1, 1})
}

func main() {
	test()
}

// Flowchart

// Start
//  |
//  v
// Initialize positive, negative, zero to 0
//  |
//  v
// For each element v in arr
//  |  |
//  |  v
//  |  Is v > 0?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | Increment positive
//  |  |    |
//  |  v    Is v < 0?
//  |  |    / \
//  |  |  Yes  No
//  |  |   |    |
//  |  |   v    v
//  |  | Increment negative
//  |  |   |    |
//  |  |   v    Increment zero
//  |  |   |
//  |  v   v
//  | Repeat loop
//  |
//  v
// Calculate positive ratio: positive / len(arr)
//  |
//  v
// Calculate negative ratio: negative / len(arr)
//  |
//  v
// Calculate zero ratio: zero / len(arr)
//  |
//  v
// Print positive ratio
//  |
//  v
// Print negative ratio
//  |
//  v
// Print zero ratio
//  |
//  v
// End
