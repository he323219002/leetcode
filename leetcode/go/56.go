package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergeLi := make([][]int, 0)
	for _, perLi := range intervals {
		if len(mergeLi) == 0 {
			mergeLi = append(mergeLi, perLi)
			continue
		}
		perLiLeft, perLiRight := perLi[0], perLi[1]
		lastNum := mergeLi[len(mergeLi)-1][1]

		if lastNum < perLiLeft {
			mergeLi = append(mergeLi, perLi)
			continue
		}
		if lastNum < perLiRight {
			mergeLi[len(mergeLi)-1][1] = perLiRight
			continue
		}
	}
	return mergeLi
}

func main() {

}
