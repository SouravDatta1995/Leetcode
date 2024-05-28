/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */
package main

// @lc code=start
type Heap struct {
	arr []int
}

func findKthLargest(nums []int, k int) int {
	h := new(Heap)

	for _, n := range nums {
		h.push(n, k)
	}
	return h.arr[0]
}

func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func (h *Heap) push(x int, k int) {
	if len(h.arr) == 0 {
		h.arr = append(h.arr, x)
	} else if len(h.arr) < k {
		h.arr = append(h.arr, x)
		i := len(h.arr) - 1
		for i > 0 && h.arr[i] < h.arr[parent(i)] {
			h.arr[i], h.arr[parent(i)], i = h.arr[parent(i)], h.arr[i], parent(i)
		}
	} else if h.arr[0] < x {
		h.arr = append(h.arr, x)
		h.pop()
	}

}

func (h *Heap) pop() int {
	res := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapify()
	return res
}

func (h *Heap) heapify() {
	i := 0
	for left(i) < len(h.arr) {
		if right(i) < len(h.arr) && h.arr[i] > h.arr[right(i)] && h.arr[right(i)] < h.arr[left(i)] {
			h.arr[i], h.arr[right(i)], i = h.arr[right(i)], h.arr[i], right(i)
		} else if h.arr[i] > h.arr[left(i)] {
			h.arr[i], h.arr[left(i)], i = h.arr[left(i)], h.arr[i], left(i)
		} else {
			break
		}
	}
}

// @lc code=end
