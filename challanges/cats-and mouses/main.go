package main

// Two cats and a mouse are at various positions on a line. You will be given their starting positions.
// Your task is to determine which cat will reach the mouse first, assuming the mouse does not move and the
// cats travel at equal speed. If the cat arrive at the same time, the mouse will be allowed to move and it
// will escape while they fight.

// If cat A catches the mouse first, print Cat A.
// If cat B catches the mouse first, print Cat B.
// If both cats reach the mouse at the same time, print Mouse C as the two cats fight and mouse escapes.

// Time Complexity: O(1) because the function performs a constant number of operations.
// Space Complexity: O(1) because the function uses a constant amount of extra space.

func catAndMouse(x int32, y int32, z int32) string {
	// Write your code here
	if abs(x-z) < abs(y-z) { // abs() is a function that returns the absolute value of a number (int32)
		return "Cat A" // if the absolute value of x-z is less than the absolute value of y-z, return "Cat A"
	} else if abs(x-z) > abs(y-z) { // if the absolute value of x-z is greater than the absolute value of y-z, return "Cat B"
		return "Cat B"
	} else {
		return "Mouse C"
	}
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func catAndMouse1(x int32, y int32, z int32) string {
	// Write your code here
	if x == z && y == z {
		return "Mouse C"
	}
	if abs(x-z) < abs(y-z) {
		return "Cat A"
	} else if abs(x-z) > abs(y-z) {
		return "Cat B"
	} else {
		return "Mouse C"
	}
}

func test() {
	println(catAndMouse(1, 2, 3)) // Cat B
	println(catAndMouse(1, 3, 2)) // Mouse C
	println(catAndMouse(2, 1, 3)) // Cat A
}

func main() {
	test()
}

// Start
//  |
//  v
// Calculate abs(x - z)
//  |
//  v
// Calculate abs(y - z)
//  |
//  v
// Is abs(x - z) < abs(y - z)?
//  / \
// Yes  No
//  |    |
//  v    v
// Return "Cat A"  Is abs(x - z) > abs(y - z)?
//                 / \
//               Yes  No
//                |    |
//                v    v
//          Return "Cat B"  Return "Mouse C"
