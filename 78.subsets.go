/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}

// @lc code=start
func subsets(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
w
		lres := len(res)
		for j := 0; j < lres; j++ {
			subset := make([]int, len(res[j]))
			copy(subset, res[j])
			subset = append(subset, nums[i])
			res = append(res, subset)
		}
	}
	return res
}

// @lc code=end
