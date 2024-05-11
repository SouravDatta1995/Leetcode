package main

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	mergeTwoLists(l1, l2)
}

type ListNode struct {
	Next *ListNode
	Val  int
}

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	resutList := &ListNode{}
	currList := resutList
	for list1 != nil && list2 != nil {
		currList.Next = &ListNode{}
		currList = currList.Next
		if list1.Val < list2.Val {
			currList.Val = list1.Val
			list1 = list1.Next
		} else {
			currList.Val = list2.Val
			list2 = list2.Next
		}

	}
	for list1 != nil {
		currList.Next = &ListNode{}
		currList = currList.Next
		currList.Val = list1.Val
		list1 = list1.Next
	}
	for list2 != nil {
		currList.Next = &ListNode{}
		currList = currList.Next
		currList.Val = list2.Val
		list2 = list2.Next
	}
	return resutList.Next
}

// @lc code=end
