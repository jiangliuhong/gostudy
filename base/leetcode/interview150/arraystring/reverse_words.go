package arraystring

import "strings"

// 151. 反转字符串中的单词
func reverseWords(s string) string {
	words := make([]string, 0)
	w := ""
	for i := len(s) - 1; i >= 0; i-- {
		if s[i:i+1] == " " {
			if w != "" {
				words = append(words, w)
				w = ""
			}
		} else {
			w = s[i:i+1] + w
		}
	}
	if w != "" {
		words = append(words, w)
	}
	return strings.Join(words, " ")
}
