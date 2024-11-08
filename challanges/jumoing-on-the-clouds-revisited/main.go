package main

import "fmt"

// A child is playing a cloud hopping game. In this game, there are sequentially numbered clouds that can be thunderheads or
// cumulus clouds. The character must jump from cloud to cloud until it reaches the start again.
// There is an array of clouds, c and an energy level e=100. The character starts from c[0] and uses 1 unit of energy to make
// a jump of size k to cloud c[(i+k)%n].
// If it lands on a thundercloud, c[i]=1, its energy (e) decreases by 2 additional units. The game ends when the character lands
// back on cloud 0.

func jumpingOnClouds(c []int32, k int32) int32 {
	// Write your code here
	var energy int32 = 100
	var i int32 = 0
	for {
		i = (i + k) % int32(len(c))
		energy--
		if c[i] == 1 {
			energy -= 2
		}
		if i == 0 {
			break
		}
	}
	return energy
}

func jumpingOnClouds1(c []int32, k int32) int32 {
	// Write your code here
	var energy int32 = 100
	var i int32 = 0
	for {
		i = (i + k) % int32(len(c))
		energy--
		if c[i] == 1 {
			energy -= 2
		}
		if i == 0 {
			break
		}
	}
	return energy

}

func test() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0}, 2))              // 96
	fmt.Println(jumpingOnClouds1([]int32{0, 0, 1, 0, 1, 0, 1, 1}, 2)) //90
}

func test1() {
	c := []int32{0, 1, 0, 0, 0, 1, 0}
	k := int32(2)
	fmt.Println(jumpingOnClouds1(c, k)) // Output: 3
}

func main() {
	test()
	test1()
}

//explain function jumpingOnClouds with its logic, in details execution, algo, flowchart, timeand space complexity and use cases
// Jumping on the Clouds Revisited

// Explanation of jumpingOnClouds Function
// The jumpingOnClouds function simulates a game where a player jumps on clouds arranged in a circular array. The player starts with a certain amount of energy and loses energy with each jump. The goal is to determine the players remaining energy after completing the game.

// Function Signature

// func jumpingOnClouds(c []int32, k int32) int32

// c: An array of integers representing the clouds. 0 indicates a safe cloud, and 1 indicates a thundercloud that reduces additional energy.
// k: The number of clouds the player jumps over in one jump.

// Logic
// Initialize a variable energy to 100 to represent the players starting energy.
// Initialize a variable i to 0 to represent the current position.
// Use a for loop to simulate the jumps:
// Calculate the new position using (i + k) % len(c).
// Decrement energy by 1 for each jump.
// If the new position is a thundercloud (c[i] == 1), decrement energy by an additional 2.
// If the new position is the starting position (i == 0), break the loop.
// Return the remaining energy.

// Algorithm
// Initialize energy to 100 and i to 0.
// Use a for loop to simulate the jumps:
// Calculate the new position using (i + k) % len(c).
// Decrement energy by 1 for each jump.
// If the new position is a thundercloud (c[i] == 1), decrement energy by an additional 2.
// If the new position is the starting position (i == 0), break the loop.
// Return the remaining energy.

// Flowchart
// Start
//  |
//  v
// Initialize energy to 100
// Initialize i to 0
//  |
//  v
// For each jump
//  |  |
//  |  v
//  |  Calculate i as (i + k) % len(c)
//  |  |
//  |  v
//  |  Decrement energy by 1
//  |  |
//  |  v
//  |  Is c[i] == 1?
//  |  / \
//  | Yes  No
//  |  |    |
//  |  v    v
//  | Decrement energy by 2
//  |  |
//  v  Is i == 0?
//  / \
// Yes  No
//  |    |
//  v    v
// Break loop  Repeat loop
//  |
//  v
// Return energy
//  |
//  v
// End

// Detailed Execution
// Given the input:

// c := []int32{0, 1, 0, 0, 0, 1, 0}
// k := int32(2)

// Initialize variables:
// energy := int32(100)
// i := int32(0)

// First Jump:
// Calculate new position
// i = (i + k) % len(c) = (0 + 2) % 7 = 2
// Decrement energy by 1
// energy = 100 - 1 = 99
// c[i] == 0: No additional energy decrement.

// Second Jump:
// Calculate new position
// i = (i + k) % len(c) = (2 + 2) % 7 = 4
// Decrement energy by 1
// energy = 99 - 1 = 98
// c[i] == 0: No additional energy decrement.

// Third Jump:
// Calculate new position
// i = (i + k) % len(c) = (4 + 2) % 7 = 6
// Decrement energy by 1
// energy = 98 - 1 = 97
// c[i] == 0: No additional energy decrement.

// Fourth Jump:
// Calculate new position
// i = (i + k) % len(c) = (6 + 2) % 7 = 1
// Decrement energy by 1
// energy = 97 - 1 = 96
// c[i] == 1: Additional energy decrement by 2.
// energy = 96 - 2 = 94

// Fifth Jump:
// Calculate new position
// i = (i + k) % len(c) = (1 + 2) % 7 = 3
// Decrement energy by 1
// energy = 94 - 1 = 93
// c[i] == 0: No additional energy decrement.

// Sixth Jump:
// Calculate new position
// i = (i + k) % len(c) = (3 + 2) % 7 = 5
// Decrement energy by 1
// energy = 93 - 1 = 92
// c[i] == 1: Additional energy decrement by 2.
// energy = 92 - 2 = 90

// Seventh Jump:
// Calculate new position
// i = (i + k) % len(c) = (5 + 2) % 7 = 0
// Decrement energy by 1
// energy = 90 - 1 = 89
// c[i] == 0: No additional energy decrement.

// Return remaining energy: 89

// Time and Space Complexity
// Time Complexity
// The time complexity of the function is O(n/k), where n is the length of the array c and k is the jump length.
// This is because the function iterates through the array c in steps of k.
// Space Complexity
// The space complexity of the function is O(1).
// This is because the function uses a constant amount of extra space regardless of the input size. The only additional storage used is for the energy and i variables.

// Real-Time Use Cases
// Game Development:

// Implementing energy mechanics in games where the player must avoid obstacles and manage energy levels.
// Robotics:

// Programming robots to navigate through environments with obstacles while managing energy consumption.
// Simulation:

// Simulating scenarios where entities move through a circular path with varying energy costs.
// Network Routing:

// Determining the energy consumption of data packets moving through a circular network with varying energy costs at different nodes.

// Summary
// Logic: Iterate through the array c in steps of k, decrementing energy for each jump and additional energy for thunderclouds, and return the remaining energy after completing the game.
// Algorithm: Initialize energy and i, use a for loop to simulate the jumps, decrement energy based on the conditions, and return the remaining energy.
// Flowchart: Visual representation of the decision-making process for calculating the remaining energy after completing the game.
// Time Complexity: O(n/k) because the function iterates through the array c in steps of k.
// Space Complexity: O(1) because the function uses a constant amount of extra space regardless of the input size.
// Use Cases: Game development, robotics, simulation, network routing.
