package main

import "fmt"

// There are a number of people who will be attending ACM-ICPC World Finals. Each of them may be well versed in a number of topics.
// Given a list of topics known by each attendee,
// presented as binary strings,
// determine the maximum number of topics a 2-person team can know.
// Each subject has a column in the binary string,
// and a '1' means the subject is known
// while '0' means it is not.
// Also determine the number of teams that know the maximum number of topics.
// Return an integer array with two elements.
// The first is the maximum number of topics known,
// and the second is the number of teams that know that number of topics.

//  Example
//  topic = ['10101', '11110', '00010']
//  the attendee data aligned for comparison:
//  10101
//  11110
//  00010
//  In this case, the teams would be pairs of attendees.
//  The first pair would know 5 topics.
//  The second pair would know 4 topics.
//  The third pair would know 2 topics.
//  In this example, the maximum number of topics a 2-person team can know is 5,
//  and there are 1 team that knows that many topics.

/*
 * Complete the 'acmTeam' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts STRING_ARRAY topic as parameter.
 */
func acmTeam(topic []string) []int32 {
	// Write your code here
	maxTopics := int32(0)
	maxTeams := int32(0)
	for i := 0; i < len(topic); i++ {
		for j := i + 1; j < len(topic); j++ {
			topics := int32(0)
			for k := 0; k < len(topic[i]); k++ {
				if topic[i][k] == '1' || topic[j][k] == '1' {
					topics++
				}
			}
			if topics > maxTopics {
				maxTopics = topics
				maxTeams = 1
			} else if topics == maxTopics {
				maxTeams++
			}
		}
	}
	return []int32{maxTopics, maxTeams}
}

func testAcmTeam() {
	topic := []string{"10101", "11100", "11010", "00101"}
	result := acmTeam(topic)
	for _, val := range result {
		fmt.Println(val)
	}

	topic1 := []string{"101", "110", "001"}
	result1 := acmTeam(topic1)
	for _, val := range result1 {
		fmt.Println(val)
	}
}

func main() {
	testAcmTeam()
}

//explain function acmTeam(topic []string) []int32
// with its in details execution, algo, flowchart, time
// and space complexity, write its optimized version, unit tests and use cases

// Explanation of acmTeam Function
// The acmTeam function determines the maximum number of topics a team of two people can know and the number of teams that can know that maximum number of topics. Each person is represented by a binary string where each bit indicates whether the person knows a particular topic (1 for yes, 0 for no

// Function Signature
// func acmTeam(topic []string) []int32

// topic: An array of binary strings representing the topics known by each attendee.

// Logic
// Iterate through all pairs of attendees.
// For each pair, calculate the number of topics they collectively know by performing a bitwise OR operation on their topic strings.
// Track the maximum number of topics known by any team and the number of teams that know that maximum number of topics.

// Algorithm
// Initialize maxTopics and maxTeams to 0.
// Iterate through all pairs of attendees using two nested loops.
// For each pair, calculate the number of topics they collectively know by performing a bitwise OR operation on their topic strings.
// If the number of topics is greater than maxTopics, update maxTopics and set maxTeams to 1.
// If the number of topics is equal to maxTopics, increment maxTeams.
// Return an array containing maxTopics and maxTeams.

// flowchart
// Start
//  |
//  v
// Initialize maxTopics to 0
// Initialize numTeams to 0
//  |
//  v
// For each pair of people (i, j)
//  |  |
//  |  v
//  |  Perform bitwise OR on topic[i] and topic[j]
//  |  |
//  |  v
//  |  Count the number of 1s in the result
//  |  |
//  |  v
//  |  Is count > maxTopics?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | Update maxTopics and numTeams  Is count == maxTopics?
//  |  |
//  |  v
//  | Increment numTeams
//  |  |
//  v  Repeat loop
//  |
//  v
// Return [maxTopics, numTeams]
//  |
//  v
// End

// Detailed Execution
// Given the input:
// topic := []string{"10101", "11100", "11010", "00101"}

// Initialize maxTopics and maxTeams to 0.
// maxTopics := 0
// maxTeams := 0

// Iterate through all pairs of attendees:
// Pair 1: (0, 1)
// Perform bitwise OR on topic[0] and topic[1]: "10101" | "11100" = "11101"
// Count the number of 1s in the result: 4
// Is 4 > 0? Yes
// Update maxTopics to 4 and numTeams to 1
// Increment numTeams: 1
// maxTopics = 4 and numTeams = 1

// Pair 2: (0, 2)
// Perform bitwise OR on topic[0] and topic[2]: "10101" | "11010" = "11111"
// Count the number of 1s in the result: 5
// Is 5 > 4? Yes
// Update maxTopics to 5 and numTeams to 1
// Increment numTeams: 1
// maxTopics = 5 and numTeams = 1

