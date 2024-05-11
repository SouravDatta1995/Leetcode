/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */
package main

// @lc code=start
type MinStack struct {
	stack []int
	min   []int
	top   int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}, -1}
}

func (this *MinStack) Push(val int) {
	if this.top == -1 {
		this.min = append(this.min, val)
	} else if this.min[this.top] < val {
		this.min = append(this.min, this.min[this.top])
	} else {
		this.min = append(this.min, val)
	}

	this.stack = append(this.stack, val)
	this.top++
	// fmt.Println(this)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.top]
	this.min = this.min[:this.top]
	this.top = this.top - 1
	// fmt.Println(this)
}

func (this *MinStack) Top() int {

	// fmt.Println(this)
	return this.stack[this.top]
}

func (this *MinStack) GetMin() int {
	// fmt.Println(this)
	return this.min[this.top]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
