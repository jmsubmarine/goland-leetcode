/*
For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""
*/

package main

func gcdOfStrings(str1 string, str2 string) string {
    // Check if str1 and str2 can be formed from the same substring.
    // If str1 + str2 is not equal to str2 + str1, then the strings cannot be divided into the same parts.
    if str1 + str2 != str2 + str1 {
        return ""
    }

    // Find the greatest common divisor (GCD) of the lengths of str1 and str2.
    x := gcd(len(str1), len(str2))
    return str1[:x]
}

// gcd finds the greatest common divisor (GCD) of two integers a and b using the Euclidean algorithm.
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}
