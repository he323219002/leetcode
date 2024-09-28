package main

import (
	"fmt"
	"math"
)

func main() {
	// asteroids := []int{5, 10, -5}
	// asteroids := []int{10, 2, -5}
	asteroids := []int{1, -2, -2, -2}
	res := asteroidCollision(asteroids)
	fmt.Println(res)
}

func asteroidCollision(asteroids []int) []int {
	// 构造栈，反向小的出栈，
	// 0..n 如果都是负数则不考虑，因为向左不会和右边相遇
	stack := make([]int, 0)

	for _, val := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, val)
			continue
		}

		for lastOne := stack[len(stack)-1]; len(stack) > 0; lastOne = stack[len(stack)-1] {
			if lastOne < 0 {
				stack = append(stack, val)
				break
			}

			if float32(lastOne)/float32(val) > 0 {
				stack = append(stack, val)
				break
			} else if math.Abs(float64(lastOne)) > math.Abs(float64(val)) {
				break
			} else if math.Abs(float64(lastOne)) == math.Abs(float64(val)) {
				stack = stack[:len(stack)-1]
				break
			} else {
				stack = stack[:len(stack)-1]
			}

			if len(stack) == 0 {
				stack = append(stack, val)
				break
			}
		}
	}
	return stack
}
