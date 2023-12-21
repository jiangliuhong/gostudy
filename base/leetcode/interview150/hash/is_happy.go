package hash

// 202. 快乐数
func isHappy(n int) bool {
	dict := map[int]bool{}
	for n != 1 {
		sum := 0
		for n > 0 {
			sum = sum + (n%10)*(n%10)
			n = n / 10
		}
		n = sum
		if dict[n] {
			break
		}
		dict[n] = true
	}
	return n == 1
}
