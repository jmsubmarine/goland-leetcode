/*
Given an integer array arr, return true if there are three consecutive odd numbers in the array. Otherwise, return false.

Example 1:

Input: arr = [2,6,4,1]
Output: false
Explanation: There are no three consecutive odds.
Example 2:

Input: arr = [1,2,34,3,4,5,7,23,12]
Output: true
*/

package main

func threeConsecutiveOdds(arr []int) bool {
    counter := 0

    for _, item := range arr {
        if item % 2 != 0 {
            counter++
        } else {
            counter = 0
        }

        if counter >= 3 {
            return true
        }
    }

    return false
}
