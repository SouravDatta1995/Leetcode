/*
 * @lc app=leetcode id=424 lang=golang
 *
 * [424] Longest Repeating Character Replacement
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(characterReplacement("AABABBA", 2))
}

// @lc code=start
func characterReplacement(s string, k int) int {
	charMap := make(map[rune]int)
	l := 0
	maxCount := 0

	for r := 0; r < len(s); r++ {

		charMap[rune(s[r])] = charMap[rune(s[r])] + 1

		// fmt.Println(l, r, charMap)
		charmax := 0
		for _, v := range charMap {
			charmax = max(charmax, v)
		}
		for (r+1-l)-charmax > k {

			charMap[rune(s[l])]--
			l++
			for _, v := range charMap {
				charmax = max(charmax, v)
			}

		}
		// fmt.Println(l, r, charMap)
		maxCount = max(maxCount, r+1-l)
	}

	return maxCount
}

// @lc code=end
