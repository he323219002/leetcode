package main

func lengthOfLongestSubstring(s string) int {
	// leetcode
	// 初始化字符下标索引
	sDic := make(map[byte]int)
	i, res, len := -1, 0, len(s)
	// 指针j遍历
	for j := 0; j < len; j++ {
		_, ok := sDic[s[j]]
		if ok {
			i = max(sDic[s[j]], i)
		}
		sDic[s[j]] = j
		res = max(res, j-i)
	}
	return res
}

func main() {
	lengthOfLongestSubstring("bs")
}
