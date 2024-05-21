/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{3, 5, 8}, 11))
}

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	var dfs func([]int, int, int)
	dfs = func(cur []int, i, total int) {
		if total == target {
			res = append(res, append([]int{}, cur...))
			return
		}
		if i >= len(candidates) || total > target {
			return
		}
		cur = append(cur, candidates[i])
		dfs(cur, i, total+candidates[i])
		dfs(cur[:len(cur)-1], i+1, total)
	}

	dfs([]int{}, 0, 0)

	return res
}

// @lc code=end
