package main

func main() {

}

func canJump(nums []int) bool {
	// 维护最远长度
	right := 0
	for index, v := range nums {
		// 只处理最远范围内的
		if index <= right {
			right = max(right, index+v)
			if right >= len(nums)-1 {
				return true
			}
		}
	}
	return false
}
