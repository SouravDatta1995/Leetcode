/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */
package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea([]int{4, 2, 0, 3, 2, 5}))
}

// @lc code=start
type Rect struct {
	index  int
	height int
}

func largestRectangleArea(heights []int) int {
	stack := []Rect{}
	heights = append(heights, 0)
	maxArea := 0
	for i, h := range heights {
		start := i
		for len(stack) > 0 && h < stack[len(stack)-1].height {
			elem := pop(&stack)
			area := elem.height * (i - elem.index)
			if maxArea < area {
				maxArea = area
			}
			start = elem.index
		}

		push(&stack, Rect{start, h})
	}

	return maxArea
}

func push(this *[]Rect, val Rect) {
	*this = append(*this, val)
}

func pop(this *[]Rect) Rect {
	val := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return val
}

// @lc code=end
