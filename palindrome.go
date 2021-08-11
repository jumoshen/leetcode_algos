package algos

func IsPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	var right int

	for x > right {
		right = right*10 + x%10
		x /= 10
	}

	return x == right || x == right/10
}
