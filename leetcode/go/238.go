package main

func main() {
	nums := []int{1, 2, 3, 4, 5}
	productExceptSelf(nums)
}

func productExceptSelf(nums []int) []int {
	// leetcode
	// 下三角 b[0] = 1,b[1] = a[0] * b[0],b[2] = b[1] * a[1],...,b[n] = b[n-1]* a[n-1]
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res[i] = 1
			continue
		}
		res[i] = res[i-1] * nums[i-1]
	}

	// 上三角 b[3] *= a[4],b[2] *= a[4]*a[5]
	// len(nums) - 1 结果不变
	var temp int = 1
	for i := len(nums) - 2; i >= 0; i-- {
		temp = temp * nums[i+1]
		res[i] = res[i] * temp
	}

	return res

}
