/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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
 func kthSmallest(root *TreeNode, k int) int {
	vals := []int{}
	stack := []int{}

	makeArr(root, vals, stack)
	return vals[k-1]

}

func makeArr(root *TreeNode, vals, stack []int)  {
	if root == nil {
        if len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		vals = append(vals, v)
        }
        return
	}
    fmt.Println(root.Val,vals,stack)
	stack = append(vals, root.Val)
	makeArr(root.Left, vals,stack)

	makeArr(root.Right,vals,stack)
	return
}

// @lc code=end
