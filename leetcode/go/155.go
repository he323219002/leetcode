package main

func main() {

}

type MinStack struct {
	stack     []int
	tempStack []int
}

func Constructor() MinStack {
	minStack := MinStack{}
	minStack.stack = make([]int, 0)
	minStack.tempStack = make([]int, 0)
	return minStack
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	// 最小栈为空
	if len(this.tempStack) == 0 {
		this.tempStack = append(this.tempStack, val)
		return
	}

	// 最小数
	curMin := this.tempStack[len(this.tempStack)-1]
	if curMin <= val {
		this.tempStack = append(this.tempStack, curMin)
	} else {
		this.tempStack = append(this.tempStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.tempStack = this.tempStack[:len(this.tempStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.tempStack[len(this.tempStack)-1]
}
