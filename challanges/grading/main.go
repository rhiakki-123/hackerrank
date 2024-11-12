package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	var result []int32

	for _, grade := range grades {
		if grade >= 38 && grade%5 >= 3 {
			result = append(result, grade+(5-grade%5))
		} else {
			result = append(result, grade)
		}
	}
	return result

}

func test() {
	grades := []int32{73, 67, 38, 33}
	result := gradingStudents(grades)
	fmt.Println(result) // Expected output: [75 67 40 33]
}

func main() {
	test()
}
