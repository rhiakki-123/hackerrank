package main

import "fmt"

// Lily likes to play games with integers. She has created a new game where she determines the
// difference between a number and its reverse. For instance, given the number 12 , its reverse is 21.
// Their difference is 9. The number 120 reversed is 21, and their difference is 99.

// She decides to apply her game to decision making. She will look at a numbered range of days
// and will only go to a movie on a beautiful day

// Given a range of numbered days, [i...j] and a number k, determine the number of days in the
// range that are beautiful. Beautiful numbers are defined as numbers where |i - reverse(i)| is
// evenly divisible by k. If a days value is a beautiful number, it is a beautiful day. Return the
// number of beautiful days in the range.

/*
 * Complete the 'beautifulDays' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER i
 *  2. INTEGER j
 *  3. INTEGER k
 */

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	var count int32 = 0
	for n := i; n <= j; n++ {
		if abs(n-reverse(n))%k == 0 {
			count++
		}
	}
	return count
}

func beautifulDays1(i int32, j int32, k int32) int32 {
	// Write your code here
	var count int32 = 0
	for n := i; n <= j; n++ {
		if abs(n-reverse(n))%k == 0 {
			count++
		}
	}
	return count
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func reverse(n int32) int32 {
	var rev int32 = 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}

func test() {
	fmt.Println(beautifulDays(20, 23, 6))             // 2
	fmt.Println(beautifulDays1(13, 45, 3))            // 33
	fmt.Println(beautifulDays1(1, 2000000, 23047885)) // 2998

}

func main() {
	test()
}

// func beautifulDays(i int32, j int32, k int32) int32

// i: The starting day of the range.
// j: The ending day of the range.
// k: The divisor used to determine if a day is beautiful.
// Algorithm
// Initialize a variable count to 0 to keep track of the number of beautiful days.
// Iterate through each day from i to j using a for loop.
// For each day, calculate its reverse.
// Calculate the absolute difference between the day and its reverse.
// Check if the absolute difference is divisible by k.
// If the absolute difference is divisible by k, increment count.
// After the loop, return count.

// Start
//  |
//  v
// Initialize count to 0
//  |
//  v
// For day = i to j
//  |  |
//  |  v
//  |  Calculate reverse of day
//  |  |
//  |  v
//  |  Calculate absolute difference between day and reverse
//  |  |
//  |  v
//  |  Is absolute difference % k == 0?
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

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n * d), where n is the number of days in the range [i, j] and d is the number of digits in the day.
// This is because the function iterates through each day in the range and calculates the reverse of each day, which takes O(d) time.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the count variable and the loop variables.

// Real-Time Use Cases
// Date Calculations:

// Determining special dates based on specific criteria, such as palindrome dates or dates with specific properties.
// Event Planning:

// Identifying special days within a range for planning events or promotions.
// Data Analysis:

// Analyzing patterns in dates for various applications, such as finance or marketing.
// Gaming:

// Implementing date-based challenges or events in games.
