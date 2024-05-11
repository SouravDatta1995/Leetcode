/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	addTwoNumbers(l1, l2)
}

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resList *ListNode
	var head *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		res := 0
		if l1 != nil {
			res += l1.Val
		}
		if l2 != nil {
			res += l2.Val
		}
		res += carry
		if res < 10 {
			res = res % 10
			carry = 0
		} else {
			carry = res / 10
			res = res - 10
		}
		node := &ListNode{Val: res, Next: nil}
		if head == nil {
			head = node
			resList = node
		} else {
			resList.Next = &ListNode{}
			resList = resList.Next
			resList.Val = node.Val
			resList.Next = node.Next
		}
		fmt.Println(node)
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		node := &ListNode{Val: carry, Next: nil}
		resList.Next = &ListNode{}
		resList = resList.Next
		resList.Val = node.Val
		resList.Next = node.Next
	}

	return head
}

// @lc code=end
