/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */
package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	temp := head
	addressMap := make(map[*Node]*Node)
	for temp != nil {
		addressMap[temp] = &Node{temp.Val, nil, nil}
		temp = temp.Next
	}

	temp = head
	for temp != nil {
		addressMap[temp].Next = addressMap[temp.Next]
		addressMap[temp].Random = addressMap[temp.Random]
		temp = temp.Next
	}
	return addressMap[head]
}

// @lc code=end
