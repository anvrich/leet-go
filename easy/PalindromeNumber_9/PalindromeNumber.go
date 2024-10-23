package main

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	origin, reversed := x, 0

	for x != 0 {
		lastDigit := x % 10
		reversed = reversed*10 + lastDigit
		x = x / 10
	}
	return origin == reversed
}
