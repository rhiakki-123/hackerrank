package main

import "fmt"

// Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers.

// Example arr = [1,3,5,7,9]

// The minimum sum is  1+3+5+7 = 16 and the maximum sum is 3+5+7+9=24 . The function prints

// 16 24

func miniMaxSum(arr []int32) {
	// Write your code here
	var min, max, sum int64
	min = 1000000000
	max = 0
	for _, v := range arr {
		sum += int64(v)
		if int64(v) < min {
			min = int64(v)
		}
		if int64(v) > max {
			max = int64(v)
		}
	}
	fmt.Println(sum-max, sum-min)
}

func miniMaxSum1(arr []int32) {
	// Write your code here
	var min, max, sum int64
	min = 1000000000
	max = 0
	for i := 0; i < len(arr); i++ {
		sum += int64(arr[i])
		if int64(arr[i]) < min {
			min = int64(arr[i])
		}
		if int64(arr[i]) > max {
			max = int64(arr[i])
		}
	}
	fmt.Println(sum-max, sum-min)
}

func test() {
	//miniMaxSum([]int32{1, 2, 3, 4, 5})
	miniMaxSum([]int32{5, 5, 5, 4, 5, 2, 6, 8, 9})
}

func main() {
	test()
}

// Explanation of for Loop Execution
// The miniMaxSum function calculates the minimum and maximum sums of 4 out of 5 elements in an array. Let's break down the execution of the for loop step-by-step.

// Given the array:
// arr := []int32{1, 2, 3, 4, 5}

// Initial State:
// min = 1000000000
// max = 0
// sum = 0
// Loop Execution:
// First Iteration (v = 1):

// sum += int64(v) → sum = 0 + 1 = 1
// if int64(v) < min → 1 < 1000000000 → Yes → min = 1
// if int64(v) > max → 1 > 0 → Yes → max = 1
// Second Iteration (v = 2):

// sum += int64(v) → sum = 1 + 2 = 3
// if int64(v) < min → 2 < 1 → No
// if int64(v) > max → 2 > 1 → Yes → max = 2
// Third Iteration (v = 3):

// sum += int64(v) → sum = 3 + 3 = 6
// if int64(v) < min → 3 < 1 → No
// if int64(v) > max → 3 > 2 → Yes → max = 3
// Fourth Iteration (v = 4):

// sum += int64(v) → sum = 6 + 4 = 10
// if int64(v) < min → 4 < 1 → No
// if int64(v) > max → 4 > 3 → Yes → max = 4
// Fifth Iteration (v = 5):

// sum += int64(v) → sum = 10 + 5 = 15
// if int64(v) < min → 5 < 1 → No
// if int64(v) > max → 5 > 4 → Yes → max = 5
// Final State:
// min = 1
// max = 5
// sum = 15
// Calculation of the Result:
// sum - max → 15 - 5 = 10
// sum - min → 15 - 1 = 14
// Output:
// 10 14

// Flowchart
// Start
//  |
//  v
// Initialize min to 1000000000, max to 0, sum to 0
//  |
//  v
// For each element v in arr
//  |  |
//  |  v
//  |  sum += int64(v)
//  |  |
//  |  v
//  |  Is int64(v) < min?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | min = int64(v)
//  |  |
//  |  v
//  |  Is int64(v) > max?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | max = int64(v)
//  |  |
//  |  v
//  | Repeat loop
//  |
//  v
// Calculate sum - max
//  |
//  v
// Calculate sum - min
//  |
//  v
// Print results
//  |
//  v
// End
