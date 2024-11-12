package main

import "fmt"

// Little Bobby loves chocolate. He frequently goes to his favorite 5 & 10 store, Penny Auntie, to buy them.
// They are having a promotion at Penny Auntie. If Bobby saves enough wrappers, he can turn them in
// for a free chocolate.

// For example, Bobby has n=15 to spend on bars of chocolate that cost c=3 each. He can turn in m=2 wrappers to receive another bar.
// Initially, he buys 5 bars and has 5 wrappers after eating them. He turns in 4 of them, leaving him with 1, for 2 more bars.
// After eating those two, he has 3 wrappers, turns in 2 leaving him with 1 wrapper and his new bar.
// Once he eats that one, he has 2 wrappers and turns them in for another bar.
// After eating that one, he only has 1 wrapper, and his feast ends.
// Overall, he has eaten 5 + 2 + 1 + 1 = 9 bars.

/*
 * Complete the 'chocolateFeast' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER c
 *  3. INTEGER m
 */

func chocolateFeast(n int32, c int32, m int32) int32 {
	// Write your code here
	bars := n / c       // initial bars
	wrappers := bars    // initial wrappers
	for wrappers >= m { // exchange wrappers for bars
		bars = bars + wrappers/m // bars from wrappers  exchange
		wrappers = wrappers/m + wrappers%m
	}
	return bars
}

func test() {
	n := int32(15)
	c := int32(3)
	m := int32(2)

	result := chocolateFeast(n, c, m)
	fmt.Println(result)

}

func main() {
	test()
}

// Example Execution
// Lets walk through an example:

// Suppose n = 15, c = 3, and m = 2.
// Initial Purchase:

// bars = 15 / 3 = 5: You can buy 5 bars with 15 units of money.
// wrappers = 5: You have 5 wrappers from the 5 bars.
// Wrapper Exchange Loop:

// First iteration:
// bars += 5 / 2 = 2: You can exchange 5 wrappers for 2 more bars.
// wrappers = 5 / 2 + 5 % 2 = 2 + 1 = 3: You now have 3 wrappers (2 from the exchange and 1 leftover).
// Second iteration:
// bars += 3 / 2 = 1: You can exchange 3 wrappers for 1 more bar.
// wrappers = 3 / 2 + 3 % 2 = 1 + 1 = 2: You now have 2 wrappers (1 from the exchange and 1 leftover).
// Third iteration:
// bars += 2 / 2 = 1: You can exchange 2 wrappers for 1 more bar.
// wrappers = 2 / 2 + 2 % 2 = 1 + 0 = 1: You now have 1 wrapper.
// The loop ends because wrappers < m.
// Return Total Bars:

// return bars: The total number of bars consumed is 5 (initial) + 2 + 1 + 1 = 9.
