/*
 * @lc app=leetcode id=1448 lang=golang
 *
 * [1448] Count Good Nodes in Binary Tree
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
func goodNodes(root *TreeNode) int {
	maxVal := root.Val
	goodNodes := 1

	goodNodes += checkNodes(root.Left, maxVal)
	goodNodes += checkNodes(root.Right, maxVal)
	return goodNodes
}

func checkNodes(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}
	goodNodes := 0
	if root.Val >= maxVal {
		goodNodes++
		maxVal = root.Val
	}
	goodNodes += checkNodes(root.Left, maxVal)
	goodNodes += checkNodes(root.Right, maxVal)
	return goodNodes
}

// @lc code=end
