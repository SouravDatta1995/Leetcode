/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1}, []int{0}))
}

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totLen := len(nums1) + len(nums2)
	finArr := make([]int, totLen/2+1)
	l, r := 0, 0
	for l+r < (totLen/2 + 1) {
		if l < len(nums1) && (r == len(nums2) || nums1[l] <= nums2[r]) {
			finArr[l+r] = nums1[l]
			l++
		} else {
			finArr[l+r] = nums2[r]
			if r < len(nums2) {
				r++
			}
		}
	}
	if totLen%2 == 0 {
		return float64(finArr[totLen/2]+finArr[totLen/2-1]) / 2
	} else {
		return float64(finArr[totLen/2])
	}
}

// @lc code=end
