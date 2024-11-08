package main

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

//  You are in charge of the cake for a childs birthday. You have decided the cake will have one candle for each year
//  of their total age. They will only be able to blow out the tallest of the candles. Count how many candles are
//  tallest.

//  Example

//  candles = [4,4,1,3]

//  The maximum height candles are 4 units high. There are 2 of them, so return 2.

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var max, count int32
	for _, v := range candles {
		if v > max {
			max = v
			count = 1
		} else if v == max {
			count++
		}
	}
	return count
}

func birthdayCakeCandles1(candles []int32) int32 {
	// Write your code here
	var max, count int32
	for i := 0; i < len(candles); i++ {
		if candles[i] > max {
			max = candles[i]
			count = 1
		} else if candles[i] == max {
			count++
		}
	}
	return count
}

func test() {
	println(birthdayCakeCandles([]int32{4, 4, 1, 3}))  // 2
	println(birthdayCakeCandles1([]int32{4, 4, 4, 3})) // 3
}

func main() {
	test()
}

// Loop Execution:
// First Iteration (v = 4):

// if v > max → 4 > 0 → Yes
// max = 4
// count = 1
// Second Iteration (v = 4):

// if v > max → 4 > 4 → No
// else if v == max → 4 == 4 → Yes
// count++ → count = 2
// Third Iteration (v = 1):

// if v > max → 1 > 4 → No
// else if v == max → 1 == 4 → No
// Fourth Iteration (v = 3):

// if v > max → 3 > 4 → No
// else if v == max → 3 == 4 → No
// Final State:
// max = 4
// count = 2
// Output:
// 2

// // Flowchart
// Start
//  |
//  v
// Initialize max to 0, count to 0
//  |
//  v
// For each element v in candles
//  |  |
//  |  v
//  |  Is v > max?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | max = v
//  | count = 1
//  |  |
//  |  v
//  |  Is v == max?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | count++
//  |  |
//  |  v
//  | Repeat loop
//  |
//  v
// Return count
//  |
//  v
// End
