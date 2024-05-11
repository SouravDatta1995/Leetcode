/*
 * @lc app=leetcode id=875 lang=golang
 *
 * [875] Koko Eating Bananas
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	// fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	// fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
	fmt.Println(minEatingSpeed([]int{332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589, 290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184}, 823855818))
}

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, num := range piles {
		if num > maxPile {
			maxPile = num
		}
	}

	left, right := 1, maxPile
	for left <= right {
		mid := (left + right) / 2
		// fmt.Println(mid)
		op := 0
		for _, v := range piles {
			// fmt.Println(pile)
			op += int(math.Ceil(float64(v) / float64(mid)))
		}
		if op > h {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left

}

// @lc code=end
