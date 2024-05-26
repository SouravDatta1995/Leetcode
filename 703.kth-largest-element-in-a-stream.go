/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 */
package main

// @lc code=start
type KthLargest struct {
	arr      []int
	heapSize int
	lgix     int
}

func Constructor(k int, nums []int) KthLargest {
	heap := KthLargest{arr: []int{}, heapSize: 0, lgix: k}
	for _, v := range nums {
		heap.Add(v)
	}
	return heap
}

func (this *KthLargest) Add(val int) int {
	if this.heapSize <= this.lgix {
		this.heapSize++
		this.arr = append(this.arr, val)
		i := this.heapSize - 1
		for i >= 0 && this.arr[parent(i)] >= this.arr[i] {
			this.arr[parent(i)], this.arr[i], i = this.arr[i], this.arr[parent(i)], parent(i)
		}
	} else if val > this.arr[0] {
		// fmt.Println(this.arr)
		this.arr[0] = val
		reorder(this)
		// fmt.Println(this.arr)
	}
	return this.arr[0]

}
func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return (2 * i) + 1
}
func right(i int) int {
	return (2 * i) + 2
}
func reorder(this *KthLargest) {
	i := 0
	if left(i) < this.heapSize {
		if right(i) < this.heapSize && this.arr[right(i)] < this.arr[i] && this.arr[right(i)] < this.arr[left(i)] {
			this.arr[i], this.arr[right(i)] = this.arr[right(i)], this.arr[i]
			i = right(i)
		} else if this.arr[left(i)] < this.arr[i] {
			this.arr[i], this.arr[left(i)] = this.arr[left(i)], this.arr[i]
			i = left(i)
		} else {
			return
		}
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end
