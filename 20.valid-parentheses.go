/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
}

// @lc code=start
func isValid(s string) bool {
	pSck := make([]byte, len(s))
	i := 0
	for _, elem := range s {
		// fmt.Println(elem, i, pSck)
		if elem == '(' || elem == '{' || elem == '[' {
			pSck[i] = byte(elem)
			i++
		} else if elem == ')' {
			if i == 0 || pSck[i-1] != '(' {
				return false
			} else {
				i--
			}
		} else if elem == '}' {
			if i == 0 || pSck[i-1] != '{' {
				return false
			} else {
				i--
			}
		} else if elem == ']' {
			if i == 0 || pSck[i-1] != '[' {
				return false
			} else {
				i--
			}
		}

	}
	if i == 0 {
		return true
	}
	return false
}

// @lc code=end
