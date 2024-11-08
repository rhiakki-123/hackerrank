package main

// Your local library needs your help! Given the expected and actual return dates for a library book, create a program that calculates the fine (if any). The fee structure is as follows:

// If the book is returned on or before the expected return date, no fine will be charged (i.e.: fine = 0).
// If the book is returned after the expected return day but still within the same calendar month and year as
// the expected return date, fine = 15 Hackos × the number of days late.
// If the book is returned after the expected return month but still within the same calendar year as the expected
// return date, the fine = 500 Hackos × the number of months late.
// If the book is returned after the calendar year in which it was expected, there is a fixed fine of 10000 Hackos.
// Charges are based only on the least precise measure of lateness. For example, whether a book is due
// January 1, 2017 or December 31, 2017, if it is returned January 1, 2018, that is a year late and the fine would be
// 10000 Hackos.

// d1 = 9, m1 = 6, y1 = 2015
// d2 = 6, m2 = 6, y2 = 2015

// The first values are the return date and the second are the due date. The result is 45.

/*
 * Complete the 'libraryFine' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER d1
 *  2. INTEGER m1
 *  3. INTEGER y1
 *  4. INTEGER d2
 *  5. INTEGER m2
 *  6. INTEGER y2
 */

// Code Implementation
func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	if y1 > y2 {
		return 10000
	} else if y1 == y2 {
		if m1 > m2 {
			return 500 * (m1 - m2)
		} else if m1 == m2 {
			if d1 > d2 {
				return 15 * (d1 - d2)
			}
		}
	}
	return 0
}

func libraryFine1(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	// Write your code here
	if y1 > y2 {
		return 10000
	}
	if y1 == y2 && m1 > m2 {
		return 500 * (m1 - m2)
	}
	if y1 == y2 && m1 == m2 && d1 > d2 {
		return 15 * (d1 - d2)
	}
	return 0
}

func test() {
	d1 := int32(9)
	m1 := int32(6)
	y1 := int32(2015)
	d2 := int32(6)
	m2 := int32(6)
	y2 := int32(2015)
	println(libraryFine(d1, m1, y1, d2, m2, y2)) // 45

	d11 := int32(6)
	m11 := int32(6)
	y11 := int32(2015)
	d21 := int32(9)
	m21 := int32(6)
	y21 := int32(2016)
	println(libraryFine(d11, m11, y11, d21, m21, y21))
}

func main() {
	test()
}

//explain function  libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 with its logic, in details execution, algo, flowchart, timeand space complexity and use cases

// Explanation of libraryFine Function
// The libraryFine function calculates the fine for returning a library book late. The fine is calculated based on the number of days, months, and years the book is returned after the due date.

// Function Signature
// func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32

// d1, m1, y1: The return date (day, month, year).
// d2, m2, y2: The due date (day, month, year).

// Logic
// If the book is returned on or before the due date, the fine is 0.
// If the book is returned after the due date but within the same month and year, the fine is 15 Hackos for each day late.
// If the book is returned after the due month but within the same year, the fine is 500 Hackos for each month late.
// If the book is returned after the due year, the fine is a fixed 10,000 Hackos.

// Algorithm
// Check if the return year (y1) is greater than the due year (y2):
// If true, return 10,000 Hackos.
// Check if the return year (y1) is equal to the due year (y2):
// If true, check if the return month (m1) is greater than the due month (m2):
// If true, return 500 Hackos for each month late.
// If the return month (m1) is equal to the due month (m2):
// Check if the return day (d1) is greater than the due day (d2):
// If true, return 15 Hackos for each day late.
// If none of the above conditions are met, return 0 Hackos.

