package palindromenumber

func isPalindrome(x int) bool {
	y, z := 0, x

	for x > 0 {
		y = y*10 + x%10
		x /= 10
	}

	return y == z
}
