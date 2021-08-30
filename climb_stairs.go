package algos

func climbStairs(n int)  int{
	p := 0; q := 0; r := 1

	for i:=1; i <= n; i++{
		p = q
		q = r
		r = p + q
	}
	return r
}
