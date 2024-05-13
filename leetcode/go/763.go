package main

import "fmt"

func main() {
	s := "ababcbacadefegdehijhklij"
	// s := "eccbbbbdec"
	// s := "eaaaabaaec"
	res := partitionLabels(s)
	fmt.Println(res)
}

func partitionLabels(s string) []int {
	// 左右指针 维护一个最大右索引
	// 如果左指针不超过最大左索引，则为一个区间
	res := make([]int, 0)
	sMap := make(map[rune]int, 0)
	for index, val := range s {
		sMap[val] = index
	}
	leftPoint, rightPoint := 0, 0

	for i := 0; i < len(s); i++ {
		if rightPoint == len(s)-1 {
			res = append(res, rightPoint-leftPoint+1)
			return res
		}

		if i <= rightPoint {
			// 字符最大索引位置,维护右指针
			cur := rune(s[i])
			rightPoint = max(rightPoint, sMap[cur])
			continue
		}
		if i > rightPoint {
			// 结算上一个区间并开始新一个区间计算
			res = append(res, rightPoint-leftPoint+1)
			leftPoint = i
			// 字符最大索引位置,维护右指针
			cur := rune(s[i])
			rightPoint = max(rightPoint, sMap[cur])

			// 如果为最后一个字符
			if i == len(s)-1 {
				res = append(res, i-rightPoint+1)
				return res
			}

			continue
		}

	}
	return res
}
