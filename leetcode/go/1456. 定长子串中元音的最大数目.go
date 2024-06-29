package main

import "fmt"

func main() {
	// res := maxVowels("abciiidef", 3)
	res := maxVowels("weallloveyou", 7)
	// res := maxVowels("leetcode", 3)
	fmt.Println(res)
}

func maxVowels(s string, k int) int {
	// 滑动窗口
	vowelMap := make(map[rune]bool, 5)
	vowelMap['a'] = true
	vowelMap['e'] = true
	vowelMap['i'] = true
	vowelMap['o'] = true
	vowelMap['u'] = true

	storeRes := 0
	for i := 0; i < k; i++ {
		if vowelMap[rune(s[i])] {
			storeRes += 1
		}
		if storeRes == k {
			return k
		}
	}
	res := storeRes

	for i := 1; i < len(s)-k+1; i++ {
		if vowelMap[rune(s[i-1])] {
			storeRes -= 1
		}
		if vowelMap[rune(s[i+k-1])] {
			storeRes += 1
			res = max(res, storeRes)
		}
		if storeRes == k {
			return k
		}
	}
	return res
}
