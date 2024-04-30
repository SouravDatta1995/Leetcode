/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
package main

import (
	"fmt"
)

func main() {
	// fmt.Println(lengthOfLongestSubstring("dvdf"))
	// fmt.Println(lengthOfLongestSubstring("b"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	result := 0
	myMap := make(map[rune]int)
	start := 0

	for i, v := range s {
		if idx, ok := myMap[v]; ok && idx >= start {
			start = idx + 1
		}
		myMap[v] = i
		result = max(result, i-start+1)
	}

	return result
}

// @lc code=end
