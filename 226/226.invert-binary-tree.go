/*
 * @lc app=leetcode id=226 lang=golang
 *
 * [226] Invert Binary Tree
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
func invertTree(root *TreeNode) *TreeNode {
	node := root
	if node == nil {
		return node
	}
	temp := node.Left
	node.Left = node.Right
	node.Right = temp
	if node.Left != nil {
		invertTree(node.Left)
	}
	if node.Right != nil {
		invertTree(node.Right)
	}
	return root
}

// @lc code=end
