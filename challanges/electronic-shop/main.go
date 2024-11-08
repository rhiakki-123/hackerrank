package main

import "fmt"

// A person wants to determine the most expensive computer keyboard and USB drive that can be purchased with
// a give budget. Given price lists for keyboards and USB drives and a budget, find the cost to buy them.
// If it is not possible to buy both items, return .

// Example
// b = 60
// keyboards = [40, 50, 60]
// drives = [5, 8, 12]

// The person can buy 50 keyboard + 8 drives, . Choose the latter as the more expensive option and return 58.

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	// Write your code here

	var max int32 = -1
	for _, k := range keyboards {
		for _, d := range drives {
			if k+d <= b && k+d > max {
				max = k + d
			}
		}
	}
	return max
}

func getMoneySpent1(keyboards []int32, drives []int32, b int32) int32 {
	// Write your code here

	var max int32 = -1
	for i := 0; i < len(keyboards); i++ {
		for j := 0; j < len(drives); j++ {
			if keyboards[i]+drives[j] <= b && keyboards[i]+drives[j] > max {
				max = keyboards[i] + drives[j]
			}
		}
	}
	return max
}

func test() {
	fmt.Println(getMoneySpent1([]int32{40, 50, 60}, []int32{5, 8, 12}, 60)) // 58
}

func main() {
	test()
}
