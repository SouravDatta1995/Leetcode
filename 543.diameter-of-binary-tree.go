/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	md := 0
	maxDepth(root, &md)
	return md
}
func maxDepth(root *TreeNode, md *int) int {
	if root == nil {
		return 0
	}
	left, right := maxDepth(root.Left, md), maxDepth(root.Right, md)
	*md = max(*md, left+right)

	return 1 + max(left, right)
}

// @lc code=end
