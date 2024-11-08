package main

import "fmt"

// HackerLand Enterprise is adopting a new viral advertising strategy. When they launch a new
// product, they advertise it to exactly 5 people on social media. On the first day, half of those
// 5 people (i.e., floor(5/2) = 2) like the advertisement and each shares it with 3 of their friends.
// At the beginning of the second day, floor(5/2) * 3 = 2 * 3 = 6 people receive the advertisement.
// Each day, floor(recipients/2) of the recipients like the advertisement and will share it with 3 friends
// on the following day. Assuming nobody receives the advertisement twice, determine how many people
// have liked the ad by the end of a given day, beginning with launch day as day 1.

// For example, assume you want to know how many have liked the ad by the end of the 5th day.

/*
 * Complete the 'viralAdvertising' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

//
//Example
//  n = 5
//  The cumulative number of likes is 24.

//  Day Shared Liked Cumulative
//  1      5     2       2
//  2      6     3       5
//  3      9     4       9
//  4     12     6      15
//  5     18     9      24

func viralAdvertising(n int32) int32 {
	// Write your code here
	var shared, liked, cumulative int32 = 5, 0, 0
	for i := int32(1); i <= n; i++ {
		liked = shared / 2
		cumulative += liked
		shared = liked * 3
	}
	return cumulative
}

func viralAdvertising1(n int32) int32 {
	// Write your code here
	var shared, liked, cumulative int32 = 5, 0, 0
	for i := int32(1); i <= n; i++ {
		liked = shared / 2
		cumulative += liked
		shared = liked * 3
	}
	return cumulative
}

func test() {
	fmt.Println(viralAdvertising1(5)) // 24
	fmt.Println(viralAdvertising1(3)) // 9
}

func main() {
	test()
}

// Explanation of viralAdvertising Function
// The viralAdvertising function calculates the cumulative number of likes a marketing campaign
// receives over a given number of days. Each day, the number of people who receive the
// advertisement is halved and rounded down, and each person who receives the advertisement
// likes it and shares it with 3 friends.

// func viralAdvertising(n int32) int32

// n: The number of days the advertisement campaign runs.
// Algorithm
// Initialize variables:
// shared to 5 (the number of people who receive the advertisement on the first day).
// cumulative to 0 (the cumulative number of likes).
// Iterate through each day from 1 to n using a for loop.
// For each day:
// Calculate the number of likes as shared // 2.
// Add the number of likes to cumulative.
// Calculate the number of people who will receive the advertisement the next day as likes * 3.
// After the loop, return cumulative.

// Start
//  |
//  v
// Initialize shared to 5
// Initialize cumulative to 0
//  |
//  v
// For day = 1 to n
//  |  |
//  |  v
//  |  Calculate likes as shared // 2
//  |  |
//  |  v
//  |  Add likes to cumulative
//  |  |
//  |  v
//  |  Calculate shared as likes * 3
//  |  |
//  v  Repeat loop
//  |
//  v
// Return cumulative
//  |
//  v
// End

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n), where n is the number of days.
// This is because the function iterates through each day exactly once to calculate the cumulative number of likes.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the shared, cumulative, and likes variables.

// Real-Time Use Cases
// Marketing Campaigns:

// Estimating the reach and impact of viral marketing campaigns over a period of time.
// Social Media:

// Predicting the spread of content (e.g., posts, videos) on social media platforms.
// Epidemiology:

// Modeling the spread of information or diseases in a population.
// Product Launches:

// Forecasting the adoption and popularity of new products based on initial reception.

// Summary
// Algorithm: Iterate through each day, calculate the number of likes, update the cumulative likes, and determine the number of people who will receive the advertisement the next day.
// Flowchart: Visual representation of the decision-making process for calculating the cumulative number of likes.
// Time Complexity: O(n) because the function iterates through each day exactly once.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Marketing campaigns, social media, epidemiology, product launches.
// This detailed explanation, flowchart, complexity analysis, and use cases show how the viralAdvertising function works step-by-step to calculate the cumulative number of likes a marketing campaign receives over a given number of da
