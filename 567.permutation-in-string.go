/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */
package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidboaoo"))
	// fmt.Println(checkInclusion("ab", "eidboaoo"))
}

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	charMap := make(map[rune]int)
	for _, r := range s1 {
		charMap[r]++
	}
	for i := 0; i < len(s2)-len(s1)+1; i++ {
		if charMap[rune(s2[i])] == 0 {
			continue
		} else {
			tempMap := make(map[rune]int)
			for k, v := range charMap {
				tempMap[k] = v
			}
			// fmt.Println(s2[i : i+len(s1)])
			for _, c := range s2[i : i+len(s1)] {
				tempMap[c]--
			}
			// fmt.Println(tempMap)
			fl := true
			for _, t := range tempMap {
				if t != 0 {
					fl = false
					break
				}
			}
			// fmt.Println(fl)
			if fl {
				return true
			}
		}
	}
	return false
}

// @lc code=end
