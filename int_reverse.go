package algos

// Reverse 数据倒序（十进制数 用余数 整数*10递归拼接）
func Reverse(x int) int {
	var recv int
	for x != 0 {
		pop := x % 10
		x /= 10
		if recv*10 > 1<<31-1 || (recv*10 == 1<<31-1 && pop > 7) {
			return 0
		}
		if recv*10 < -1<<31 || (recv*10 == -1<<31 && pop < 8) {
			return 0
		}
		recv = recv*10 + pop
	}
	return recv
}
