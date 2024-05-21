/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
 func isValidBST(root *TreeNode) bool {

	return validLeftSubTree(root.Left, root.Val, -10000000000) && validRightSubTree(root.Right, root.Val, 10000000000)

}

func validLeftSubTree(root *TreeNode, minVal int, maxVal int) bool {
	if root == nil {
		return true
	}
    // fmt.Println(root.Val,minVal,maxVal)
	if root.Val < minVal && root.Val > maxVal {
		return validLeftSubTree(root.Left, root.Val, maxVal) && validRightSubTree(root.Right, root.Val, minVal)
	}
	return false
}
func validRightSubTree(root *TreeNode, minVal int, maxVal int) bool {
	if root == nil {
		return true
	}
    // fmt.Println(root.Val,minVal,maxVal)
	if root.Val > minVal && root.Val < maxVal {
		return validLeftSubTree(root.Left, root.Val, minVal) && validRightSubTree(root.Right, root.Val, maxVal)
	}
	return false
}

// @lc code=end
