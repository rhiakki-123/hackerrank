package main

import (
	"fmt"
	"math"
	"strings"
)

// An English text needs to be encrypted using the following encryption scheme.
// First, the spaces are removed from the text. Let L be the length of this text.
// Then, characters are written into a grid, whose rows and columns have the following constraints:

//   floor(sqrt(L)) <= row <= column <= ceil(sqrt(L))

// For example
//  s = if man was meant to stay on the ground god would have given us roots

// After removing spaces, the string is 54 characters long. sqrt(54) is between 7 and 8,
// so it is written in the form of a grid with 7 rows and 8 columns.

//  ifmanwas
//  meanttos
//  tayonthe
//  groundgo
//  dwouldha
//  vegivenu
//  sroots

// Ensure that rows x columns >= L

// If multiple grids satisfy the above conditions, choose the one with the minimum area, i.e. rows x columns.

// The encoded message is obtained by displaying the characters in a column, inserting a space, and then displaying the next column and inserting a space, and so on.
// For example, the encoded message for the above rectangle is:

// imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn sseoau

// You will be given a message to encode and print the encoded message.

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func encryption(s string) string {
	// Write your code here
	s = strings.ReplaceAll(s, " ", "")
	l := len(s)
	row := int32(math.Floor(math.Sqrt(float64(l)))) //
	col := int32(math.Ceil(math.Sqrt(float64(l))))
	if row*col < int32(l) {
		row++
	}
	var result string
	for i := int32(0); i < col; i++ {
		for j := int32(0); j < row; j++ {
			if int32(j)*col+int32(i) < int32(l) { // if the index is less than the length of the string
				result += string(s[int32(j)*col+int32(i)]) // add the character to the result
			}
		}
		result += " "

		//remove the last space
		if i == col-1 {
			result = result[:len(result)-1]
		}
	}
	return result
}

func test() {
	// Test case
	s := "if man was meant to stay on the ground god would have given us roots"
	result := encryption(s)
	fmt.Println(result) // Expected output: "tcte eht saes"

}

func main() {
	test()
}

// Explain function func encryption(s string) string
// with its in details execution, algo, flowchart, time
// and space complexity, write its optimized version, unit tests and use cases

// Explanation of encryption Function
// The encryption function encrypts a given string s using a specific algorithm.
// The algorithm involves arranging the characters of the string into a grid
// and then reading the characters column-wise to form the encrypted message.

// Function Signature
// func encryption(s string) string

// Parameters:
// s: The input string to be encrypted.
// Returns:
// A string representing the encrypted message.

// Logic
// Remove Spaces:
// Remove all spaces from the input string s.
// Determine Grid Dimensions:
// Calculate the length of the modified string s.
// Determine the number of rows and columns for the grid based on the length of the string.
// Rows: rows = floor(sqrt(length))
// Columns: cols = ceil(sqrt(length))
// Ensure that rows * cols is at least the length of the string.
// Fill the Grid:
// Fill the grid row-wise with the characters of the string.
// Read Column-wise:
// Read the characters column-wise to form the encrypted message.

// Algorithm
// Remove spaces from the input string s.
// s = strings.ReplaceAll(s, " ", "") // Remove spaces from the string s.

// Determine Grid Dimensions:
// length := len(s)
// rows := int(math.Floor(math.Sqrt(float64(length)))) // Calculate the number of rows.
// cols := int(math.Ceil(math.Sqrt(float64(length)))) // Calculate the number of columns.
// if rows * cols < length { // Ensure that rows * cols is at least the length of the string.
//     rows++
// }

// Fill the Grid and Read Column-wise:
// Initialize a slice of strings to store the encrypted message.
// Iterate over the columns and construct the encrypted message by reading characters column-wise.

// Start
//  |
//  v
// Remove spaces from s
//  |
//  v
// Calculate length of s
//  |
//  v
// Determine rows and cols
//  |
//  v
// If rows * cols < length, increment rows
//  |
//  v
// Initialize encryptedMessage slice
//  |
//  v
// For each column from 0 to cols-1
//  |  |
//  |  v
//  |  For each row from 0 to rows-1
//  |    Add character to encryptedMessage
//  |
//  v
// Join encryptedMessage with spaces
//  |
//  v
// Return encryptedMessage
//  |
//  v
// End

