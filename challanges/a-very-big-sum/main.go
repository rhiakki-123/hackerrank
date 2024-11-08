package main

import (
	"fmt"
)

/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

func aVeryBigSum(ar []int64) int64 {
	// Write your code here
	var sum int64          // 0
	for _, v := range ar { // 1000000001, 1000000002, 1000000003, 1000000004, 1000000005
		sum += v // 0 + 1000000001 = 1000000001
	}
	return sum // 5000000015
}

func aVeryBigSum1(ar []int64) int64 {
	// Write your code here
	var sum int64                  // 0
	for i := 0; i < len(ar); i++ { // 1000000001, 1000000002, 1000000003, 1000000004, 1000000005
		sum += ar[i] // 0 + 1000000001 = 1000000001
	} // 5000000015
	return sum
}

func test() {
	fmt.Println(aVeryBigSum1([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})) // 5000000015

	fmt.Println(aVeryBigSum1([]int64{1, 2, 3, 4, 5})) // 15
}

func main() {
	test() // run test cases

}
