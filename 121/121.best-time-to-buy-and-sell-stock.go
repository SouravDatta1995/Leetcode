/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */
package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

// @lc code=start
func maxProfit(prices []int) int {
	min := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > profit {
			profit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return profit
}

// @lc code=end
