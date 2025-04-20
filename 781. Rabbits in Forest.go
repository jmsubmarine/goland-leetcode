/*

There is a forest with an unknown number of rabbits. We asked n rabbits "How many rabbits have the same color as you?" and collected the answers in an integer array answers where answers[i] is the answer of the ith rabbit.

Given the array answers, return the minimum number of rabbits that could be in the forest.

Example 1:

Input: answers = [1,1,2]
Output: 5

Explanation:
The two rabbits that answered "1" could both be the same color, say red.
The rabbit that answered "2" can't be red or the answers would be inconsistent.
Say the rabbit that answered "2" was blue.
Then there should be 2 other blue rabbits in the forest that didn't answer into the array.
The smallest possible number of rabbits in the forest is therefore 5: 3 that answered plus 2 that didn't.

Example 2:
Input: answers = [10,10,10]
Output: 11

Constraints:

1 <= answers.length <= 1000
0 <= answers[i] < 1000

*/

package main

func numRabbits(answers []int) int {
    // Create a map to count how many times each answer appears
    count := map[int]int{}

    // Iterate through the answers and count their occurrences
    for _, answer := range answers {
        count[answer]++
    }

    total := 0

    // Iterate over the map to calculate the total number of rabbits
    for answerValue, answerCount := range count {
        // Calculate the group size (answer + 1)
        groupSize := answerValue + 1
        // Calculate the number of full groups needed to fit all rabbits with this answer
        groups := (answerCount + groupSize - 1) / groupSize
        // Add the total number of rabbits in these groups to the result
        total += groups * groupSize
    }

    return total
}
