package main

import (
	"fmt"
)

func main() {
	k, n := 3, 7
	res := combinationSum3(k, n)
	fmt.Println(res)
}

func combinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	if n < k {
		return ans
	}

	var backSearch func(tempRes *[]int, target int, floor int)
	backSearch = func(tempRes *[]int, target int, floor int) {
		if target == 0 && len(*tempRes) == k {
			ans = append(ans, append([]int(nil), (*tempRes)...))
			*tempRes = (*tempRes)[:len(*tempRes)-1]
			return
		} else if floor > target {
			*tempRes = (*tempRes)[:len(*tempRes)-1]
			return
		}
		for i := floor; i <= 9; i++ {
			if i > target {
				continue
			}
			*tempRes = append((*tempRes), i)
			backSearch(tempRes, target-i, i+1)
		}
		if len(*tempRes) > 0 {
			*tempRes = (*tempRes)[:len(*tempRes)-1]
		}
	}

	res := make([]int, 0)
	backSearch(&res, n, 1)

	return ans
}
