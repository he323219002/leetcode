package main

import "fmt"

func main() {
	res := numTilings(4)
	fmt.Println(res)
}

func numTilings(n int) int {
	// d[1] = 1, d[2] = 2,d[3] = 5,
	// d[4] = d[3]*2-3 + d[2]*d[2] = 11
	// d[5] = d[4] * 2 - 3 + d[3] * 2 = 19 +
	// d[n] = d[n-2] * d[2] + d[n-1] * 2 -3

	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 5
	}

	d := make([]int, 3)
	d[0] = 1
	d[1] = 2
	d[2] = 5

	for i := 4; i <= n; i++ {
		d[0], d[1] = d[1], d[2]
		d[2] = d[0]*2 + d[1]*2 - 3
	}
	return d[2]
}
