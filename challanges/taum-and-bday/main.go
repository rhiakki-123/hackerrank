package main

import "fmt"

// Taum is planning to celebrate the birthday of his friend, Diksha.
// There are two types of gifts that Diksha wants from Taum:
// one is black and the other is white.
// To make her happy, Taum has to buy b black gifts and w white gifts.

// The cost of each black gift is bc units.
// The cost of each white gift is wc units.
// The cost to convert each black gift into white gift or vice versa is z units.

// Determine the minimum cost of Dikshas gifts.

// For example, Taum wants to buy b = 3 black gifts and w = 5 white gifts.
// The cost of black gifts is bc = 3 units.
// The cost of white gifts is wc = 4 units.
// The cost to convert black gifts into white gifts is z = 1 units.

// He can buy a black gift for 3 and convert it to a white gift for 1,
// making the total cost of each white gift 4.
// That matches the cost of a white gift,
// so he can do that or just buy black gifts and white gifts.
// Either way, the overall cost is  (3*3) + (5*4) = 29.

/*
 * Complete the 'taumBday' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER b
 *  2. INTEGER w
 *  3. INTEGER bc
 *  4. INTEGER wc
 *  5. INTEGER z
 */

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	var cost int64
	if bc > wc+z {
		cost = int64(wc)*int64(b+w) + int64(z)*int64(b)
	} else if wc > bc+z {
		cost = int64(bc)*int64(b+w) + int64(z)*int64(w)
	} else {
		cost = int64(bc)*int64(b) + int64(wc)*int64(w)
	}
	return cost
}

func test() {
	fmt.Println(taumBday(10, 10, 1, 1, 1)) // 20
	fmt.Println(taumBday(5, 9, 2, 3, 4))   // 37
	fmt.Println(taumBday(3, 6, 9, 1, 1))   // 12
	fmt.Println(taumBday(7, 7, 4, 2, 1))   // 35
	fmt.Println(taumBday(3, 5, 3, 4, 1))   // 29
}

func main() {
	test()
}

//explain function func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64
// with its in details execution, algo, flowchart, time
// and space complexity, write its optimized version, unit tests and use cases

// Explanation of taumBday Function
// The taumBday function calculates the minimum cost for Taum to buy b black gifts and w white gifts, given the costs of black and white gifts and the cost to convert one type of gift to another.

// Function Signature
// func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64

// Parameters
// b: The number of black gifts Taum wants to buy.
// w: The number of white gifts Taum wants to buy.
// bc: The cost of each black gift.
// wc: The cost of each white gift.
// z: The cost to convert one type of gift to another.

// Logic
// Calculate the cost of buying all black gifts directly.
// Calculate the cost of buying all white gifts directly.
// Calculate the cost of converting white gifts to black gifts and then buying them.
// Calculate the cost of converting black gifts to white gifts and then buying them.
// The minimum cost is the minimum of the above four costs.

// Algorithm
// Calculate the cost of buying all black gifts directly: costBlack = b * bc.
// Calculate the cost of buying all white gifts directly: costWhite = w * wc.
// Calculate the cost of converting white gifts to black gifts and then buying them: costConvertWhiteToBlack = w * (bc + z).
// Calculate the cost of converting black gifts to white gifts and then buying them: costConvertBlackToWhite = b * (wc + z).
// The total cost is the sum of the minimum of the direct and conversion costs for black and white gifts.
// Return the total cost.

// Flowchart
// Start
//  |
//  v
// Calculate costBlack as b * bc
//  |
//  v
// Calculate costWhite as w * wc
//  |
//  v
// Calculate costConvertWhiteToBlack as w * (bc + z)
//  |
//  v
// Calculate costConvertBlackToWhite as b * (wc + z)
//  |
//  v
// Calculate totalCost as min(costBlack, costConvertWhiteToBlack) + min(costWhite, costConvertBlackToWhite)
//  |
//  v
// Return totalCost
//  |
//  v
// End

