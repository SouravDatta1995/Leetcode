/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
package main

import (
	"fmt"
)

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	fmt.Println(removeNthFromEnd(l1, 2))
}

type ListNode struct {
	Next *ListNode
	Val  int
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = reverseList(head)
	temp1, temp2 := head, head.Next
	if n == 1 {
		head = head.Next
		return reverseList(head)
	}
	for i := 1; i < n-1; i++ {
		temp1 = temp1.Next
		temp2 = temp2.Next
	}
	temp1.Next = temp2.Next
	return reverseList(head)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

// @lc code=end
