package main

import "fmt"

// Given the time in numerals we may convert it into words, as shown below:

// 5:00 -> five o clock
// 5:01 -> one minute past five
// 5:10 -> ten minutes past five
// 5:30 -> half past five
// 5:40 -> twenty minutes to six
// 5:45 -> quarter to six
// 5:47 -> thirteen minutes to six
// 5:28 -> twenty eight minutes past five

// At minutes = 0, use o clock.
// For 1 <= minutes <= 30,
// use past, and for 30 < minutes use to.
// Note the space between the apostrophe and clock in o clock.
// Write a program that prints the time in words for the input given.

//optimal solution
func timeInWords(h int32, m int32) string {
	numberWords := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "quarter", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", "twenty one", "twenty two", "twenty three", "twenty four", "twenty five", "twenty six", "twenty seven", "twenty eight", "twenty nine", "half"}

	if m == 0 {
		return fmt.Sprintf("%s o' clock", numberWords[h])
	} else if m <= 30 {
		if m == 15 || m == 30 {
			return fmt.Sprintf("%s past %s", numberWords[m], numberWords[h])
		} else if m == 1 {
			return fmt.Sprintf("one minute past %s", numberWords[h])
		} else {
			return fmt.Sprintf("%s minutes past %s", numberWords[m], numberWords[h])
		}
	} else {
		nextHour := h + 1
		if nextHour == 13 {
			nextHour = 1
		}
		remainingMinutes := 60 - m
		if remainingMinutes == 15 {
			return fmt.Sprintf("quarter to %s", numberWords[nextHour])
		} else if remainingMinutes == 1 {
			return fmt.Sprintf("one minute to %s", numberWords[nextHour])
		} else {
			return fmt.Sprintf("%s minutes to %s", numberWords[remainingMinutes], numberWords[nextHour])
		}
	}
}

func timeInWords1(h int32, m int32) string {
	// Write your code here
	hours := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"}
	minutes := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "quarter", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", "twenty one", "twenty two", "twenty three", "twenty four", "twenty five", "twenty six", "twenty seven", "twenty eight", "twenty nine", "half"}

	if m == 0 {
		return hours[h-1] + " o' clock"
	} else if m == 1 {
		return minutes[m-1] + " minute past " + hours[h-1]
	} else if m == 15 {
		return minutes[m-1] + " past " + hours[h-1]
	}
	if m < 30 {
		return minutes[m-1] + " minutes past " + hours[h-1]
	} else if m == 30 {
		return minutes[m-1] + " past " + hours[h-1]
	}
	if m == 59 {
		return minutes[60-m-1] + " minute to " + hours[h]
	}
	return minutes[60-m-1] + " minutes to " + hours[h]
}

func testTimeInWords() {
	// Test case
	h := int32(5)
	m := int32(45)
	result := timeInWords(h, m)
	fmt.Println(result) // Expected output: thirteen minutes to six
}

func main() {
	testTimeInWords()
}

// Explain in depth  function func timeInWords(h int32, m int32) string
// Logic
// Flowchart
// loop execution for each element in the array
// Algorithm
// Time Complexity
// Space Complexity
// Optimized version
// unit tests
// use cases

// Function Purpose
// timeInWords converts a given time in hours (h) and minutes (m) into its equivalent English words representation.

// Logic
// Define arrays for the words representing numbers and special cases for minutes.
// Handle different cases based on the value of m:
// If m is 0, return "o' clock".
// If m is between 1 and 30, return the appropriate phrase.
// If m is greater than 30, return the phrase for the next hour minus the remaining minutes.

// Flowchart
// graph TD
//     A[Start] --> B[Input: h, m]
//     B --> C[Define numberWords and minuteWords arrays]
//     C --> D[If m == 0]
//     D --> E[Return "o' clock"]
//     D --> F[Else if m <= 30]
//     F --> G[Return appropriate phrase]
//     F --> H[Else]
//     H --> I[Return phrase for next hour minus remaining minutes]
//     I --> J[End]

// Algorithm
// func timeInWords(h int32, m int32) string {
//     numberWords := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "quarter", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", "twenty one", "twenty two", "twenty three", "twenty four", "twenty five", "twenty six", "twenty seven", "twenty eight", "twenty nine", "half"}

//     if m == 0 {
//         return fmt.Sprintf("%s o' clock", numberWords[h])
//     } else if m <= 30 {
//         if m == 15 || m == 30 {
//             return fmt.Sprintf("%s past %s", numberWords[m], numberWords[h])
//         } else if m == 1 {
//             return fmt.Sprintf("one minute past %s", numberWords[h])
//         } else {
//             return fmt.Sprintf("%s minutes past %s", numberWords[m], numberWords[h])
//         }
//     } else {
//         nextHour := h + 1
//         if nextHour == 13 {
//             nextHour = 1
//         }
//         remainingMinutes := 60 - m
//         if remainingMinutes == 15 {
//             return fmt.Sprintf("quarter to %s", numberWords[nextHour])
//         } else if remainingMinutes == 1 {
//             return fmt.Sprintf("one minute to %s", numberWords[nextHour])
//         } else {
//             return fmt.Sprintf("%s minutes to %s", numberWords[remainingMinutes], numberWords[nextHour])
//         }
//     }
// }

// Loop Execution for Each Element
// There are no explicit loops in this function. The function uses conditional statements to determine the correct phrase based on the value of m.
// Time Complexity
// The time complexity is O(1) because the function performs a constant number of operations regardless of the input size.
// Space Complexity
// The space complexity is O(1) because the function uses a fixed amount of extra space for the numberWords array and a few variables.

// unit tests
// func TestTimeInWords(t *testing.T) {
//     tests := []struct {
//         h, m int32
//         want string
//     }{
//         {5, 0, "five o' clock"},
//         {5, 1, "one minute past five"},
//         {5, 10, "ten minutes past five"},
//         {5, 15, "quarter past five"},
//         {5, 30, "half past five"},
//         {5, 40, "twenty minutes to six"},
//         {5, 45, "quarter to six"},
//         {5, 47, "thirteen minutes to six"},
//         {12, 0, "twelve o' clock"},
//         {12, 59, "one minute to one"},
//     }

//     for _, tt := range tests {
//         got := timeInWords(tt.h, tt.m)
//         if got != tt.want {
//             t.Errorf("timeInWords(%v, %v) = %v; want %v", tt.h, tt.m, got, tt.want)
//         }
//     }
// }

// Use Cases
// Converting digital time to a more human-readable format.
// Applications where time needs to be displayed in words for better user experience.
// Educational tools for teaching children how to tell time.
// Voice assistants that need to verbalize the time.
