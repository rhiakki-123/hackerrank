package main

import "fmt"

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

// Time and Space Complexity
// Time Complexity: O(ranked + player)
// The time complexity of the function is O(ranked + player) because it iterates over the ranked array and the player array once.
// Space Complexity: O(ranked)
// The space complexity of the function is O(ranked) because it uses a uniqueRanked array to store the unique scores from the ranked array. The size of the uniqueRanked array is at most the size of the ranked array.

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	var uniqueRanked []int32
	for _, r := range ranked {
		if len(uniqueRanked) == 0 || uniqueRanked[len(uniqueRanked)-1] != r {
			uniqueRanked = append(uniqueRanked, r)
		}
	}
	var res []int32            // result array to store the ranks of the player
	j := len(uniqueRanked) - 1 // index of the last element in the uniqueRanked array
	for _, p := range player { // iterate over the player array to calculate the rank of each player
		for j >= 0 && p >= uniqueRanked[j] { // find the rank of the player by comparing their score with the scores in the uniqueRanked array
			j--
		}
		res = append(res, int32(j+2)) // add the rank of the player to the result array
	}
	return res
}

func test() {
	fmt.Println(climbingLeaderboard([]int32{100, 90, 90, 80, 75, 60}, []int32{50, 65, 77, 90, 102})) // [6, 5, 4, 2, 1]
}

func main() {
	test()
}

// Time and Space Complexity
// Time Complexity:
// The time complexity of the function is O(ranked + player) because it iterates over the ranked array and the player array once.
// Space Complexity:
// The space complexity of the function is O(ranked) because it uses a uniqueRanked array to store the unique scores from the ranked array. The size of the uniqueRanked array is at most the size of the ranked array.
// Real-Time Use Cases
// Leaderboard Ranking:
// Determining the rank of a player on a leaderboard based on their scores
// Competitive Programming:
// Calculating the rank of a participant in a programming contest based on their performance
// Online Gaming:
// Assigning a rank to players based on their scores in an online game
// Competitive Sports:
// Ranking athletes based on their performance in a sports competition
// Academic Grading:
// Assigning a rank to students based on their grades in a class or exam
// Job Interviews:
// Ranking candidates based on their performance in an interview process
// Sales Performance:
// Ranking sales representatives based on their sales performance
// Customer Satisfaction:
// Ranking customers based on their satisfaction scores
// Employee Performance:
// Ranking employees based on their performance evaluations
// Product Reviews:
// Ranking products based on customer reviews and ratings
// Social Media Influence:
// Ranking users based on their social media influence and engagement
// Music Charts:
// Ranking songs based on their popularity and sales
// Movie Ratings:
// Ranking movies based on their box office performance and critical reviews
// Book Rankings:
// Ranking books based on their sales and reader reviews
// Website Traffic:
// Ranking websites based on their traffic and engagement metrics
// Search Engine Results:
// Ranking search results based on relevance and popularity
// Financial Markets:
// Ranking financial assets based on their performance and risk factors
// Real Estate Markets:
// Ranking properties based on their location, size, and amenities
// Travel Destinations:
// Ranking travel destinations based on visitor reviews and attractions
// Restaurant Ratings:
// Ranking restaurants based on customer reviews and ratings
// Health and Fitness:
// Ranking individuals based on their fitness levels and health metrics
// Educational Institutions:
// Ranking schools and universities based on academic performance and student outcomes
// Job Market:
// Ranking job opportunities based on salary, location, and job requirements
// Investment Opportunities:
// Ranking investment options based on risk, return, and market conditions
// Business Performance:
// Ranking businesses based on revenue, growth, and market share
// Productivity Rankings:
// Ranking individuals or teams based on their productivity levels and output
// Environmental Rankings:
// Ranking countries or regions based on their environmental sustainability and conservation efforts
// Social Impact Rankings:
// Ranking organizations based on their social impact and community contributions
// Technology Rankings:
// Ranking technology companies based on innovation, market share, and product quality
// Healthcare Rankings:
// Ranking healthcare providers based on patient outcomes and quality of care
// Legal Rankings:
// Ranking law firms based on their expertise, reputation, and client satisfaction
// Political Rankings:
// Ranking political candidates based on their policies, leadership qualities, and public support
// Social Rankings:
// Ranking individuals or groups based on their social status, influence, and connections
// Cultural Rankings:
// Ranking cultural institutions or events based on their historical significance and popularity
// Entertainment Rankings:
// Ranking entertainment products or events based on their audience engagement and critical acclaim
// Fashion Rankings:
// Ranking fashion designers or brands based on their creativity, style, and market appeal
// Food Rankings:
// Ranking food products or restaurants based on taste, quality, and presentation
// Travel Rankings:
// Ranking travel destinations or experiences based on visitor reviews and recommendations
// Technology Rankings:
// Ranking technology products or companies based on innovation, performance, and user experience
