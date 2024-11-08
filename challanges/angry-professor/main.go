package main

import "fmt"

// A Discrete Mathematics professor has a class of students. Frustrated with their lack of discipline,
// the professor decides to cancel class if fewer than some number of students are present when
// class starts. Arrival times go from on time (arrivalTime <= 0) to arrived late (arrivalTime > 0).

// Given the arrival time of each student and a threshhold number of attendees, determine if
// the class is canceled

// Example
// n = 5
// k = 3
// a = [-2, -1, 0, 1, 2]

// The first 3 students arrived on. The last 2 were late. The threshold is 3 students,
//so class will go on. Return YES.

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	var count int32 = 0
	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			count++
		}
	}
	if count >= k {
		return "NO"
	}
	return "YES"
}

func test() {
	fmt.Println(angryProfessor(3, []int32{-2, -1, 0, 1, 2})) // NO
	fmt.Println(angryProfessor(2, []int32{0, -1, 2, 1}))     // NO
	fmt.Println(angryProfessor(4, []int32{-1, -3, 4, 2}))    // YES
}

func main() {
	test()
}

// Algorithm
// Initialize a variable count to 0 to keep track of the number of students who arrive on time.
// Iterate through each element in the a array using a for loop.
// For each element, check if the arrival time is less than or equal to 0.
// If the arrival time is less than or equal to 0, increment count.
// After the loop, check if count is greater than or equal to k.
// If count is greater than or equal to k, return "NO" (class is not canceled).
// If count is less than k, return "YES" (class is canceled).

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n), where n is the length of the a array.
// This is because the function iterates through each element in the a array exactly once to count the number of students who arrive on time.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the count variable and the loop variables.

// Start
//  |
//  v
// Initialize count to 0
//  |
//  v
// For i = 0 to len(a)-1
//  |  |
//  |  v
//  |  Is a[i] <= 0?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | Increment count
//  |  |
//  v  Repeat loop
//  |
//  v
// Is count >= k?
//  / \
// Yes  No
//  |    |
//  v    v
// Return "NO"  Return "YES"
//  |
//  v
// End

// Use Cases
// Classroom Management:

// Determining whether a class should be canceled based on student attendance.
// Event Planning:

// Deciding whether to proceed with an event based on the number of attendees who arrive on time.
// Workplace Meetings:

// Determining whether a meeting should be canceled based on the number of employees who arrive on time.
// Transportation Services:

// Deciding whether to cancel a bus or train service based on the number of passengers who arrive on time.
