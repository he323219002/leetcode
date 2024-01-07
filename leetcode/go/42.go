package main

func trap(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	leftMax := 0
	rightMax := 0
	for {
		if left >= right {
			break
		}
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left += 1
		} else {
			ans += rightMax - height[right]
			right -= 1
		}
	}
	return ans
}

func main() {

}
