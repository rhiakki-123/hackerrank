package main

import "fmt"

// Lexicographical order is often known as alphabetical order when dealing with strings.
// A string is greater than another string
// if it comes later in a lexicographically sorted list.

// Given a word, create a new word by swapping some or all of its characters. This new word must meet two criteria:
// It must be greater than the original word
// It must be the smallest word that meets the first condition

// For example, given the word abcd, the next largest word is abdc.

/*
 * Complete the 'biggerIsGreater' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING w as parameter.
 */

func biggerIsGreater(w string) string {
	// Convert string to a slice of runes for easier manipulation
	runes := []rune(w)
	n := len(runes)

	// Step 1: Find the pivot
	i := n - 2
	for i >= 0 && runes[i] >= runes[i+1] {
		i--
	}
	if i == -1 {
		return "no answer"
	}

	// Step 2: Find the successor
	j := n - 1
	for runes[j] <= runes[i] {
		j--
	}

	// Step 3: Swap the pivot and successor
	runes[i], runes[j] = runes[j], runes[i]

	// Step 4: Reverse the suffix
	left, right := i+1, n-1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

// func testBiggerIsGreater() {
// 	fmt.Println(biggerIsGreater("abdc")) // Output: "acbd"
// 	fmt.Println(biggerIsGreater("abcd")) // Output: "abdc"
// 	fmt.Println(biggerIsGreater("dcba")) // Output: "no answer"
// 	fmt.Println(biggerIsGreater("ab"))   // Output: "ba"
// }

func biggerIsGreater1(w string) string {
	// Write your code here
	runes := []rune(w)
	n := len(runes)

	// Find non-increasing suffix
	i := n - 1
	for i > 0 && runes[i-1] >= runes[i] {
		i--
	}

	// Now i is the head index of the suffix
	if i <= 0 {
		return "no answer"
	}

	// Let's find the rightmost element that is greater than the head
	j := n - 1
	for runes[j] <= runes[i-1] {
		j--
	}

	// Now the value at j is greater than the value at i-1
	runes[i-1], runes[j] = runes[j], runes[i-1]
	j = n - 1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

func test() {
	// Test case
	w := "ab"
	result := biggerIsGreater(w)
	fmt.Println(result)                   // Expected output: "ba"
	fmt.Println(biggerIsGreater1("hefg")) // Expected output: "hegf"
	fmt.Println(biggerIsGreater("dhck"))  // Expected output: "dhkc"
}

func main() {
	test()

}

// Explanation of biggerIsGreater Function
// The biggerIsGreater function finds the next lexicographical permutation of a given string w. If no such permutation exists (i.e., the string is the highest possible permutation), it returns "no answer".

// Function Signature
// func biggerIsGreater(w string) string

// Parameters:
// w: The input string for which the next lexicographical permutation is to be found.
// Returns:
// A string representing the next lexicographical permutation of w, or "no answer" if no such permutation exists.

// Find the Rightmost Character that is Smaller than its Next Character:

// Traverse the string from right to left to find the first character that is smaller than its next character.
// Lets call this character pivot.
// Find the Smallest Character on the Right of the Pivot that is Larger than the Pivot:

// Traverse the string from right to left to find the smallest character that is larger than the pivot. Lets call this character successor.
// Swap the Pivot and Successor:

// Swap the pivot and successor characters.
// Reverse the Suffix:
// Reverse the substring to the right of the pivot to get the next lexicographical permutation.

// Detailed Explanation:
// example: "abdc"

// Step 1: Find the pivot
// Traverse the string from right to left to find the first character that is smaller than its next character. In this case, the pivot is 'b' at index 1.
// w[2] = 'd' and w[3] = 'c' (not smaller)
// w[1] = 'b' and w[2] = 'd' (smaller)
// Pivot found at index 1 (pivot = 'b').

// Step 2: Find the successor
// Traverse from right to left to find the smallest character that is larger than the pivot.
// w[3] = 'c' (larger than pivot)
// Successor found at index 3 (successor = 'c').

// Step 3: Swap the pivot and successor
// Swap the pivot and successor characters.
// Swap w[1] and w[3].
// w = "adbc"

// Step 4: Reverse the suffix
// Reverse the substring to the right of the pivot.
// Reverse the substring to the right of the pivot (from index 2 to the end).
// w = "acbd"

// The next lexicographical permutation of "abdc" is "acbd".

// Time and Space Complexity
// Time Complexity
// Finding the Pivot: O(n)
// Traverse the string from right to left.
// Finding the Successor: O(n)
// Traverse the string from right to left.
// Swapping and Reversing: O(n)
// Swap the pivot and successor, and reverse the suffix.
// Overall Time Complexity: O(n)
// Space Complexity
// Additional Space: O(n)
// Convert the string to a slice of runes for easier manipulation.
// Overall Space Complexity: O(n)

// Summary
// Logic: Find the next lexicographical permutation of a given string by identifying the pivot and successor, swapping them, and reversing the suffix.
// Algorithm: Traverse the string to find the pivot and successor, swap them, and reverse the suffix.
// Time Complexity: O(n) because the function traverses the string multiple times.
// Space Complexity: O(n) because the function uses additional space to store the string as a slice of runes.
// Use Cases: Generating the next permutation in lexicographical order, solving permutation-related problems.

// Start
//  |
//  v
// Convert string w to a slice of runes
//  |
//  v
// Find the Pivot:
// Initialize i = len(w) - 2
// While i >= 0 and w[i] >= w[i + 1]
//     Decrement i
//  |
//  v
// Is i == -1?
//  / \
// Yes  No
//  |    |
//  v    v
// Return "no answer"  Find the Successor:
//                     Initialize j = len(w) - 1
//                     While w[j] <= w[i]
//                         Decrement j
//  |
//  v
// Swap the Pivot and Successor:
// Swap w[i] and w[j]
//  |
//  v
// Reverse the Suffix:
// Initialize left = i + 1
// Initialize right = len(w) - 1
// While left < right
//     Swap w[left] and w[right]
//     Increment left
//     Decrement right
//  |
//  v
// Convert the slice of runes back to a string
//  |
//  v
// Return the resulting string
//  |
//  v
// End

// example Execution
// w = "abdc"
// 1. convert string w to a slice of runes
// runes = []rune("abdc") = ['a', 'b', 'd', 'c']

// 2. Find the Pivot:
// i := 2 // runes[2] = 'd', runes[3] = 'c' (not smaller)
// i := 1 // runes[1] = 'b', runes[2] = 'd' (smaller)
// // Pivot found at index 1

// 3. Find the Successor:
// j := 3 // runes[3] = 'c' (larger than pivot 'b')
// // Successor found at index 3

// 4. Swap the Pivot and Successor:
// runes[1], runes[3] = runes[3], runes[1] // runes = ['a', 'c', 'd', 'b']

// 5. Reverse the Suffix:
// left := 2, right := 3
// runes[2], runes[3] = runes[3], runes[2] // runes = ['a', 'c', 'b', 'd']

// 6. Convert the slice of runes back to a string
// result := string(runes) // result = "acbd"

// 7. Return the resulting string
// Return "acbd"

// End
