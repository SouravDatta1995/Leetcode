/*
 * @lc app=leetcode id=1046 lang=golang
 *
 * [1046] Last Stone Weight
 */
package main

import "fmt"

func main() {
	fmt.Println(lastStoneWeight([]int{10, 5, 4, 10, 3, 1, 7, 8}))
}

// @lc code=start
type Heap struct {
	arr []int
}

func lastStoneWeight(stones []int) int {
	var heap Heap

	for _, stone := range stones {
		push(&heap, stone)
		fmt.Println(heap)
	}
	for len(heap.arr) > 1 {
		s1 := pop(&heap)
		s2 := pop(&heap)
		fmt.Println(heap, s1, s2)
		if s1 > s2 {
			push(&heap, s1-s2)
		}
	}
	if len(heap.arr) == 0 {
		return 0
	}
	return heap.arr[0]
}

func push(this *Heap, x int) {
	if len(this.arr) == 0 {
		this.arr = append(this.arr, x)
	} else {
		this.arr = append(this.arr, x)
		i := len(this.arr) - 1

		for i > 0 && this.arr[i] > this.arr[parent(i)] {
			this.arr[i], this.arr[parent(i)], i = this.arr[parent(i)], this.arr[i], parent(i)
		}
	}
}

func pop(this *Heap) int {
	res := this.arr[0]
	this.arr[0] = this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]
	heapify(this)
	return res
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

func heapify(this *Heap) {
	i := 0
	for left(i) < len(this.arr) {
		if right(i) < len(this.arr) && this.arr[right(i)] > this.arr[i] && this.arr[right(i)] > this.arr[left(i)] {
			this.arr[i], this.arr[right(i)] = this.arr[right(i)], this.arr[i]
			i = right(i)
		} else if this.arr[left(i)] > this.arr[i] {
			this.arr[i], this.arr[left(i)] = this.arr[left(i)], this.arr[i]
			i = left(i)
		} else {
			return
		}
	}
}

// @lc code=end
