/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */
package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{[]int{1}}, 0))
}

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1
	for left <= right {
		mid := (left + right) / 2
		// fmt.Println("a", left, right, matrix[mid][0])
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right >= 0 {
		return binarySearch(matrix[right], target)
	}
	return false
}

func binarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		// fmt.Println(left, right, arr[mid])
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// @lc code=end
