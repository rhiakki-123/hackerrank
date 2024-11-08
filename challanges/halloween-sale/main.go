package main

import "fmt"

// You wish to buy video games from the famous online video game store Mist.
// Usually, all games are sold at the same price, p dollars. However, they are planning to have the seasonal Halloween Sale next month
// in which you can buy games at a cheaper price. Specifically, the first game will cost p dollars,
// and every subsequent game will cost d dollars less than the previous one.
// This continues until the cost becomes less than or equal to m dollars,
// after which every game will cost m dollars.
// How many games can you buy during the Halloween Sale?

// Example
// p = 20
// d = 3
// m = 6
// s = 70

// The first game will cost 20 dollars. Every subsequent game will cost 3 dollars less, so the second game will cost 17 dollars.
// The third game will cost 14 dollars, and the fourth will cost 11 dollars. Then, the cost will decrease to 6 dollars,
// so the fifth and final game will cost 6 dollars.

// You have 70 dollars in your Mist wallet. You can buy 5 games: 20 + 17 + 14 + 11 + 6 = 68 dollars.

/*
 * Complete the 'howManyGames' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER p
 *  2. INTEGER d
 *  3. INTEGER m
 *  4. INTEGER s
 */

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	// Write your code here
	var games int32
	for s >= p {
		s -= p
		games++
		if p-d >= m {
			p -= d
		} else {
			p = m
		}
	}
	return games
}

func testHowManyGames() {
	// Test case
	p := int32(20)
	d := int32(3)
	m := int32(6)
	s := int32(80)
	result := howManyGames(p, d, m, s)
	fmt.Println(result) // Expected output: 5
}

func main() {
	testHowManyGames()
}

// Explain in depth func howManyGames(p int32, d int32, m int32, s int32) int32
// Logic
// Flowchart
// loop execution for each element in the array
// Algorithm
// Time Complexity
// Space Complexity
// Optimized version
// unit tests
// use cases

// Function Purpose
// howManyGames calculates the maximum number of games that can be bought with a given
// amount of money s, starting with an initial price p, decreasing by d after each
// purchase until the price reaches a minimum price m.

// Logic
// Start with the initial price p.
// Deduct the price from the total money s and count the game.
// Decrease the price by d for the next game, but ensure it doesnt go below m.
// Repeat until there is not enough money to buy another game.

// Flowchart
// graph TD
//     A[Start] --> B[Input: p, d, m, s]
//     B --> C[Initialize count to 0]
//     C --> D[While s >= p]
//     D --> E[Deduct p from s]
//     E --> F[Increment count]
//     F --> G[Decrease p by d]
//     G --> H[If p < m, set p to m]
//     H --> D
//     D --> I[End loop]
//     I --> J[Return count]
//     J --> K[End]

// Algorithm
// func howManyGames(p int32, d int32, m int32, s int32) int32 {
//     count := int32(0)

//     for s >= p {
//         s -= p
//         count++
//         if p-d > m {
//             p -= d
//         } else {
//             p = m
//         }
//     }

//     return count
// }

// Loop Execution for Each Element
// Initialization:

// count is set to 0.
// Loop Execution:

// While s >= p:
// Deduct p from s.
// Increment count.
// Decrease p by d, but ensure it doesnt go below m.

// p := int32(20)
// d := int32(3)
// m := int32(6)
// s := int32(80)

// Initial state:
// count = 0
// p = 20
// s = 80

// Iteration 1:
// s = 80 - 20 = 60
// count = 1
// p = 20 - 3 = 17

// Iteration 2:
// s = 60 - 17 = 43
// count = 2
// p = 17 - 3 = 14

// Iteration 3:
// s = 43 - 14 = 29
// count = 3
// p = 14 - 3 = 11

// Iteration 4:
// s = 29 - 11 = 18
// count = 4
// p = 11 - 3 = 8

// Iteration 5:
// s = 18 - 8 = 10
// count = 5
// p = 8 - 3 = 5

// teration 6:
// s = 10 - 5 = 5
// count = 6
// p = 5 (remains 5 as it is the minimum price)

// Iteration 7:
// s = 5 - 5 = 0
// count = 7
// p = 5

// End of loop:
// Return count = 7

// Time Complexity
// The time complexity is O(n), where n is the number of games that can be bought. This is because we are iterating until the money runs out.
// Space Complexity
// The space complexity is O(1) as we are using a constant amount of extra space.

// Optimized Version
// The given implementation is already optimized with a time complexity of O(n) and space complexity of O(1).

// Unit Tests
// func TestHowManyGames(t *testing.T) {
//     tests := []struct {
//         p, d, m, s int32
//         want       int32
//     }{
//         {20, 3, 6, 80, 7},
//         {20, 3, 6, 85, 8},
//         {20, 3, 6, 100, 9},
//         {20, 3, 6, 10, 0},
//         {20, 3, 6, 20, 1},
//     }

//     for _, tt := range tests {
//         got := howManyGames(tt.p, tt.d, tt.m, tt.s)
//         if got != tt.want {
//             t.Errorf("howManyGames(%v, %v, %v, %v) = %v; want %v",
//                      tt.p, tt.d, tt.m, tt.s, got, tt.want)
//         }
//     }
// }

// Use Cases
// Budget planning for purchasing items with decreasing prices.
// Simulating purchasing strategies in games or applications.
// Financial analysis for bulk purchasing with discounts.
// Inventory management where prices decrease over time.