// Flowchart
// Start
//  |
//  v
// Is y1 > y2?
//  / \
// Yes  No
//  |    |
//  v    v
// Return 10000  Is y1 == y2?
//               / \
//             Yes  No
//              |    |
//              v    v
//       Is m1 > m2?  Return 0
//       / \
//     Yes  No
//      |    |
//      v    v
// Return 500 * (m1 - m2)  Is m1 == m2?
//                         / \
//                       Yes  No
//                        |    |
//                        v    v
//                 Is d1 > d2?  Return 0
//                 / \
//               Yes  No
//                |    |
//                v    v
//         Return 15 * (d1 - d2)  Return 0
//  |
//  v
// End

// Detailed Execution
// Given the inputs:

// d1 := int32(9)
// m1 := int32(6)
// y1 := int32(2015)
// d2 := int32(6)
// m2 := int32(6)
// y2 := int32(2015)

// Check if y1 > y2:
// y1 > y2 // 2015 > 2015 // false

// Check if y1 == y2:
// y1 == y2 // 2015 == 2015 // true

// Check if m1 > m2:
// m1 > m2 // 6 > 6 // false

// Check if m1 == m2:
// m1 == m2 // 6 == 6 // true

// Check if d1 > d2:
// d1 > d2 // 9 > 6 // true

// Return 15 * (d1 - d2):
// 15 * (9 - 6) // 15 * 3 // 45

// Final State:
// d1 = 9
// m1 = 6
// y1 = 2015
// d2 = 6
// m2 = 6
// y2 = 2015
// Output:
// 45

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(1).
// This is because the function performs a constant number of operations regardless of the input size.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the fine variable.

// Real-Time Use Cases
// Library Management:

// Calculating fines for overdue books in a library system.
// Rental Services:

// Determining late fees for rental services such as car rentals, equipment rentals, etc.
// Subscription Services:

// Calculating penalties for late payments in subscription-based services.

// GitHub Copilot
// Explanation of libraryFine Function
// The libraryFine function calculates the fine for returning a library book late. The fine is calculated based on the number of days, months, and years the book is returned after the due date.

// Function Signature
// d1, m1, y1: The return date (day, month, year).
// d2, m2, y2: The due date (day, month, year).
// Logic
// If the book is returned on or before the due date, the fine is 0.
// If the book is returned after the due date but within the same month and year, the fine is 15 Hackos for each day late.
// If the book is returned after the due month but within the same year, the fine is 500 Hackos for each month late.
// If the book is returned after the due year, the fine is a fixed 10,000 Hackos.
// Algorithm
// Check if the return year (y1) is greater than the due year (y2):
// If true, return 10,000 Hackos.
// Check if the return year (y1) is equal to the due year (y2):
// If true, check if the return month (m1) is greater than the due month (m2):
// If true, return 500 Hackos for each month late.
// If the return month (m1) is equal to the due month (m2):
// Check if the return day (d1) is greater than the due day (d2):
// If true, return 15 Hackos for each day late.
// If none of the above conditions are met, return 0 Hackos.
// Flowchart
// Detailed Execution
// Given the inputs:

// Check if y1 > y2:

// Check if y1 == y2:

// Check if m1 > m2:

// Check if m1 == m2:

// Check if d1 > d2:

// Calculate the fine:

// Final State:
// fine = 45
// Output:
// 45
// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(1).
// This is because the function performs a constant number of operations regardless of the input size.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the fine variable.
// Real-Time Use Cases
// Library Management:

// Calculating fines for overdue books in a library system.
// Rental Services:

// Determining late fees for rental services such as car rentals, equipment rentals, etc.
// Subscription Services:

// Calculating penalties for late payments in subscription-based services.
// Code Implementation
// Summary
// Logic: Calculate the fine based on the number of days, months, and years the book is returned after the due date.
// Algorithm: Check if the return year is greater than the due year, if the return month is greater than the due month, and if the return day is greater than the due day, and calculate the fine accordingly.
// Flowchart: Visual representation of the decision-making process for calculating the fine for returning a library book late.
// Time Complexity: O(1) because the function performs a constant number of operations regardless of the input size.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Library management, rental services, subscription services.
// This detailed explanation, flowchart, complexity analysis, and use cases show how the libraryFine function works step-by-step to calculate the fine for returning a library book late.
