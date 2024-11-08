package main

import "fmt"

// This is a staircase of size : n= 4
//    #
//   ##
//  ###
// ####

// Its base and height are both equal to . It is drawn using # symbols and spaces. The last line is not preceded by any spaces.

// Write a program that prints a staircase of size .

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here
	for i := 1; i <= int(n); i++ {
		for j := 1; j <= int(n); j++ {
			if j <= int(n)-i {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func test() {
	staircase(4)
}

func main() {
	test()
}

// flowchart

// Start
//  |
//  v
// Define function staircase(n)
//  |
//  v
// For i = 1 to 4
//  |  |
//  |  v
//  |  For j = 1 to 4
//  |  |  |
//  |  |  v
//  |  |  Is j <= 4 - i?
//  |  |  / \
//  |  | Yes  No
//  |  |  |    |
//  |  |  v    v
//  |  | Print " "
//  |  |  |    |
//  |  |  v    Print "#"
//  |  |  |
//  |  v  Repeat inner loop
//  |
//  v
// Print newline
//  |
//  v
// Repeat outer loop
//  |
//  v
// End

// Detailed Execution with Values
// First Iteration (i = 1):

// j = 1: j <= 4 - 1 → 1 <= 3 → Yes → Print " "
// j = 2: j <= 4 - 1 → 2 <= 3 → Yes → Print " "
// j = 3: j <= 4 - 1 → 3 <= 3 → Yes → Print " "
// j = 4: j <= 4 - 1 → 4 <= 3 → No → Print "#"
// Print newline
// Second Iteration (i = 2):

// j = 1: j <= 4 - 2 → 1 <= 2 → Yes → Print " "
// j = 2: j <= 4 - 2 → 2 <= 2 → Yes → Print " "
// j = 3: j <= 4 - 2 → 3 <= 2 → No → Print "#"
// j = 4: j <= 4 - 2 → 4 <= 2 → No → Print "#"
// Print newline
// Third Iteration (i = 3):

// j = 1: j <= 4 - 3 → 1 <= 1 → Yes → Print " "
// j = 2: j <= 4 - 3 → 2 <= 1 → No → Print "#"
// j = 3: j <= 4 - 3 → 3 <= 1 → No → Print "#"
// j = 4: j <= 4 - 3 → 4 <= 1 → No → Print "#"
// Print newline
// Fourth Iteration (i = 4):

// j = 1: j <= 4 - 4 → 1 <= 0 → No → Print "#"
// j = 2: j <= 4 - 4 → 2 <= 0 → No → Print "#"
// j = 3: j <= 4 - 4 → 3 <= 0 → No → Print "#"
// j = 4: j <= 4 - 4 → 4 <= 0 → No → Print "#"
// Print newline
// End
