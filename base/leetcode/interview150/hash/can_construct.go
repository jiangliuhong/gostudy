package hash

// 383. 赎金信
// 给你两个字符串：ransomNote 和 magazine ，
// 判断 ransomNote 能不能由 magazine 里面的字符构成。
func canConstruct(ransomNote string, magazine string) bool {
	rd := map[rune]int{}
	for _, r := range ransomNote {
		rd[r]++
	}
	for _, m := range magazine {
		rd[m]--
		if rd[m] <= 0 {
			delete(rd, m)
		}
	}
	return len(rd) == 0
}
