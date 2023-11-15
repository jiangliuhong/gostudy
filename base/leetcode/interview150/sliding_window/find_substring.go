package sliding_window

// 30. 串联所有单词的子串
func findSubstring(s string, words []string) []int {
	wordSize, wordItemSize, ls := len(words), len(words[0]), len(s)
	var res []int

	for i := 0; i+wordItemSize*wordSize <= ls; i++ {
		wordDict := map[string]int{}
		for j := 0; j < wordSize; j++ {
			subS := s[i+j*wordItemSize : i+(j+1)*wordItemSize]
			wordDict[subS]++
		}
		for _, word := range words {
			wordDict[word]--
			if wordDict[word] == 0 {
				delete(wordDict, word)
			}
		}
		if len(wordDict) == 0 {
			res = append(res, i)
		}

	}
	return res
}

func findSubstring2(s string, words []string) []int {
	wordItemSize, ls := len(words[0]), len(s)
	wordDict := map[string]int{}
	var res []int
	for _, item := range words {
		wordDict[item] = wordDict[item] + 1
	}
	for i := 0; i+wordItemSize <= ls; i++ {
		wordDictTemp := map[string]int{}
		for k, v := range wordDict {
			wordDictTemp[k] = v
		}
		for j := i; j+wordItemSize <= ls && len(wordDictTemp) > 0; j = j + wordItemSize {
			subS := s[j : j+wordItemSize]
			if v, ok := wordDictTemp[subS]; ok {
				v = v - 1
				if v <= 0 {
					delete(wordDictTemp, subS)
				} else {
					wordDictTemp[subS] = v
				}
			} else {
				break
			}
		}
		if len(wordDictTemp) == 0 {
			res = append(res, i)
		}
	}
	return res
}

/*
func findSubstring(s string, words []string) []int {
	res := []int{}
	wordItemSize := len(words[0])
	indexDict := map[int][]int{}
	for i := 0; i < len(s); i++ {
		if i+wordItemSize >= len(s) {
			break
		}
		subS := s[i : i+wordItemSize]
		for j := 0; j < len(words); j++ {
			if subS == words[j] {
				var indexDictItem []int
				if item, ok := indexDict[j]; ok {
					indexDictItem = item
				}
				indexDictItem = append(indexDictItem, i)
				indexDict[j] = indexDictItem
			}
		}
	}
	for item := range indexDict {

	}
	return res
}
*/
