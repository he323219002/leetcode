package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	res := eraseOverlapIntervals(intervals)
	fmt.Println(res)
}

func eraseOverlapIntervals(intervals [][]int) int {
	// 去除最少的重复 = 找最多的不重复
	// leetcode 所有区间按右端点排序，找小的右端点
	count := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	tempRes := make([][]int, 0)
	for len(intervals) > 0 {
		if len(tempRes) == 0 {
			tempRes = append(tempRes, intervals[0])
			continue
		}
		if intervals[0][0] < tempRes[len(tempRes)-1][1] {
			intervals = intervals[1:]
			continue
		}
		tempRes = append(tempRes, intervals[0])

	}

	return count - len(tempRes)
}
