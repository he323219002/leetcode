package main

import "fmt"

func main() {
	res := climbStairs(2)
	fmt.Println(res)
}

func climbStairs(n int) int {
	// d[1] = 1  d[0] = 1 d
	// d[n] = d[n-1] + d[n-2]
	if n == 1 {
		return 1
	}

	d := make([]int, 3)
	d[0], d[1] = 1, 1
	for i := 1; i < n; i++ {
		if i > 1 {
			d[0] = d[1]
			d[1] = d[2]
		}
		d[2] = d[0] + d[1]
	}

	return d[2]
}
