package main

import (
	"sort"
)

// David has several containers, each with a number of balls in it. He has just enough containers
// to sort each type of ball he has into its own container.
// David wants to sort the balls using his sort method.
// David wants to perform some number of swap operations such that:
// Each container contains only balls of the same type.
// No two balls of the same type are located in different containers.
// You must perform q queries where each query is in the form of a matrix, M, with dimensions n x n.
// For each query, print Possible on a new line if David can satisfy the conditions above for the given matrix; otherwise, print Impossible instead.

/*
 * Complete the 'organizingContainers' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts 2D_INTEGER_ARRAY container as parameter.
 */

//optimized version

func organizingContainers(container [][]int32) string {
	n := len(container)
	containerCapacities := make([]int64, n)
	typeQuantities := make([]int64, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			value := int64(container[i][j])
			containerCapacities[i] += value // Total balls in container i
			typeQuantities[j] += value      // Total balls of type j
		}
	}

	// Sort the totals to compare them
	sort.Slice(containerCapacities, func(i, j int) bool { return containerCapacities[i] < containerCapacities[j] })
	sort.Slice(typeQuantities, func(i, j int) bool { return typeQuantities[i] < typeQuantities[j] })

	// Compare the sorted totals
	for i := 0; i < n; i++ {
		if containerCapacities[i] != typeQuantities[i] {
			return "Impossible"
		}
	}

	return "Possible"
}

func organizingContainers1(container [][]int32) string {
	var n = len(container)
	var rowSum = make([]int32, n)
	var colSum = make([]int32, n)
	for i := 0; i < n; i++ { // calculate rowSum[i] and colSum[j]
		for j := 0; j < n; j++ { // for each container i and ball j in it add the number of balls to rowSum[i] and colSum[j] respectively
			rowSum[i] += container[i][j] // sum of balls in each container i
			colSum[j] += container[i][j] // sum of balls of type j in all containers
		}
	}
	for i := 0; i < n; i++ { // check if rowSum[i] == colSum[j] for some j
		var j int
		for j = i; j < n; j++ { // for each container i check if there is a container j with the same number of balls	 of type i in it	 and swap them if found
			if rowSum[i] == colSum[j] {
				colSum[j] = colSum[i]
				break
			}
		}
		if j == n {
			return "Impossible"
		}
	}
	return "Possible"
}

func test() {
	var container = [][]int32{
		{1, 1},
		{1, 1},
	}
	println(organizingContainers(container)) // Possible

	container = [][]int32{
		{0, 2},
		{1, 1},
	}
	println(organizingContainers(container)) // Impossible

	container = [][]int32{
		{1, 3, 1},
		{2, 1, 2},
		{3, 3, 3},
	}
	println(organizingContainers1(container)) // Impossible

	container = [][]int32{
		{0, 2, 1},
		{1, 1, 1},
		{2, 0, 0},
	}
	println(organizingContainers(container)) // Possible

}

func main() {
	test()
}

// Explain function func organizingContainers(container [][]int32) string
// with its in details execution, algo, flowchart, time
// and space complexity, write its optimized version, unit tests and use cases

// Explanation of organizingContainers Function
// The organizingContainers function determines whether it is possible to rearrange the balls
// in containers such that each container contains only balls of the same type.
// The initial arrangement of balls is given in a two-dimensional array,
// where each row represents a container, and each column represents a type of ball.

// Function Signature
// func organizingContainers(container [][]int32) string

// Parameters:
// container: A two-dimensional slice of integers where container[i][j] represents the number of balls of type j in container i.
// Returns:
// A string "Possible" if the containers can be reorganized so that each contains only one type of ball.
// A string "Impossible" otherwise.

// Logic
// Calculate Container Capacities:

// For each container (row), calculate the total number of balls it contains.
// Store these totals in a slice called containerCapacities.
// Calculate Ball Type Totals:

// For each ball type (column), calculate the total number of balls of that type across all containers.
// Store these totals in a slice called typeQuantities.
// Compare Capacities and Quantities:

// Sort both containerCapacities and typeQuantities.
// If the sorted lists are equal, it means that its possible to rearrange the balls so that each container contains only one type of ball.
// If not, its impossible.

// Algorithm
// Initialization:
// Determine the number of containers (and types) n.
// Initialize containerCapacities and typeQuantities slices of length n.

// Compute Totals:
// Loop over the containers and ball types to sum up capacities and quantities
// for i := 0; i < n; i++ {
//     for j := 0; j < n; j++ {
//         containerCapacities[i] += container[i][j]   // Sum of balls in container i
//         typeQuantities[j] += container[i][j]        // Total balls of type j
//     }
// }

// Handle Large Numbers:
// Convert sums to int64 to prevent integer overflow for large input values.

// Sort the Totals:
// Sort containerCapacities and typeQuantities in ascending order.

// Compare the Totals:
// Check if each corresponding element in the sorted lists is equal.
// If all elements match, return "Possible".
// Otherwise, return "Impossible".

// Flowchart
// Start
//  |
//  v
// Initialize containerCapacities and typeQuantities slices
//  |
//  v
// For i from 0 to n-1
//  |  For j from 0 to n-1
//  |    containerCapacities[i] += container[i][j]
//  |    typeQuantities[j] += container[i][j]
//  |
//  v
// Convert totals to int64 to prevent overflow
//  |
//  v
// Sort containerCapacities and typeQuantities
//  |
//  v
// For i from 0 to n-1
//  |  If containerCapacities[i] != typeQuantities[i]
//  |    Return "Impossible"
//  |
//  v
// Return "Possible"

// Detailed Execution
// Example Input:
// container := [][]int32{
//     {1, 4},
//     {2, 3},
// }

