package main

import "fmt"

// You will be given a square chess board with one queen and a number of obstacles placed on it.
// Determine how many squares the queen can attack. A queen is standing on an n x n chessboard.
// The chess boards rows are numbered from 1 to n, going from bottom to top.
// Its columns are numbered from 1 to n, going from left to right.

// Each square is referenced by a tuple, (r, c), describing the row, r, and column, c,
// where the square is located.

// The queen is standing at position (rq, cq). In a single move, she can attack any square in
// any of the eight directions (left, right, up, down, and the four diagonals).
// In the diagram below, the green circles denote all the cells the queen can attack from (4, 4):

// queen-attack.png

// There are obstacles on the chessboard. An obstacle at location (r, c) does not prevent the
// queen from attacking the cell (r, c). The queen cannot move to a square that has an obstacle.

// Given the queens position and the locations of all the obstacles, find and print the number of
// squares the queen can attack from her position at (rq, cq).

// queensAttack has the following parameters:
// - int n: the number of rows and columns in the board
// - nt k: the number of obstacles on the board
// - int r_q: the row number of the queen's position
// - int c_q: the column number of the queen's position
// - int obstacles[k][2]: each element is an array of 2 integers, the row and column of an obstacle

// Returns
// - int: the number of squares the queen can attack

/*
 * Complete the 'queensAttack' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER r_q
 *  4. INTEGER c_q
 *  5. 2D_INTEGER_ARRAY obstacles
 */

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func abs(a int32) int32 {
	if a < 0 {
		return -a
	}
	return a
}

//optimized version
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	up := n - r_q
	down := r_q - 1
	left := c_q - 1
	right := n - c_q
	upLeft := min(n-r_q, c_q-1)
	upRight := min(n-r_q, n-c_q)
	downLeft := min(r_q-1, c_q-1)
	downRight := min(r_q-1, n-c_q)

	for _, obstacle := range obstacles {
		r_o := obstacle[0]
		c_o := obstacle[1]

		if r_o == r_q {
			if c_o < c_q {
				left = min(left, c_q-c_o-1)
			} else {
				right = min(right, c_o-c_q-1)
			}
		} else if c_o == c_q {
			if r_o < r_q {
				down = min(down, r_q-r_o-1)
			} else {
				up = min(up, r_o-r_q-1)
			}
		} else if abs(r_o-r_q) == abs(c_o-c_q) {
			if r_o > r_q && c_o < c_q {
				upLeft = min(upLeft, r_o-r_q-1)
			} else if r_o > r_q && c_o > c_q {
				upRight = min(upRight, r_o-r_q-1)
			} else if r_o < r_q && c_o < c_q {
				downLeft = min(downLeft, r_q-r_o-1)
			} else if r_o < r_q && c_o > c_q {
				downRight = min(downRight, r_q-r_o-1)
			}
		}
	}

	totalMoves := up + down + left + right + upLeft + upRight + downLeft + downRight
	return totalMoves
}

func queensAttack1(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	// Write your code here
	obs := make(map[[2]int32]bool)     // obstacles  map to check if there is an obstacle at a given position
	for _, obsPos := range obstacles { // add obstacles to the map
		obs[[2]int32{obsPos[0], obsPos[1]}] = true // mark the position as an obstacle  in the map
	}

	directions := [8][2]int32{ // directions queen can move
		{1, 0},   // up
		{1, 1},   // up right
		{0, 1},   // right
		{-1, 1},  // down right
		{-1, 0},  // down
		{-1, -1}, // down left
		{0, -1},  // left
		{1, -1},  // up left
	}

	count := int32(0)                // count of squares the queen can attack
	for _, dir := range directions { // loop through all directions
		r := r_q + dir[0]                        // get the next position in the direction
		c := c_q + dir[1]                        // get the next position in the direction
		for r > 0 && r <= n && c > 0 && c <= n { // check if the position is within the board limits and there is no obstacle
			if obs[[2]int32{r, c}] { // if there is an obstacle at the position break
				break // break the loop
			}
			count++     // increment the count of squares the queen can attack from the position (r, c)
			r += dir[0] // move to the next position in the direction
			c += dir[1] // move to the next position in the direction
		}
	}
	return count // return the count of squares the queen can attack
}

