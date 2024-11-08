package main

import (
	"fmt"
)

// Given a time in -hour AM/PM format, convert it to military (24-hour) time.

// Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
// - 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

// s = "12:01:00PM"
// Return '12:01:00'.
// s="12:01:00AM"
// Return '00:01:00'.

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	if s[len(s)-2:] == "AM" { // 12:01:00AM
		if s[:2] == "12" { // 12:01:00
			return "00" + s[2:len(s)-2] // 00:01:00
		}
		return s[:len(s)-2] // 12:01:00
	} else { // 12:01:00PM
		if s[:2] == "12" { // 12:01:00
			return s[:len(s)-2] // 12:01:00
		}
		return fmt.Sprintf("%02d", (int(s[0]-'0')*10+int(s[1]-'0'))+12) + s[2:len(s)-2] // 12:01:00
	}
}

func test() {
	fmt.Println(timeConversion("12:01:00PM")) // 12:01:00
	fmt.Println(timeConversion("12:01:00AM")) // 00:01:00
	fmt.Println(timeConversion("07:51:00PM")) // 12:01:00
	fmt.Println(timeConversion("09:01:00AM")) // 00:01:00
}

func main() {
	test()
}
