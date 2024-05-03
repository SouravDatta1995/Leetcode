/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */
package main

import "fmt"

func main() {
	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temp))
}

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	stack := [][]int{}
	daysArr := make([]int, len(temperatures))
	for i, temp := range temperatures {
		for len(stack) > 0 && stack[len(stack)-1][0] < temp {
			daysArr[stack[len(stack)-1][1]] = i - stack[len(stack)-1][1]
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, []int{temp, i})

	}
	return daysArr
}

// @lc code=end
