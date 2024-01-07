package main

import "fmt"

func maxArea(height []int) int {
	i, j, res := 0, len(height)-1, 0
	for {
		if i == j {
			break
		}
		if height[i] < height[j] {
			res = max(res, height[i]*(j-i))
			i += 1
		} else {
			res = max(res, height[j]*(j-i))
			j -= 1
		}
	}
	return res
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(height)
	fmt.Println(result)
}
