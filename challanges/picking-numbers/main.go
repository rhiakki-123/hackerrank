package main

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

// Time and Space Complexity
// Time Complexity: // The time complexity of the function is O(n^2) because there are two nested loops that iterate over the input array a.
// Space Complexity: // The space complexity of the function is O(1) because it uses a constant amount of extra space regardless of the input size.

func pickingNumbers(a []int32) int32 {
	var max int32 = 0
	for i := 0; i < len(a); i++ {
		var count int32 = 0
		for j := 0; j < len(a); j++ {
			if a[j] == a[i] || a[j] == a[i]+1 { // if the element at index j is equal to the element at index i or the element at index j is equal to the element at index i + 1, increment count by 1
				count++
			}
		}
		if count > max {
			max = count
		}
	}
	return max
}

func test() {
	println(pickingNumbers([]int32{4, 6, 5, 3, 3, 1})) // 3
	println(pickingNumbers([]int32{1, 2, 2, 3, 1, 2})) // 5
}

func main() {
	test()
}

// Time and Space Complexity
// Time Complexity:
// The time complexity of the function is O(n^2) because there are two nested loops that iterate over the input array a.
// Space Complexity:
// The space complexity of the function is O(1) because it uses a constant amount of extra space regardless of the input size.

// Real-Time Use Cases
// Data Analysis:
// Finding the longest subsequence of similar values in a dataset.
// Gaming:
// Determining the longest sequence of similar scores or levels in a game.
// Inventory Management:
// Identifying the longest sequence of similar items in an inventory list.
