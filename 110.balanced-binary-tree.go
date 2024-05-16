/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	_, res := maxDepth(root)
	return res
}
func maxDepth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	l, lt := maxDepth(root.Left)
	r, rt := maxDepth(root.Right)
	if !lt || !rt || l-r > 1 || l-r < -1 {
		return 0, false
	}

	return 1 + max(l, r), true
}

// @lc code=end