// Pair 3: (0, 3)
// Perform bitwise OR on topic[0] and topic[3]: "10101" | "00101" = "10101"
// Count the number of 1s in the result: 3
// Is 3 > 5? No
// Is 3 == 5? No
// maxTopics = 5 and numTeams = 1

// Pair 4: (1, 2)
// Perform bitwise OR on topic[1] and topic[2]: "11100" | "11010" = "11110"
// Count the number of 1s in the result: 4
// Is 4 > 5? No
// Is 4 == 5? No
// maxTopics = 5 and numTeams = 1

// Pair 5: (1, 3)
// Perform bitwise OR on topic[1] and topic[3]: "11100" | "00101" = "11101"
// Count the number of 1s in the result: 4
// Is 4 > 5? No
// Is 4 == 5? No
// maxTopics = 5 and numTeams = 1

// Pair 6: (2, 3)
// Perform bitwise OR on topic[2] and topic[3]: "11010" | "00101" = "11111"
// Count the number of 1s in the result: 5
// Is 5 > 5? No
// Is 5 == 5? Yes
// Increment numTeams: 2
// maxTopics = 5 and numTeams = 2

// Return [5, 2]

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n^2 * m), where n is the number of people and m is the length of the topic strings.
// This is because the function iterates through all possible pairs of people (O(n^2)) and performs a bitwise OR operation and counts the number of 1s in the result (O(m)).
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the variables maxTopics and numTeams.

// Real-Time Use Cases
// Team Formation:
// Forming teams with the maximum combined knowledge in a competitive programming contest.
// Collaboration Analysis:
// Analyzing the potential collaboration between individuals based on their combined knowledge.
// Resource Allocation:
// Allocating resources to teams with the maximum combined skills or knowledge.

// Optimized Version
// package main

// import (
//     "fmt"
//     "strconv"
// )

// func countOnes(s string) int32 {
//     count := int32(0)
//     for _, char := range s {
//         if char == '1' {
//             count++
//         }
//     }
//     return count
// }

// func acmTeam(topic []string) []int32 {
//     maxTopics := int32(0)
//     numTeams := int32(0)

//     for i := 0; i < len(topic); i++ {					   //
//         for j := i + 1; j < len(topic); j++ {
//             topicsKnown := ""
//             for k := 0; k < len(topic[i]); k++ {
//                 if topic[i][k] == '1' || topic[j][k] == '1' {
//                     topicsKnown += "1"
//                 } else {
//                     topicsKnown += "0"
//                 }
//             }
//             count := countOnes(topicsKnown)
//             if count > maxTopics {
//                 maxTopics = count
//                 numTeams = 1
//             } else if count == maxTopics {
//                 numTeams++
//             }
//         }
//     }

//     return []int32{maxTopics, numTeams}
// }

// func testAcmTeam() {
//     topic := []string{"10101", "11100", "11010", "00101"}
//     fmt.Println(acmTeam(topic)) // Output: [5, 2]
// }

// func main() {
//     testAcmTeam()
// }

// Unit Tests
// package main

// import "testing"

// func TestAcmTeam(t *testing.T) {
//     tests := []struct {
//         topic    []string
//         expected []int32
//     }{
//         {[]string{"10101", "11100", "11010", "00101"}, []int32{5, 2}},
//         {[]string{"101", "110", "001"}, []int32{3, 1}},
//         {[]string{"111", "111", "111"}, []int32{3, 3}},
//         {[]string{"000", "000", "000"}, []int32{0, 3}},
//         {[]string{"10101", "11100", "11010"}, []int32{5, 1}},
//     }

//     for _, test := range tests {
//         result := acmTeam(test.topic)
//         if result[0] != test.expected[0] || result[1] != test.expected[1] {
//             t.Errorf("For input %v, expected %v but got %v", test.topic, test.expected, result)
//         }
//     }
// }

// Summary
// Logic: Calculate the maximum number of topics a team of two people can know by performing a bitwise OR operation on their topic strings and count the number of teams that can know that maximum number of topics.
// Algorithm: Initialize variables to store the maximum number of topics and the number of teams, iterate through all possible pairs of people, perform a bitwise OR operation on their topic strings, update the maximum number of topics and the number of teams, and return the result.
// Flowchart: Visual representation of the decision-making process for calculating the maximum number of topics a team of two people can know and the number of teams that can know that maximum number of topics.
// Time Complexity: O(n^2 * m) because the function iterates through all possible pairs of people and performs a bitwise OR operation and counts the number of 1s in the result.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Team formation, collaboration analysis, resource allocation.
