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
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
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
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	var prev *ListNode
	slow.Next = nil
	for second != nil {
		tmp := second.Next
		second.Next = prev
		prev = second
		second = tmp
	}
	first := head
	second = prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}

}

// @lc code=end
