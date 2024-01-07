package main

func main() {

}

func rob(nums []int) int {
	prev, curr := 0, 0
	for _, v := range nums {
		prev, curr = curr, max(curr, prev+v)
	}

	return curr
}
