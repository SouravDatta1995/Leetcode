/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */
package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}
	var backtrack func(int, []int)

	backtrack = func(i int, subset []int) {
		if i == len(nums) {
			res = append(res, append([]int{}, subset...))
			return
		}

		subset = append(subset, nums[i])
		backtrack(i+1, subset)
		subset = subset[:len(subset)-1]

		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}

		backtrack(i+1, subset)
	}
	backtrack(0, []int{})
	return res
}

// @lc code=end
