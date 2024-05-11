/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */
package main

import (
	"strconv"
)

// @lc code=start
func evalRPN(tokens []string) int {

	tArr := []string{}

	for _, elem := range tokens {
		l := len(tArr)
		switch elem {
		case "+":
			val1, _ := strconv.Atoi(tArr[l-2])
			val2, _ := strconv.Atoi(tArr[l-1])
			tArr = append(tArr[:l-2], strconv.Itoa(val1+val2))
			break
		case "-":
			val1, _ := strconv.Atoi(tArr[l-2])
			val2, _ := strconv.Atoi(tArr[l-1])
			tArr = append(tArr[:l-2], strconv.Itoa(val1-val2))
			break
		case "*":
			val1, _ := strconv.Atoi(tArr[l-2])
			val2, _ := strconv.Atoi(tArr[l-1])
			tArr = append(tArr[:l-2], strconv.Itoa(val1*val2))
			break
		case "/":
			val1, _ := strconv.Atoi(tArr[l-2])
			val2, _ := strconv.Atoi(tArr[l-1])
			tArr = append(tArr[:l-2], strconv.Itoa(val1/val2))
			break
		default:
			tArr = append(tArr, elem)
		}
	}
	res, _ := strconv.Atoi(tArr[0])
	return res
}

// @lc code=end
