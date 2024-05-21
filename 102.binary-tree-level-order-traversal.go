/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	q := make([]*TreeNode, 1)
	q[0] = root
	level := 0
	for len(q) > 0 {
		currLen := len(q)
		for i := 0; i < currLen; i++ {
			node := q[0]
			q = q[1:]
			if len(res) <= level {
				res = append(res, []int{node.Val})
			} else {
				res[level] = append(res[level], node.Val)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		level++
	}
	return res
}

// @lc code=end
