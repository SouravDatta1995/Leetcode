/*
 * @lc app=leetcode id=605 lang=golang
 *
 * [605] Can Place Flowers
 */
package main

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := n
	if len(flowerbed) == 1 && ((flowerbed[0] == 0 && n == 1) || n == 0) {
		return true
	}

	for i, v := range flowerbed {
		if v != 1 {
			if i == 0 {
				if i < len(flowerbed) && flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					count--
				}
			} else if i == len(flowerbed)-1 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				count--
			} else if flowerbed[i-1] == 0 {
				if flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					count--
				}
			}

		}
	}
	return count <= 0
}

// @lc code=end