// Step-by-Step Execution:

// Initialization:
// n := 2
// containerCapacities := make([]int64, n) // [0, 0]
// typeQuantities := make([]int64, n)      // [0, 0]

// Compute Totals:
// Container 0:
// containerCapacities[0] = 1 + 4 = 5
// typeQuantities[0] += 1 (from container[0][0]) => typeQuantities[0] = 1
// typeQuantities[1] += 4 (from container[0][1]) => typeQuantities[1] = 4

// Container 1:
// containerCapacities[1] = 2 + 3 = 5
// typeQuantities[0] += 2 (from container[1][0]) => typeQuantities[0] = 3
// typeQuantities[1] += 3 (from container[1][1]) => typeQuantities[1] = 7

// Convert to int64:
// Ensure all calculations are done using int64 to handle large sums.

// Sort the Totals:
// containerCapacities after sorting: [5, 5]
// typeQuantities after sorting: [3, 7]

// Compare the Totals:
// Compare containerCapacities[0] and typeQuantities[0]: 5 vs 3 (not equal)
// Since they are not equal, return "Impossible".

// Time and Space Complexity
// Time Complexity
// Calculating Totals: O(n²)
// Nested loops over n containers and n types.
// Sorting Totals: O(n log n)
// Sorting two slices of length n.
// Comparing Totals: O(n)
// Comparing each element in the sorted lists.
// Overall Time Complexity: O(n²)
// The nested loops dominate the time complexity.

// Space Complexity
// Additional Space: O(n)
// Two slices of length n for capacities and quantities.
// Overall Space Complexity: O(n)

// Optimized Version
// The function can be optimized by using a single loop to calculate the totals and comparing them in the same loop.
// This approach reduces the time complexity to O(n) and simplifies the logic.

// package main

// import (
//     "fmt"
//     "sort"
// )

// func organizingContainers(container [][]int32) string {
//     n := len(container)
//     containerCapacities := make([]int64, n)
//     typeQuantities := make([]int64, n)

//     for i := 0; i < n; i++ {
//         for j := 0; j < n; j++ {
//             value := int64(container[i][j])
//             containerCapacities[i] += value   // Total balls in container i
//             typeQuantities[j] += value        // Total balls of type j
//         }
//     }

//     // Sort the totals to compare them
//     sort.Slice(containerCapacities, func(i, j int) bool { return containerCapacities[i] < containerCapacities[j] })
//     sort.Slice(typeQuantities, func(i, j int) bool { return typeQuantities[i] < typeQuantities[j] })

//     // Compare the sorted totals
//     for i := 0; i < n; i++ {
//         if containerCapacities[i] != typeQuantities[i] {
//             return "Impossible"
//         }
//     }

//     return "Possible"
// }

// func main() {
//     container1 := [][]int32{
//         {1, 1},
//         {1, 1},
//     }
//     fmt.Println(organizingContainers(container1)) // Output: Possible

//     container2 := [][]int32{
//         {0, 2, 1},
//         {1, 1, 1},
//         {2, 0, 0},
//     }
//     fmt.Println(organizingContainers(container2)) // Output: Impossible
// }

// Unit Tests
// The organizingContainers function can be tested using unit tests to verify its correctness.

// package main

// import (
//     "testing"
// )

// func TestOrganizingContainers(t *testing.T) {
//     tests := []struct {
//         container [][]int32
//         expected  string
//     }{
//         {
//             container: [][]int32{
//                 {1, 1},
//                 {1, 1},
//             },
//             expected: "Possible",
//         },
//         {
//             container: [][]int32{
//                 {0, 2, 1},
//                 {1, 1, 1},
//                 {2, 0, 0},
//             },
//             expected: "Impossible",
//         },
//         {
//             container: [][]int32{
//                 {2, 0, 0},
//                 {0, 2, 0},
//                 {0, 0, 2},
//             },
//             expected: "Possible",
//         },
//         {
//             container: [][]int32{
//                 {999336263, 998799923},
//                 {998799923, 999763019},
//             },
//             expected: "Possible",
//         },
//         {
//             container: [][]int32{
//                 {997612619, 934920795, 998879231, 999926463},
//                 {960369681, 997828120, 999792735, 979622676},
//                 {999013654, 998634077, 997988323, 958769423},
//                 {997409523, 999301350, 940952923, 993020546},
//             },
//             expected: "Possible",
//         },
//     }

//     for _, test := range tests {
//         result := organizingContainers(test.container)
//         if result != test.expected {
//             t.Errorf("For input %v, expected %s but got %s", test.container, test.expected, result)
//         }
//     }
// }

// Real-Time Use Cases
// Warehouse Reorganization:

// Determining if a set of products can be reorganized so that each warehouse contains only one type of product, based on existing stock levels.
// Resource Allocation in Cloud Computing:

// Allocating virtual machines to servers such that each server handles only one type of application, considering the capacities and resource demands.
// Chemical Storage Management:

// Organizing chemicals in storage containers so that each container holds only one type of chemical, crucial for safety regulations.
// Logistics and Shipping:

// Arranging cargo in shipping containers where each container contains only one type of good, optimizing for customs regulations or handling requirements.

// Summary
// Logic: The problem reduces to checking if the multiset of container capacities equals the multiset of total quantities for each ball type.
// Algorithm:
// Calculate the total balls in each container.
// Calculate the total balls of each type.
// Sort and compare these totals.
// Flowchart: Outlines the step-by-step process from initialization to comparison.
// Time Complexity: O(n²)
// Due to iterating over an n x n matrix to calculate totals.
// Space Complexity: O(n)
// Storage for capacities and quantities.
// Use Cases: Applicable in various fields where reorganization is needed under specific constraints.