func testQueensAttack() {
	n := int32(4)
	k := int32(0)
	r_q := int32(4)
	c_q := int32(4)
	obstacles := [][]int32{}
	result := queensAttack(n, k, r_q, c_q, obstacles)
	fmt.Println(result) // 9

	n = int32(5)
	k = int32(3)
	r_q = int32(4)
	c_q = int32(3)
	obstacles = [][]int32{{5, 5}, {4, 2}, {2, 3}}
	result = queensAttack(n, k, r_q, c_q, obstacles)
	fmt.Println(result) // 10

	n = int32(8)
	k = int32(2)
	r_q = int32(4)
	c_q = int32(4)
	obstacles = [][]int32{{3, 5}, {5, 4}}
	result = queensAttack(n, k, r_q, c_q, obstacles)
	fmt.Println(result) // 23
}

func main() {
	testQueensAttack()
}

//explain function func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32
// with its in details execution, algo, flowchart, time
// and space complexity, write its optimized version, unit tests and use cases

// Explanation of queensAttack Function
// The queensAttack function calculates the number of squares a queen can attack on an n x n chessboard, given the queens position and the positions of obstacles

// Function Signature
// func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32

// n: The number of rows and columns in the chessboard
// k: The number of obstacles on the chessboard
// r_q: The row number of the queens position
// c_q: The column number of the queens position
// obstacles: A 2D array containing the positions of the obstacles

// Logic
// Create a map to store the positions of the obstacles
// Define the directions in which the queen can move (up, down, left, right, and diagonals)
// Iterate through each direction and calculate the number of squares the queen can attack in that direction
// Check for obstacles in each direction and stop if an obstacle is encountered
// Return the total number of squares the queen can attack

// Algorithm
// Initialize variables to store the maximum possible moves in each direction.
// Iterate through the obstacles and update the maximum possible moves in each direction based on the positions of the obstacles.
// Sum the maximum possible moves in all eight directions.
// Return the total number of squares the queen can attack.

// flowchart
// Start
//  |
//  v
// Initialize max moves in each direction
//  |
//  v
// For each obstacle in obstacles
//  |  |
//  |  v
//  |  Update max moves based on obstacle position
//  |  |
//  v  Repeat loop
//  |
//  v
// Sum max moves in all directions
//  |
//  v
// Return total moves
//  |
//  v
// End

// Detailed Execution
// Given the inputs:
// n := int32(8)
// k := int32(1)
// r_q := int32(4)
// c_q := int32(4)
// obstacles := [][]int32{{3, 5}}

// The queen is at position (4, 4) on an 8x8 chessboard with one obstacle at position (3, 5).
// Initialize max moves in each direction:
// up := n - r_q // 8 - 4 = 4
// down := r_q - 1  // 4 - 1 = 3
// left := c_q - 1 // 4 - 1 = 3
// right := n - c_q // 8 - 4 = 4
// upLeft := min(n - r_q, c_q - 1) // min(4, 3) = 3
// upRight := min(n - r_q, n - c_q) // min(4, 4) = 4
// downLeft := min(r_q - 1, c_q - 1) // min(3, 3) = 3
// downRight := min(r_q - 1, n - c_q) // min(3, 4) = 3

// Update max moves based on obstacle position:
// r_o := 3 // Obstacle at (3, 5)
// c_o := 5 //
// Obstacle at (3, 5)
// if r_o == r_q {                                 // if the obstacle is in the same row as the queen
//     if c_o < c_q { 							    // if the obstacle is to the left of the queen
//         left = min(left, c_q - c_o - 1)         // update the maximum moves to the left
//     } else {
//         right = min(right, c_o - c_q - 1)
//     }
// } else if c_o == c_q {
//     if r_o < r_q {
//         down = min(down, r_q - r_o - 1)
//     } else {
//         up = min(up, r_o - r_q - 1)
//     }
// } else if abs(r_o - r_q) == abs(c_o - c_q) {
//     if r_o > r_q && c_o < c_q {
//         upLeft = min(upLeft, r_o - r_q - 1)
//     } else if r_o > r_q && c_o > c_q {
//         upRight = min(upRight, r_o - r_q - 1)
//     } else if r_o < r_q && c_o < c_q {
//         downLeft = min(downLeft, r_q - r_o - 1)
//     } else if r_o < r_q && c_o > c_q {
//         downRight = min(downRight, r_q - r_o - 1)
//     }
// }

// Sum max moves in all directions:
// totalMoves := up + down + left + right + upLeft + upRight + downLeft + downRight // 4 + 3 + 3 + 4 + 3 + 4 + 3 + 3 = 27

