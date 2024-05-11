/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */
package main

// @lc code=start
func findDuplicate(nums []int) int {
	numMap := make([]bool, 100001)

	for _, n := range nums {
		if numMap[n] == true {
			return n
		}
		numMap[n] = true
	}
	return 0
}

// @lc code=end
