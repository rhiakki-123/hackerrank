package main

import "fmt"

// When a contiguous block of text is selected in a PDF viewer, the selection is highlighted with
// a blue rectangle. In this PDF viewer, each word is highlighted independently. For example:

// There is a list of 26 character heights aligned by index to their letters. For example, 'a' is at
// index  and 'z' is at index . There will also be a string. Using the letter heights given,
// determine the area of the rectangle highlight in mm^2 assuming all letters are 1mm wide.

/*
 * Complete the 'designerPdfViewer' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY h
 *  2. STRING word
 */

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	var max int32 = 0
	for _, c := range word {
		if h[c-'a'] > max {
			max = h[c-'a']
		}
	}
	return max * int32(len(word))
}

func designerPdfViewer1(h []int32, word string) int32 {
	// Write your code here
	var max int32 = 0
	for i := 0; i < len(word); i++ {
		if h[word[i]-'a'] > max {
			max = h[word[i]-'a']
		}
	}
	return max * int32(len(word))
}

func test() {
	fmt.Println(designerPdfViewer1([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 3, 1, 4, 1, 2, 1, 3, 1, 4, 1, 3, 1, 3, 1, 3, 1, 4}, "abc"))  // 9
	fmt.Println(designerPdfViewer1([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 3, 1, 4, 1, 2, 1, 3, 1, 4, 1, 3, 1, 3, 1, 3, 1, 4}, "zaba")) // 16
}

func main() {
	test()
}

// // Algorithm
// Algorithm
// Initialize a variable max to 0 to keep track of the maximum height of the letters in the word.
// Iterate through each character in the word using a for loop.
// For each character, calculate its height using the h array.
// If the height of the current character is greater than max, update max to the height of the current character.
// After the loop, calculate the area of the highlighted rectangle by multiplying max by the length of the word.
// Return the calculated area.

// Start
//  |
//  v
// Initialize max to 0
//  |
//  v
// For i = 0 to len(height)-1
//  |  |
//  |  v
//  |  Is height[i] > max?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  |  Update max to height[i]
//  |  |
//  v  Repeat loop
//  |
//  v
// Is max > k?
//  / \
// Yes  No
//  |    |
//  v    v
// Return max - k  Return 0
//  |
//  v
// End

// Use Cases
// Text Highlighting in PDF Viewers:

// Determining the area of the highlighted text in a PDF viewer based on the heights of the letters.
// Text Rendering in Graphics:

// Calculating the bounding box for rendering text in graphical applications.
// Font Size Adjustment:

// Adjusting the font size dynamically based on the height of the tallest letter in a word.
// Text Selection in Editors:

// Calculating the area of the selected text in text editors for highlighting purposes.
// UI/UX Design:

// Designing user interfaces where the height of text elements needs to be considered for layout purposes.

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n), where n is the length of the word.
// This is because the function iterates through each character in the word exactly once to find the maximum height.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the max variable and the loop variables.
// Summary
// Time Complexity: O(n) because the function iterates through each character in the word exactly once.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