// Return total moves: 27

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(k), where k is the number of obstacles.
// This is because the function iterates through the obstacles to update the maximum possible moves in each direction.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the variables representing the maximum possible moves in each direction.

// Real-Time Use Cases
// Game Development:
// Implementing chess game mechanics where the queens possible moves are calculated based on obstacles.
// Pathfinding:
// Calculating the possible moves of an entity in a grid-based environment with obstacles.
// Robotics:
// Programming robots to navigate through environments with obstacles based on their movement capabilities.

// Optimized Version
// package main

// import (
//     "fmt"
//     "math"
// )

// func min(a, b int32) int32 {
//     if a < b {
//         return a
//     }
//     return b
// }

// func abs(a int32) int32 {
//     if a < 0 {
//         return -a
//     }
//     return a
// }

// func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
//     up := n - r_q
//     down := r_q - 1
//     left := c_q - 1
//     right := n - c_q
//     upLeft := min(n - r_q, c_q - 1)
//     upRight := min(n - r_q, n - c_q)
//     downLeft := min(r_q - 1, c_q - 1)
//     downRight := min(r_q - 1, n - c_q)

//     for _, obstacle := range obstacles {
//         r_o := obstacle[0]
//         c_o := obstacle[1]

//         if r_o == r_q {
//             if c_o < c_q {
//                 left = min(left, c_q - c_o - 1)
//             } else {
//                 right = min(right, c_o - c_q - 1)
//             }
//         } else if c_o == c_q {
//             if r_o < r_q {
//                 down = min(down, r_q - r_o - 1)
//             } else {
//                 up = min(up, r_o - r_q - 1)
//             }
//         } else if abs(r_o - r_q) == abs(c_o - c_q) {
//             if r_o > r_q && c_o < c_q {
//                 upLeft = min(upLeft, r_o - r_q - 1)
//             } else if r_o > r_q && c_o > c_q {
//                 upRight = min(upRight, r_o - r_q - 1)
//             } else if r_o < r_q && c_o < c_q {
//                 downLeft = min(downLeft, r_q - r_o - 1)
//             } else if r_o < r_q && c_o > c_q {
//                 downRight = min(downRight, r_q - r_o - 1)
//             }
//         }
//     }

//     totalMoves := up + down + left + right + upLeft + upRight + downLeft + downRight
//     return totalMoves
// }

// func testQueensAttack() {
//     n := int32(8)
//     k := int32(1)
//     r_q := int32(4)
//     c_q := int32(4)
//     obstacles := [][]int32{{3, 5}}
//     fmt.Println(queensAttack(n, k, r_q, c_q, obstacles)) // Output: 24
// }

// func main() {
//     testQueensAttack()
// }

// Unit Tests

// package main

// import "testing"

// func TestQueensAttack(t *testing.T) {
//     tests := []struct {
//         n         int32
//         k         int32
//         r_q       int32
//         c_q       int32
//         obstacles [][]int32
//         expected  int32
//     }{
//         {8, 1, 4, 4, [][]int32{{3, 5}}, 24},
//         {8, 0, 4, 4, [][]int32{}, 27},
//         {8, 2, 4, 4, [][]int32{{3, 5}, {5, 4}}, 23},
//         {5, 3, 4, 3, [][]int32{{5, 5}, {4, 2}, {2, 3}}, 10},
//     }

//     for _, test := range tests {
//         result := queensAttack(test.n, test.k, test.r_q, test.c_q, test.obstacles)
//         if result != test.expected {
//             t.Errorf("For input n=%d, k=%d, r_q=%d, c_q=%d, obstacles=%v, expected %d but got %d", test.n, test.k, test.r_q, test.c_q, test.obstacles, test.expected, result)
//         }
//     }
// }

// Summary
// Logic: Calculate the maximum possible moves in each of the eight directions (up, down, left, right, and the four diagonals) based on the queen's position and the positions of obstacles.
// Algorithm: Initialize variables to store the maximum possible moves in each direction, iterate through the obstacles to update the maximum possible moves, sum the maximum possible moves in all directions, and return the total number of squares the queen can attack.
// Flowchart: Visual representation of the decision-making process for calculating the number of squares a queen can attack on an n x n chessboard.
// Time Complexity: O(k) because the function iterates through the obstacles to update the maximum possible moves in each direction.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Game development, pathfinding, robotics.
