package main

import "fmt"

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(temperatures)
	fmt.Println(res)
}

func dailyTemperatures(temperatures []int) []int {
	// 两个栈，分别存温度和天数
	tempStack := make([]int, 0)
	dayStack := make([]int, 0)
	res := make([]int, len(temperatures))

	for day, temp := range temperatures {
		// 温度，1.如果栈是空的，则入栈 2.小于等于前一个，则入栈
		if len(tempStack) == 0 {
			tempStack = append(tempStack, temp)
			// 对应的天数入栈
			dayStack = append(dayStack, day)
			continue
		} else if temp <= tempStack[len(tempStack)-1] {
			tempStack = append(tempStack, temp)
			dayStack = append(dayStack, day)
			continue
		}
		// 循环
		for temp > tempStack[len(tempStack)-1] {
			tempStack = tempStack[:len(tempStack)-1]
			lastDay := dayStack[len(dayStack)-1]
			dayStack = dayStack[:len(dayStack)-1]
			res[lastDay] = day - lastDay

			if len(tempStack) == 0 {
				tempStack = append(tempStack, temp)
				dayStack = append(dayStack, day)
				break
			} else if temp <= tempStack[len(tempStack)-1] {
				tempStack = append(tempStack, temp)
				dayStack = append(dayStack, day)
				break
			}
		}
	}

	return res
}
