/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	slices.Sort(candidates)
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
		dfs(cur, i+1, total+candidates[i])
		for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
			i++
		}
		dfs(cur[:len(cur)-1], i+1, total)
	}

	dfs([]int{}, 0, 0)

	return res
}

// @lc code=end