// Detailed Explanation
// Example Input
//  s := "have a nice day"

//  Step-by-Step Execution:
// 1. Remove spaces from the input string s.
//     s = "haveaniceday"
// 2. Calculate the length of the modified string s.
//     length = 12
// 3. Determine the number of rows and columns for the grid.
//     rows = int(math.Floor(math.sqrt(float64(13)))) = 3
//     cols = int(math.Floor(math.ceil(sqrt(13)))) = 4
// if 3 * 4 < 12, increment rows to 4.
// 4. Fill the grid row-wise with the characters of the string.
// Initialize encryptedMessage as an empty slice of strings.
// Iterate over the columns and construct the encrypted message
// encryptedMessage := []string{}
// for c := 0; c < cols; c++ {
//     var colString string
//     for r := 0; r < rows; r++ {
//         index := r * cols + c
//         if index < length {
//             colString += string(s[index])
//         }
//     }
//     encryptedMessage = append(encryptedMessage, colString)
// }

// 5. Read the characters column-wise to form the encrypted message.
// Join the strings in the encryptedMessage slice with spaces to form the final encrypted message.
// encryptedMessageString := strings.Join(encryptedMessage, " ")

// Output:
// encryptedMessageString = "hae and via ecy"

// Time and Space Complexity
// Time Complexity
// Removing Spaces: O(n)
// Iterates through the string to remove spaces.
// Calculating Dimensions: O(1)
// Constant time operations to calculate rows and columns.
// Filling the Grid and Reading Column-wise: O(n)
// Iterates through the string to fill the grid and read column-wise.
// Overall Time Complexity: O(n)
// Space Complexity
// Additional Space: O(n)
// Stores the encrypted message and intermediate strings.
// Overall Space Complexity: O(n

// Optimized Version
// The encryption function can be optimized by combining the steps of filling the grid and reading column-wise into a single step.
// This optimization reduces the number of iterations and improves the efficiency of the algorithm.

// package main

// import (
//     "fmt"
//     "math"
//     "strings"
// )

// func encryption(s string) string {
//     s = strings.ReplaceAll(s, " ", "")
//     length := len(s)
//     rows := int(math.Floor(math.Sqrt(float64(length))))
//     cols := int(math.Ceil(math.Sqrt(float64(length))))
//     if rows*cols < length {
//         rows++
//     }

//     encryptedMessage := []string{}
//     for c := 0; c < cols; c++ {
//         var colString string
//         for r := 0; r < rows; r++ {
//             index := r*cols + c
//             if index < length {
//                 colString += string(s[index])
//             }
//         }
//         encryptedMessage = append(encryptedMessage, colString)
//     }

//     return strings.Join(encryptedMessage, " ")
// }

// func testEncryption() {
//     s := "have a nice day"
//     fmt.Println(encryption(s)) // Output: "hae and via ecy"
// }

// func main() {
//     testEncryption()
// }

// package main

// import "testing"

// func TestEncryption(t *testing.T) {
//     tests := []struct {
//         input    string
//         expected string
//     }{
//         {"have a nice day", "hae and via ecy"},
//         {"feed the dog", "fto ehg ee dd"},
//         {"chill out", "clu hlt io"},
//         {"if man was meant to stay on the ground god would have given us roots", "imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn sseoau"},
//         {"", ""},
//     }

//     for _, test := range tests {
//         result := encryption(test.input)
//         if result != test.expected {
//             t.Errorf("For input %q, expected %q but got %q", test.input, test.expected, result)
//         }
//     }
// }

// Summary
// Logic: Encrypt the input string by arranging characters into a grid and reading them column-wise.
// Algorithm: Remove spaces, calculate grid dimensions, fill the grid, read column-wise, and return the encrypted message.
// Flowchart: Visual representation of the encryption process.
// Time Complexity: O(n) because the function iterates through the string to remove spaces, fill the grid, and read column-wise.
// Space Complexity: O(n) because the function uses additional space to store the encrypted message and intermediate strings.
// Use Cases: Secure communication, data obfuscation, text transformation.
