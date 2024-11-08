package main

import (
	"fmt"
	"strconv"
)

// A modified Kaprekar number is a positive whole number with a special property.
// If you square it, then split the number into two integers and sum those integers,
// you have the same value you started with.

// func kaprekarNumbers(p int32, q int32) []int32 {
// 	// Write your code here
// 	var res []int32
// 	for i := p; i <= q; i++ {
// 		if isKaprekar(i) {
// 			res = append(res, i)
// 		}
// 	}
// 	if len(res) == 0 {
// 		return []int32{-1}
// 	}
// 	return res
// }

// func isKaprekar(n int32) bool {
// 	square := int64(n) * int64(n)
// 	squareStr := strconv.FormatInt(square, 10)
// 	d := len(squareStr) - int(len(strconv.Itoa(int(n))))
// 	left, _ := strconv.Atoi(squareStr[:d])
// 	right, _ := strconv.Atoi(squareStr[d:])
// 	return left+right == int(n)
// }

func kaprekarNumbers(p int32, q int32) {
	for i := p; i <= q; i++ {
		if isKaprekar(i) {
			fmt.Println(i)
		}
	}
}

func isKaprekar(n int32) bool {
	square := int64(n) * int64(n)
	squareStr := strconv.FormatInt(square, 10)
	d := len(squareStr) - int(len(strconv.FormatInt(int64(n), 10)))
	left, _ := strconv.Atoi(squareStr[:d])
	right, _ := strconv.Atoi(squareStr[d:])
	return left+right == int(n)
}

func test() {
	// Test case
	p := int32(400)
	q := int32(700)
	kaprekarNumbers(p, q) // Expected output: 1 9 45 55 99
}

func main() {
	test()
}
