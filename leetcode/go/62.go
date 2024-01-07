package main

func main() {

}

func uniquePaths(m int, n int) int {
	d := make([]int, n)
	for i := range d {
		d[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			d[j] = d[j-1] + d[j]
		}
	}
	return d[n-1]
}
