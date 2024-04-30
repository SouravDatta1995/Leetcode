/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
package main

import (
	"fmt"
)

func main() {
	inp := 3
	fmt.Println(generateParenthesis(inp))
}

// @lc code=start
func generateParenthesis(n int) []string {
	var res []string
	recurParenthisis(0, 0, n, &res, "")
	return res

}
func recurParenthisis(open, close, max int, stack *[]string, seq string) {
	fmt.Println(open, close, max, stack)
	if close == open && open == max {
		*stack = append(*stack, seq)
		return
	}

	if open < max {
		seq += "("
		open++
		recurParenthisis(open, close, max, stack, seq)
		open--
		seq = seq[:len(seq)-1]
		fmt.Println(seq)
	}
	if close < open {
		seq += ")"
		close++
		recurParenthisis(open, close, max, stack, seq)
		close--
		seq = seq[:len(seq)-1]
		fmt.Println(seq)
	}
	// fmt.Println(*stack)
}

// @lc code=end
