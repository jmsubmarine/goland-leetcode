/*

Given an integer num, return a string of its base 7 representation.

*/

package main

func convertToBase7(num int) string {
    if num == 0 {
		return "0"
	}

    negative := num < 0
	if negative {
		num = -num
	}

    digits := ""
	for num > 0 {
		remainder := num % 7
		digits = strconv.Itoa(remainder) + digits
		num /= 7
	}

	if negative {
		digits = "-" + digits
	}

	return digits
}
