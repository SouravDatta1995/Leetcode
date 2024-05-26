/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}
	temp := nums[0]
	cur = append(cur, nums[0])
	nums = nums[1:]
	if len(nums) == 0 {
		res = append(res, append([]int{}, cur...))
		return res
	}

	for range len(nums) {
		permute(nums)
	}
	nums = append(nums, temp)

	return res
}

// @lc code=end
