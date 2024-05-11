/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	first, second := head, head.Next

	for second.Next != nil && second.Next.Next != nil && first != second {
		first = first.Next
		second = second.Next.Next
	}
	return first == second
}

// @lc code=end