// Detailed Execution
// Given the inputs:
// b := int32(10)
// w := int32(10)
// bc := int32(1)
// wc := int32(1)
// z := int32(1)

// Calculate the cost of buying all black gifts directly:
// costBlack = b * bc = 10 * 1 = 10

// Calculate the cost of buying all white gifts directly:
// costWhite = w * wc = 10 * 1 = 10

// Calculate the cost of converting white gifts to black gifts and then buying them:
// costConvertWhiteToBlack = w * (bc + z) = 10 * (1 + 1) = 20

// Calculate the cost of converting black gifts to white gifts and then buying them:
// costConvertBlackToWhite = b * (wc + z) = 10 * (1 + 1) = 20

// The total cost is the sum of the minimum of the direct and conversion costs for black and white gifts:
// totalCost = min(costBlack, costConvertWhiteToBlack) + min(costWhite, costConvertBlackToWhite)
// 		   = min(10, 20) + min(10, 20)
// 		   = 10 + 10
// 		   = 20

// Return the total cost: 20

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(1).
// This is because the function performs a constant number of operations regardless of the input size.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the variables representing the costs.

// Real-Time Use Cases
// Cost Optimization:

// Calculating the minimum cost of purchasing items with the option to convert one type of item to another.
// Budget Planning:

// Determining the most cost-effective way to allocate a budget for purchasing different types of items.
// Resource Allocation:

// Optimizing the allocation of resources when there are conversion costs between different types of resources.

// Optimized Version of taumBday Function
// package main

// import "fmt"

// func min(a, b int64) int64 {
//     if a < b {
//         return a
//     }
//     return b
// }

// func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
//     b64 := int64(b)
//     w64 := int64(w)
//     bc64 := int64(bc)
//     wc64 := int64(wc)
//     z64 := int64(z)

//     costBlack := b64 * bc64
//     costWhite := w64 * wc64
//     costConvertWhiteToBlack := w64 * (bc64 + z64)
//     costConvertBlackToWhite := b64 * (wc64 + z64)

//     totalCost := min(costBlack, costConvertWhiteToBlack) + min(costWhite, costConvertBlackToWhite)
//     return totalCost
// }

// func testTaumBday() {
//     b := int32(10)
//     w := int32(10)
//     bc := int32(1)
//     wc := int32(1)
//     z := int32(1)
//     fmt.Println(taumBday(b, w, bc, wc, z)) // Output: 20
// }

// func main() {
//     testTaumBday()
// }

// Unit Tests
// package main

// import "testing"

// func TestTaumBday(t *testing.T) {
//     tests := []struct {
//         b        int32
//         w        int32
//         bc       int32
//         wc       int32
//         z        int32
//         expected int64
//     }{
//         {10, 10, 1, 1, 1, 20},
//         {5, 9, 2, 3, 4, 37},
//         {3, 6, 9, 1, 1, 12},
//         {7, 7, 4, 2, 1, 35},
//         {3, 3, 1, 9, 2, 12},
//     }

//     for _, test := range tests {
//         result := taumBday(test.b, test.w, test.bc, test.wc, test.z)
//         if result != test.expected {
//             t.Errorf("For input b=%d, w=%d, bc=%d, wc=%d, z=%d, expected %d but got %d", test.b, test.w, test.bc, test.wc, test.z, test.expected, result)
//         }
//     }
// }

// Summary
// Logic: Calculate the minimum cost for Taum to buy b black gifts and w white gifts by considering the direct and conversion costs for both types of gifts.
// Algorithm: Initialize variables to store the costs, calculate the direct and conversion costs, find the minimum costs, and return the total cost.
// Flowchart: Visual representation of the decision-making process for calculating the minimum cost for Taum to buy b black gifts and w white gifts.
// Time Complexity: O(1) because the function performs a constant number of operations regardless of the input size.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Cost optimization, budget planning, resource allocation.
