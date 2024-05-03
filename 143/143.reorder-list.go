/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */
package main

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
	reorderList(l1)
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
func reorderList(head *ListNode) {
	temp1 := head
	for temp1.Next.Next != nil {
		temp2 := head
		for temp2.Next.Next != nil {
			temp2 = temp2.Next
		}
		last := temp2.Next
		temp2.Next = nil
		second := temp1.Next
		temp1.Next = last
		last.Next = second
		temp1 = temp1.Next
	}
}

// @lc code=end
