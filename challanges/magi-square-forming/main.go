package main

// We define a magic square to be an  n * n matrix of distinct positive integers from 1 to n^2 where the sum of any row, column,
// or diagonal of length  is always equal to the same number: the magic constant.

// You will be given a 3 * 3 matrix s of integers in the inclusive range [1, 9]. We can convert any digit a to any other digit b in
// the range [1, 9] at cost of |a - b|. Given s, convert it into a magic square at minimal cost. Print this cost on a new line.

// Note: The resulting magic square must contain distinct integers in the inclusive range [1, 9].

// Example
// s = [[5, 3, 4], [1, 5, 8], [6, 4, 2]]
// Here is the initial matrix:
// 5 3 4
// 1 5 8
// 6 4 2
// We can convert it to the following magic square:
// 8 3 4
// 1 5 9
// 6 7 2
// This took three replacements at a cost of |5-8| + |8-9| + |4-7| = 7.

// Function Description
// Complete the formingMagicSquare function in the editor below.
// formingMagicSquare has the following parameter(s):

// int s[3][3]: a 3 x 3 array of integers

// Returns
// int: the minimal total cost of converting the input square to a magic square

// Input Format
// Each of the lines contains three space-separated integers of row s[i].

// Constraints
// s[i][j] ∈ [1, 9]

// Sample Input 0
// 4 9 2
// 3 5 7
// 8 1 5

// Sample Output 0
// 1

/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

//  Time and Space Complexity
// Time Complexity:

// The time complexity of the function is O(1) because there are a constant number of possible 3x3 magic squares (8), and the function performs a constant number of operations for each magic square.
// Space Complexity:

// The space complexity of the function is O(1) because it uses a constant amount of extra space regardless of the input size.

func formingMagicSquare(s [][]int32) int32 {
	// Write your code here
	var magicSquares = [][][]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}
	var minCost int32 = 81
	for _, magicSquare := range magicSquares {
		var cost int32 = 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				cost += abs(s[i][j] - magicSquare[i][j])
			}
		}
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func formingMagicSquare1(s [][]int32) int32 {
	// Write your code here
	var magicSquares = [][][]int32{
		{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
		{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
		{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
		{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
		{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
		{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
		{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
		{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
	}
	var minCost int32 = 81
	for _, magicSquare := range magicSquares {
		var cost int32 = 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				cost += abs(s[i][j] - magicSquare[i][j])
			}
		}
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func test() {
	var s = [][]int32{{4, 9, 2}, {3, 5, 7}, {8, 1, 5}}
	println(formingMagicSquare(s)) // 1
}

func main() {
	test()
}

// Start
//  |
//  v
// Define all possible 3x3 magic squares
//  |
//  v
// Initialize minCost to a large value
//  |
//  v
// For each magic square
//  |  |
//  |  v
//  |  Calculate the cost to convert s to the current magic square
//  |  |
//  |  v
//  |  Is the calculated cost < minCost?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | Update minCost
//  |  |
//  v  Repeat loop
//  |
//  v
// Return minCost
//  |
//  v
// End
