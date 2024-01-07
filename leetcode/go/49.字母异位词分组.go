package main

import (
	"sort"
)

// class Solution:
//     def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
//         str_dic = dict()
//         for index,str in enumerate(strs):
//             new_str = "".join(sorted(str))
//             index_arr = str_dic.get(new_str,[])
//             index_arr.append(index)
//             str_dic[new_str] = index_arr

//         res = []
//         for k,v in str_dic.items():
//             res.append([strs[i] for i in v])
//         return res

// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		// 排序
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		// 构造
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	// 构造，长度可以为0
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
