package main

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var sum1, sum2 int32            // 0
	for i := 0; i < len(arr); i++ { // 11, 5, -12 // 4, 5, 10 //
		sum1 += arr[i][i]            // 11 + 5 + -12 = 4 // 4 + 5 + 10 = 19
		sum2 += arr[i][len(arr)-1-i] // 4 + 5 + 10 = 19 // 11 + 5 + -12 = 4
	}
	return abs(sum1 - sum2) // 15
}

//Loop Execution:
// First Iteration (i = 0):

// sum1 += arr[0][0] → sum1 = 0 + 11 = 11
// sum2 += arr[0][2] → sum2 = 0 + 4 = 4
// Second Iteration (i = 1):

// sum1 += arr[1][1] → sum1 = 11 + 5 = 16
// sum2 += arr[1][1] → sum2 = 4 + 5 = 9
// Third Iteration (i = 2):

// sum1 += arr[2][2] → sum1 = 16 + (-12) = 4
// sum2 += arr[2][0] → sum2 = 9 + 10 = 19
// }

func abs(a int32) int32 { // 15
	if a < 0 { // false
		return -a
	}
	return a // 15
}

func test() { //
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	println(diagonalDifference(arr)) // 15
}

func main() {
	test() // run test cases
}

// ## Explanation:

// Here is a flowchart that represents the execution steps of the `for` loop in the

// diagonalDifference

//  function:

// ```plaintext
// Start
//  |
//  v
// Initialize sum1 and sum2 to 0
//  |
//  v
// i = 0
//  |
//  v
// Is i < len(arr)?
//  / \
// Yes  No
//  |    |
//  v    v
// Add arr[i][i] to sum1
//  |
//  v
// Add arr[i][len(arr)-1-i] to sum2
//  |
//  v
// Increment i (i = i + 1)
//  |
//  v
// Repeat loop
//  |
//  v
// Calculate abs(sum1 - sum2)
//  |
//  v
// Return the result
//  |
//  v
// End
// ```

// // ### Explanation of the Flowchart:

// // 1. **Start**: The function

// // diagonalDifference

// //  is called.
// // 2. **Initialize sum1 and sum2 to 0**: Two variables

// // sum1

// //  and

// // sum2

// //  are initialized to 0.
// // 3. **i = 0**: The loop variable

// // i

// //  is initialized to 0.
// // 4. **Is i < len(arr)?**: The condition of the `for` loop is checked.
// //    - **Yes**: If

// // i

// //  is less than the length of the array, the loop continues.
// //      - **Add arr[i][i] to sum1**: The element at the primary diagonal position

// // arr[i][i]

// //  is added to

// // sum1

// // .
// //      - **Add arr[i][len(arr)-1-i] to sum2**: The element at the secondary diagonal position

// // arr[i][len(arr)-1-i]

// //  is added to

// // sum2

// // .
// //      - **Increment i (i = i + 1)**: The loop variable

// // i

// //  is incremented by 1.
// //      - **Repeat loop**: The loop condition is checked again.
// //    - **No**: If

// // i

// //  is not less than the length of the array, the loop ends.
// // 5. **Calculate abs(sum1 - sum2)**: The absolute difference between

// // sum1

// //  and

// // sum2

// //  is calculated using the

// // abs

// //  function.
// // 6. **Return the result**: The result of the absolute difference is returned.
// // 7. **End**: The function execution ends.
